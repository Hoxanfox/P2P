package pool

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// SocketConfig contiene la configuración para el pool de sockets
type SocketConfig struct {
	SocketPool struct {
		MaxConnections      int   `yaml:"max_connections"`
		InactiveTimeout     int   `yaml:"inactive_timeout"`
		HealthCheckInterval int   `yaml:"health_check_interval"`
		BufferSize          int   `yaml:"buffer_size"`
		WriteTimeout        int64 `yaml:"write_timeout"`
	} `yaml:"socket_pool"`
}

// ClientConnection encapsula una conexión de cliente y metadatos relacionados
type ClientConnection struct {
	Conn       net.Conn
	ID         uuid.UUID
	LastActive time.Time
	mu         sync.RWMutex // Para acceso seguro a LastActive
}

// SocketPool gestiona todas las conexiones de socket de los clientes
type SocketPool struct {
	connections      map[uuid.UUID]*ClientConnection
	mu               sync.RWMutex
	config           *SocketConfig
	log              *logrus.Logger
	connectionCount  int32
	healthCheckTimer *time.Ticker
	done             chan struct{}
	
	// Funciones callback para eventos del pool
	OnConnectionClosed func(id uuid.UUID)
}

// LoadSocketConfig carga la configuración desde un archivo YAML
func LoadSocketConfig(configPath string) (*SocketConfig, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error leyendo archivo de configuración: %w", err)
	}

	var config SocketConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error decodificando configuración YAML: %w", err)
	}

	return &config, nil
}

// NewSocketPool crea un nuevo pool de sockets desde un archivo de configuración
func NewSocketPool(configPath string) (*SocketPool, error) {
	config, err := LoadSocketConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("error cargando configuración: %w", err)
	}

	// Crear logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	pool := &SocketPool{
		connections: make(map[uuid.UUID]*ClientConnection),
		config:      config,
		log:         logger,
		done:        make(chan struct{}),
	}

	// Iniciar health checks en una goroutine separada
	healthCheckInterval := time.Duration(config.SocketPool.HealthCheckInterval) * time.Second
	pool.healthCheckTimer = time.NewTicker(healthCheckInterval)
	
	go pool.healthCheckLoop()

	logger.WithField("max_connections", config.SocketPool.MaxConnections).
		Info("Pool de sockets inicializado correctamente")
	return pool, nil
}

// Register añade una nueva conexión al pool
func (p *SocketPool) Register(id uuid.UUID, conn net.Conn) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Verificar si ya existe una conexión con este ID
	if _, exists := p.connections[id]; exists {
		// Cerrar la conexión existente antes de reemplazarla
		p.closeConnectionLocked(id)
	}

	// Verificar límite de conexiones máximas
	if len(p.connections) >= p.config.SocketPool.MaxConnections {
		return fmt.Errorf("se alcanzó el límite máximo de conexiones (%d)", p.config.SocketPool.MaxConnections)
	}

	// Registrar la nueva conexión
	p.connections[id] = &ClientConnection{
		Conn:       conn,
		ID:         id,
		LastActive: time.Now(),
	}

	p.log.WithFields(logrus.Fields{
		"client_id":         id.String(),
		"remote_addr":       conn.RemoteAddr().String(),
		"active_connections": len(p.connections),
	}).Info("Cliente registrado en el pool de sockets")

	return nil
}

// Get obtiene una conexión del pool por ID
func (p *SocketPool) Get(id uuid.UUID) (net.Conn, bool) {
	p.mu.RLock()
	client, exists := p.connections[id]
	p.mu.RUnlock()

	if !exists {
		return nil, false
	}

	// Actualizar el tiempo de última actividad
	client.mu.Lock()
	client.LastActive = time.Now()
	client.mu.Unlock()

	return client.Conn, true
}

// Release libera una conexión del pool y la cierra
func (p *SocketPool) Release(id uuid.UUID) {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	p.closeConnectionLocked(id)
}

// closeConnectionLocked cierra y elimina una conexión (debe ser llamado con el mutex adquirido)
func (p *SocketPool) closeConnectionLocked(id uuid.UUID) {
	client, exists := p.connections[id]
	if !exists {
		return
	}

	// Cerrar la conexión
	if client.Conn != nil {
		client.Conn.Close()
	}
	
	// Eliminar del mapa
	delete(p.connections, id)

	p.log.WithFields(logrus.Fields{
		"client_id":         id.String(),
		"active_connections": len(p.connections),
	}).Info("Cliente liberado del pool de sockets")
	
	// Ejecutar callback si está definido
	if p.OnConnectionClosed != nil {
		go p.OnConnectionClosed(id)
	}
}

// Broadcast envía un mensaje a múltiples clientes
func (p *SocketPool) Broadcast(ids []uuid.UUID, frame []byte) error {
	if len(frame) == 0 {
		return fmt.Errorf("frame vacío, no se puede enviar")
	}

	var failedClients []uuid.UUID
	var mu sync.Mutex // Para acceso seguro a failedClients
	var wg sync.WaitGroup

	for _, id := range ids {
		wg.Add(1)
		
		go func(clientID uuid.UUID) {
			defer wg.Done()
			
			// Obtener la conexión
			p.mu.RLock()
			client, exists := p.connections[clientID]
			p.mu.RUnlock()
			
			if !exists {
				mu.Lock()
				failedClients = append(failedClients, clientID)
				mu.Unlock()
				return
			}
			
			// Establecer timeout para la escritura
			writeTimeout := time.Duration(p.config.SocketPool.WriteTimeout) * time.Millisecond
			deadline := time.Now().Add(writeTimeout)
			client.Conn.SetWriteDeadline(deadline)
			
			// Enviar el frame
			_, err := client.Conn.Write(frame)
			
			// Limpiar timeout
			client.Conn.SetWriteDeadline(time.Time{})
			
			if err != nil {
				p.log.WithFields(logrus.Fields{
					"client_id": clientID.String(),
					"error":     err.Error(),
				}).Error("Error enviando mensaje a cliente")
				
				mu.Lock()
				failedClients = append(failedClients, clientID)
				mu.Unlock()
				
				// Cerrar la conexión si es un error de EOF o de red
				if err == io.EOF || isNetworkError(err) {
					p.Release(clientID)
				}
				return
			}
			
			// Actualizar tiempo de actividad
			client.mu.Lock()
			client.LastActive = time.Now()
			client.mu.Unlock()
		}(id)
	}

	// Esperar a que terminen todas las goroutines
	wg.Wait()

	if len(failedClients) > 0 {
		return fmt.Errorf("fallo al enviar mensaje a %d clientes", len(failedClients))
	}

	return nil
}

// isNetworkError determina si un error es un error de red
func isNetworkError(err error) bool {
	if _, ok := err.(net.Error); ok {
		return true
	}
	return false
}

// healthCheckLoop realiza comprobaciones periódicas de salud en las conexiones
func (p *SocketPool) healthCheckLoop() {
	for {
		select {
		case <-p.healthCheckTimer.C:
			p.performHealthCheck()
		case <-p.done:
			p.healthCheckTimer.Stop()
			return
		}
	}
}

// performHealthCheck verifica todas las conexiones y cierra las zombies
func (p *SocketPool) performHealthCheck() {
	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now()
	inactiveThreshold := time.Duration(p.config.SocketPool.InactiveTimeout) * time.Second
	zombieCount := 0

	for id, client := range p.connections {
		client.mu.RLock()
		lastActive := client.LastActive
		client.mu.RUnlock()

		// Verificar si la conexión está inactiva
		if now.Sub(lastActive) > inactiveThreshold {
			// Conexión zombie, cerrarla
			p.log.WithFields(logrus.Fields{
				"client_id":     id.String(),
				"last_active":   lastActive,
				"inactive_time": now.Sub(lastActive).String(),
			}).Info("Cerrando conexión inactiva (zombie)")
			
			p.closeConnectionLocked(id)
			zombieCount++
		}
	}

	if zombieCount > 0 {
		p.log.WithFields(logrus.Fields{
			"zombie_count":     zombieCount,
			"active_connections": len(p.connections),
		}).Info("Health check completado, conexiones zombie eliminadas")
	}
}

// Close cierra todas las conexiones y detiene el pool
func (p *SocketPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Detener el health check
	close(p.done)

	// Cerrar todas las conexiones
	for id := range p.connections {
		p.closeConnectionLocked(id)
	}

	p.log.Info("Pool de sockets cerrado")
}

// GetMetrics devuelve métricas del pool para monitorización
func (p *SocketPool) GetMetrics() map[string]interface{} {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	// Calcular el número de conexiones activas
	activeCount := len(p.connections)
	
	// Calcular el porcentaje de uso
	maxConns := p.config.SocketPool.MaxConnections
	utilizationPct := float64(activeCount) / float64(maxConns) * 100.0
	
	// Crear mapa de métricas
	metrics := map[string]interface{}{
		"active_connections": activeCount,
		"max_connections":    maxConns,
		"utilization_pct":    utilizationPct,
	}
	
	return metrics
}
