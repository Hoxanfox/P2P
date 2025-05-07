package pool

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// PeerPoolConfig contiene la configuración del pool de conexiones entre pares
type PeerPoolConfig struct {
	PeerPool struct {
		MaxPeers         int    `yaml:"max_peers"`
		DialTimeout      string `yaml:"dial_timeout"`
		HandshakeTimeout string `yaml:"handshake_timeout"`
		
		Reconnect struct {
			BaseDelay    string  `yaml:"base_delay"`
			MaxDelay     string  `yaml:"max_delay"`
			MaxAttempts  int     `yaml:"max_attempts"`
			JitterFactor float64 `yaml:"jitter_factor"`
		} `yaml:"reconnect"`
		
		TLS struct {
			CertFile   string `yaml:"cert_file"`
			KeyFile    string `yaml:"key_file"`
			CAFile     string `yaml:"ca_file"`
			ServerName string `yaml:"server_name"`
		} `yaml:"tls"`
		
		BufferSize   int `yaml:"buffer_size"`
		MaxFrameSize int `yaml:"max_frame_size"`
		
		Keepalive struct {
			Interval   string `yaml:"interval"`
			Timeout    string `yaml:"timeout"`
			MaxMissed  int    `yaml:"max_missed"`
		} `yaml:"keepalive"`
	} `yaml:"peer_pool"`
}

// PeerConnectionPool gestiona conexiones a otros nodos P2P
type PeerConnectionPool struct {
	config         *PeerPoolConfig
	connections    map[uuid.UUID]*PeerConn
	mu             sync.RWMutex
	log            *logrus.Logger
	dialTimeout    time.Duration
	handshakeTimeout time.Duration
	baseDelay      time.Duration
	maxDelay       time.Duration
	maxAttempts    int
	jitterFactor   float64
	keepaliveInterval time.Duration
	peerStateNotifier func(uuid.UUID, PeerState)
}

// LoadPeerPoolConfig carga la configuración desde un archivo YAML
func LoadPeerPoolConfig(configPath string) (*PeerPoolConfig, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error leyendo archivo de configuración: %w", err)
	}

	var config PeerPoolConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error decodificando configuración YAML: %w", err)
	}

	return &config, nil
}

// NewPeerConnectionPool crea un nuevo pool de conexiones P2P
func NewPeerConnectionPool(configPath string) (*PeerConnectionPool, error) {
	config, err := LoadPeerPoolConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("error cargando configuración: %w", err)
	}

	// Crear logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Parsear duraciones de la configuración
	dialTimeout, err := time.ParseDuration(config.PeerPool.DialTimeout)
	if err != nil {
		return nil, fmt.Errorf("error en la configuración dial_timeout: %w", err)
	}
	
	handshakeTimeout, err := time.ParseDuration(config.PeerPool.HandshakeTimeout)
	if err != nil {
		return nil, fmt.Errorf("error en la configuración handshake_timeout: %w", err)
	}
	
	baseDelay, err := time.ParseDuration(config.PeerPool.Reconnect.BaseDelay)
	if err != nil {
		return nil, fmt.Errorf("error en la configuración reconnect.base_delay: %w", err)
	}
	
	maxDelay, err := time.ParseDuration(config.PeerPool.Reconnect.MaxDelay)
	if err != nil {
		return nil, fmt.Errorf("error en la configuración reconnect.max_delay: %w", err)
	}
	
	keepaliveInterval, err := time.ParseDuration(config.PeerPool.Keepalive.Interval)
	if err != nil {
		return nil, fmt.Errorf("error en la configuración keepalive.interval: %w", err)
	}

	pool := &PeerConnectionPool{
		config:         config,
		connections:    make(map[uuid.UUID]*PeerConn),
		log:            logger,
		dialTimeout:    dialTimeout,
		handshakeTimeout: handshakeTimeout,
		baseDelay:      baseDelay,
		maxDelay:       maxDelay,
		maxAttempts:    config.PeerPool.Reconnect.MaxAttempts,
		jitterFactor:   config.PeerPool.Reconnect.JitterFactor,
		keepaliveInterval: keepaliveInterval,
	}

	// Iniciar rutina de keepalive para todas las conexiones
	go pool.keepaliveLoop()

	logger.WithField("max_peers", config.PeerPool.MaxPeers).
		Info("Pool de conexiones P2P inicializado correctamente")
	return pool, nil
}

// SetPeerStateNotifier establece una función para notificar cambios de estado
func (p *PeerConnectionPool) SetPeerStateNotifier(notifier func(uuid.UUID, PeerState)) {
	p.peerStateNotifier = notifier
}

// DialAndRegister conecta a un peer y lo registra en el pool
func (p *PeerConnectionPool) DialAndRegister(peer PeerInfo) (*PeerConn, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Verificar si ya existe una conexión
	if conn, exists := p.connections[peer.ID]; exists {
		p.log.WithField("peer_id", peer.ID).Debug("Conexión ya existente, reutilizando")
		return conn, nil
	}

	// Verificar límite de conexiones
	if len(p.connections) >= p.config.PeerPool.MaxPeers {
		return nil, fmt.Errorf("límite de peers alcanzado (%d)", p.config.PeerPool.MaxPeers)
	}

	// Preparar certificados para TLS
	certFile := p.config.PeerPool.TLS.CertFile
	keyFile := p.config.PeerPool.TLS.KeyFile
	caFile := p.config.PeerPool.TLS.CAFile
	serverName := p.config.PeerPool.TLS.ServerName

	// Crear configuración TLS
	tlsConfig, err := createClientTLSConfig(certFile, keyFile, caFile, serverName)
	if err != nil {
		return nil, fmt.Errorf("error configurando TLS: %w", err)
	}

	// Establecer conexión
	address := fmt.Sprintf("%s:%d", peer.Address, peer.Port)
	
	p.log.WithFields(logrus.Fields{
		"peer_id":  peer.ID,
		"address":  address,
		"node_name": peer.NodeName,
	}).Info("Conectando a peer")
	
	// Crear un dialer con timeout
	dialer := &net.Dialer{
		Timeout: p.dialTimeout,
	}
	
	// Conectar usando TLS
	ctx, cancel := context.WithTimeout(context.Background(), p.handshakeTimeout)
	defer cancel()
	
	conn, err := tls.DialWithDialer(dialer, "tcp", address, tlsConfig)
	if err != nil {
		p.log.WithFields(logrus.Fields{
			"peer_id": peer.ID,
			"address": address,
			"error":   err.Error(),
		}).Error("Error conectando a peer")
		return nil, fmt.Errorf("error conectando a peer: %w", err)
	}
	
	// Realizar handshake TLS con timeout
	if err := conn.HandshakeContext(ctx); err != nil {
		conn.Close()
		p.log.WithFields(logrus.Fields{
			"peer_id": peer.ID,
			"address": address,
			"error":   err.Error(),
		}).Error("Error en handshake TLS")
		return nil, fmt.Errorf("error en handshake TLS: %w", err)
	}

	// Crear la estructura PeerConn
	peerConn := NewPeerConn(
		peer.ID,
		peer,
		conn,
		p.config.PeerPool.BufferSize,
		p.config.PeerPool.MaxFrameSize,
		p.log,
	)
	
	// Establecer callback para notificar cambios de estado
	peerConn.SetOnStateChange(func(id uuid.UUID, state PeerState) {
		if p.peerStateNotifier != nil {
			p.peerStateNotifier(id, state)
		}
	})
	
	// Guardarla en el mapa
	p.connections[peer.ID] = peerConn

	p.log.WithFields(logrus.Fields{
		"peer_id":  peer.ID,
		"address":  address,
		"node_name": peer.NodeName,
	}).Info("Peer conectado exitosamente")

	// Iniciar monitoreo de esta conexión
	go p.monitorConnection(peer.ID)

	return peerConn, nil
}

// Get obtiene una conexión existente
func (p *PeerConnectionPool) Get(peerID uuid.UUID) (*PeerConn, bool) {
	p.mu.RLock()
	conn, exists := p.connections[peerID]
	p.mu.RUnlock()
	
	if exists && conn.state == PeerStateConnected {
		return conn, true
	}
	
	return nil, false
}

// Close cierra una conexión específica
func (p *PeerConnectionPool) Close(peerID uuid.UUID) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	conn, exists := p.connections[peerID]
	if !exists {
		return nil // No existe, no hay error
	}
	
	// Cerrar la conexión
	err := conn.Close()
	
	// Eliminar del mapa
	delete(p.connections, peerID)
	
	p.log.WithField("peer_id", peerID).Info("Conexión con peer cerrada")
	
	return err
}

// CloseAll cierra todas las conexiones
func (p *PeerConnectionPool) CloseAll() {
	p.mu.Lock()
	
	// Copiar las claves para evitar modificar el mapa durante la iteración
	var keys []uuid.UUID
	for k := range p.connections {
		keys = append(keys, k)
	}
	
	p.mu.Unlock()
	
	// Cerrar cada conexión
	for _, id := range keys {
		_ = p.Close(id)
	}
	
	p.log.Info("Todas las conexiones con peers cerradas")
}

// monitorConnection verifica el estado de una conexión y maneja reconexiones
func (p *PeerConnectionPool) monitorConnection(peerID uuid.UUID) {
	attempt := 0
	
	for {
		// Esperar a que la conexión se desconecte
		p.mu.RLock()
		conn, exists := p.connections[peerID]
		p.mu.RUnlock()
		
		if !exists {
			// La conexión ya no existe, salir de la goroutine
			return
		}
		
		// Si está conectado, seguir monitoreando
		if conn.state == PeerStateConnected {
			time.Sleep(time.Second)
			continue
		}
		
		// Si ya pasamos el número máximo de intentos, marcar como definitivamente desconectado
		if attempt >= p.maxAttempts {
			p.mu.Lock()
			if conn, exists := p.connections[peerID]; exists {
				conn.setState(PeerStateDisconnected)
			}
			p.mu.Unlock()
			
			p.log.WithFields(logrus.Fields{
				"peer_id": peerID,
				"attempts": attempt,
			}).Warn("Máximo número de intentos de reconexión alcanzado, desistiendo")
			
			return
		}
		
		// Calcular retraso con backoff exponencial y jitter
		delay := p.calculateBackoff(attempt)
		
		// Marcar como reconectando
		p.mu.Lock()
		if conn, exists := p.connections[peerID]; exists {
			conn.setState(PeerStateReconnecting)
		}
		p.mu.Unlock()
		
		p.log.WithFields(logrus.Fields{
			"peer_id": peerID,
			"attempt": attempt,
			"delay_ms": delay.Milliseconds(),
		}).Info("Esperando para reintentar conexión")
		
		// Esperar antes de reintentar
		time.Sleep(delay)
		
		// Intentar reconectar
		p.mu.RLock()
		peerInfo := conn.PeerInfo
		p.mu.RUnlock()
		
		// Realizar reconexión como una nueva conexión (reemplazando la antigua)
		_, err := p.DialAndRegister(peerInfo)
		if err != nil {
			p.log.WithFields(logrus.Fields{
				"peer_id": peerID,
				"attempt": attempt,
				"error": err.Error(),
			}).Error("Falló el intento de reconexión")
			
			attempt++
			continue
		}
		
		// Si llegamos aquí, la reconexión fue exitosa
		p.log.WithFields(logrus.Fields{
			"peer_id": peerID,
			"attempts": attempt,
		}).Info("Reconexión exitosa")
		
		return
	}
}

// calculateBackoff calcula el tiempo de espera para reconexión con backoff exponencial y jitter
func (p *PeerConnectionPool) calculateBackoff(attempt int) time.Duration {
	// Fórmula de backoff exponencial: baseDelay * 2^attempt
	backoff := float64(p.baseDelay) * math.Pow(2, float64(attempt))
	
	// Aplicar límite máximo
	if backoff > float64(p.maxDelay) {
		backoff = float64(p.maxDelay)
	}
	
	// Aplicar jitter para evitar tormentas de reconexión
	jitter := backoff * p.jitterFactor
	backoff = backoff + rand.Float64()*jitter*2 - jitter
	
	return time.Duration(backoff)
}

// keepaliveLoop envía mensajes keepalive a todos los peers conectados
func (p *PeerConnectionPool) keepaliveLoop() {
	ticker := time.NewTicker(p.keepaliveInterval)
	defer ticker.Stop()
	
	for range ticker.C {
		p.mu.RLock()
		peers := make([]uuid.UUID, 0, len(p.connections))
		for id, conn := range p.connections {
			if conn.state == PeerStateConnected {
				peers = append(peers, id)
			}
		}
		p.mu.RUnlock()
		
		for _, id := range peers {
			conn, exists := p.Get(id)
			if !exists {
				continue
			}
			
			// Enviar keepalive
			if err := conn.SendFrame(FrameTypeKeepAlive, nil); err != nil {
				p.log.WithFields(logrus.Fields{
					"peer_id": id,
					"error": err.Error(),
				}).Error("Error enviando keepalive")
				
				// Si hay error de escritura, cerrar la conexión para que se reconecte
				go p.Close(id)
			}
		}
	}
}

// GetAllPeerIDs devuelve los IDs de todos los peers conectados
func (p *PeerConnectionPool) GetAllPeerIDs() []uuid.UUID {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	result := make([]uuid.UUID, 0, len(p.connections))
	for id, conn := range p.connections {
		if conn.state == PeerStateConnected {
			result = append(result, id)
		}
	}
	
	return result
}

// GetMetrics devuelve métricas del pool y de cada conexión
func (p *PeerConnectionPool) GetMetrics() map[string]interface{} {
	p.mu.RLock()
	defer p.mu.RUnlock()
	
	// Métricas generales del pool
	metrics := map[string]interface{}{
		"total_peers": len(p.connections),
		"max_peers": p.config.PeerPool.MaxPeers,
		"utilization_pct": float64(len(p.connections)) / float64(p.config.PeerPool.MaxPeers) * 100.0,
	}
	
	// Contar por estado
	connected := 0
	reconnecting := 0
	disconnected := 0
	
	for _, conn := range p.connections {
		switch conn.state {
		case PeerStateConnected:
			connected++
		case PeerStateReconnecting:
			reconnecting++
		case PeerStateDisconnected:
			disconnected++
		}
	}
	
	metrics["connected_peers"] = connected
	metrics["reconnecting_peers"] = reconnecting
	metrics["disconnected_peers"] = disconnected
	
	// Extraer métricas detalladas por peer
	peerMetrics := make(map[string]interface{})
	for id, conn := range p.connections {
		peerMetrics[id.String()] = conn.GetMetrics()
	}
	
	metrics["peers"] = peerMetrics
	
	return metrics
}
