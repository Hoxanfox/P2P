package pool

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// PeerState representa el estado actual de un peer
type PeerState string

const (
	PeerStateConnected    PeerState = "CONNECTED"
	PeerStateDisconnected PeerState = "DISCONNECTED"
	PeerStateReconnecting PeerState = "RECONNECTING"
)

// PeerInfo contiene la información necesaria para conectar con un peer
type PeerInfo struct {
	ID        uuid.UUID
	Address   string
	Port      int
	NodeName  string
	PublicKey string
}

// String devuelve una representación en cadena del peer
func (p PeerInfo) String() string {
	return fmt.Sprintf("%s (%s:%d)", p.NodeName, p.Address, p.Port)
}

// PeerFrame representa un mensaje encapsulado entre nodos P2P
type PeerFrame struct {
	Type    uint16
	Payload []byte
}

const (
	// Tipos de frames
	FrameTypeData      uint16 = 0x0001
	FrameTypeKeepAlive uint16 = 0x0002
	FrameTypeACK       uint16 = 0x0003
	FrameTypeNACK      uint16 = 0x0004
	FrameTypeClose     uint16 = 0x0005
)

// frameHeader es el tamaño del encabezado de un frame
const frameHeader = 6 // 2 bytes de tipo + 4 bytes de longitud

// PeerConn encapsula una conexión con otro nodo P2P
type PeerConn struct {
	ID             uuid.UUID
	PeerInfo       PeerInfo
	conn           net.Conn
	tlsConn        *tls.Conn
	state          PeerState
	lastActivity   time.Time
	metrics        *PeerMetrics
	sendMutex      sync.Mutex
	recvMutex      sync.Mutex
	log            *logrus.Logger
	bufferSize     int
	maxFrameSize   int
	onStateChange  func(uuid.UUID, PeerState)
	ctx            context.Context
	cancel         context.CancelFunc
	closed         int32
}

// PeerMetrics almacena métricas para monitorear la conexión con un peer
type PeerMetrics struct {
	BytesSent        int64
	BytesReceived    int64
	MessagesSent     int64
	MessagesReceived int64
	Retries          int64
	RTTSum           int64
	RTTCount         int64
	Errors           int64
	LastError        error
	mu               sync.RWMutex
}

// NewPeerConn crea una nueva conexión de peer
func NewPeerConn(id uuid.UUID, peerInfo PeerInfo, conn net.Conn, bufferSize, maxFrameSize int, logger *logrus.Logger) *PeerConn {
	ctx, cancel := context.WithCancel(context.Background())
	
	var tlsConn *tls.Conn
	if tc, ok := conn.(*tls.Conn); ok {
		tlsConn = tc
	}
	
	return &PeerConn{
		ID:           id,
		PeerInfo:     peerInfo,
		conn:         conn,
		tlsConn:      tlsConn,
		state:        PeerStateConnected,
		lastActivity: time.Now(),
		metrics: &PeerMetrics{
			mu: sync.RWMutex{},
		},
		log:          logger,
		bufferSize:   bufferSize,
		maxFrameSize: maxFrameSize,
		ctx:          ctx,
		cancel:       cancel,
	}
}

// SetOnStateChange establece una función callback para cambios de estado
func (p *PeerConn) SetOnStateChange(callback func(uuid.UUID, PeerState)) {
	p.onStateChange = callback
}

// setState cambia el estado del peer y llama al callback si está configurado
func (p *PeerConn) setState(newState PeerState) {
	p.state = newState
	if p.onStateChange != nil {
		p.onStateChange(p.ID, newState)
	}
}

// SendFrame envía un frame al peer
func (p *PeerConn) SendFrame(frameType uint16, payload []byte) error {
	if atomic.LoadInt32(&p.closed) != 0 {
		return fmt.Errorf("conexión cerrada")
	}
	
	if len(payload) > p.maxFrameSize {
		return fmt.Errorf("tamaño de payload excede el máximo permitido: %d > %d", len(payload), p.maxFrameSize)
	}
	
	// Proteger el envío con mutex para evitar escrituras concurrentes
	p.sendMutex.Lock()
	defer p.sendMutex.Unlock()
	
	// Medir tiempo para RTT si es un mensaje de datos
	startTime := time.Now()
	
	// Crear buffer para el frame completo
	frameSize := frameHeader + len(payload)
	frame := make([]byte, frameSize)
	
	// Escribir tipo y longitud
	binary.BigEndian.PutUint16(frame[0:2], frameType)
	binary.BigEndian.PutUint32(frame[2:6], uint32(len(payload)))
	
	// Copiar payload
	copy(frame[frameHeader:], payload)
	
	// Enviar frame
	deadline := time.Now().Add(10 * time.Second)
	if err := p.conn.SetWriteDeadline(deadline); err != nil {
		p.metrics.mu.Lock()
		p.metrics.Errors++
		p.metrics.LastError = err
		p.metrics.mu.Unlock()
		return err
	}
	
	n, err := p.conn.Write(frame)
	if err != nil {
		p.metrics.mu.Lock()
		p.metrics.Errors++
		p.metrics.LastError = err
		p.metrics.mu.Unlock()
		return err
	}
	
	// Actualizar métricas
	atomic.AddInt64(&p.metrics.BytesSent, int64(n))
	atomic.AddInt64(&p.metrics.MessagesSent, 1)
	
	// Actualizar RTT solo para mensajes de datos
	if frameType == FrameTypeData {
		rtt := time.Since(startTime)
		atomic.AddInt64(&p.metrics.RTTSum, int64(rtt))
		atomic.AddInt64(&p.metrics.RTTCount, 1)
	}
	
	// Actualizar timestamp de actividad
	p.lastActivity = time.Now()
	
	return nil
}

// ReceiveFrame recibe un frame del peer
func (p *PeerConn) ReceiveFrame() (*PeerFrame, error) {
	if atomic.LoadInt32(&p.closed) != 0 {
		return nil, fmt.Errorf("conexión cerrada")
	}
	
	// Proteger la recepción con mutex
	p.recvMutex.Lock()
	defer p.recvMutex.Unlock()
	
	// Leer encabezado
	headerBuf := make([]byte, frameHeader)
	deadline := time.Now().Add(10 * time.Second)
	if err := p.conn.SetReadDeadline(deadline); err != nil {
		p.metrics.mu.Lock()
		p.metrics.Errors++
		p.metrics.LastError = err
		p.metrics.mu.Unlock()
		return nil, err
	}
	
	if _, err := io.ReadFull(p.conn, headerBuf); err != nil {
		p.metrics.mu.Lock()
		p.metrics.Errors++
		p.metrics.LastError = err
		p.metrics.mu.Unlock()
		return nil, err
	}
	
	// Decodificar encabezado
	frameType := binary.BigEndian.Uint16(headerBuf[0:2])
	payloadLen := binary.BigEndian.Uint32(headerBuf[2:6])
	
	// Verificar tamaño máximo
	if payloadLen > uint32(p.maxFrameSize) {
		err := fmt.Errorf("tamaño de frame recibido excede el máximo: %d > %d", payloadLen, p.maxFrameSize)
		p.metrics.mu.Lock()
		p.metrics.Errors++
		p.metrics.LastError = err
		p.metrics.mu.Unlock()
		return nil, err
	}
	
	// Leer payload
	payload := make([]byte, payloadLen)
	if payloadLen > 0 {
		if _, err := io.ReadFull(p.conn, payload); err != nil {
			p.metrics.mu.Lock()
			p.metrics.Errors++
			p.metrics.LastError = err
			p.metrics.mu.Unlock()
			return nil, err
		}
	}
	
	// Actualizar métricas
	readSize := int64(frameHeader) + int64(payloadLen)
	atomic.AddInt64(&p.metrics.BytesReceived, readSize)
	atomic.AddInt64(&p.metrics.MessagesReceived, 1)
	
	// Actualizar timestamp de actividad
	p.lastActivity = time.Now()
	
	return &PeerFrame{
		Type:    frameType,
		Payload: payload,
	}, nil
}

// Close cierra la conexión con el peer
func (p *PeerConn) Close() error {
	if !atomic.CompareAndSwapInt32(&p.closed, 0, 1) {
		return nil // Ya está cerrado
	}
	
	// Cancelar cualquier goroutine asociada
	p.cancel()
	
	// Intentar enviar frame de cierre (no importa si falla)
	_ = p.SendFrame(FrameTypeClose, nil)
	
	// Cambiar estado
	p.setState(PeerStateDisconnected)
	
	return p.conn.Close()
}

// AvgRTT devuelve el tiempo promedio de ida y vuelta (RTT)
func (p *PeerConn) AvgRTT() time.Duration {
	sum := atomic.LoadInt64(&p.metrics.RTTSum)
	count := atomic.LoadInt64(&p.metrics.RTTCount)
	
	if count == 0 {
		return 0
	}
	
	return time.Duration(sum / count)
}

// GetMetrics devuelve las métricas actualizadas de la conexión
func (p *PeerConn) GetMetrics() map[string]interface{} {
	metrics := make(map[string]interface{})
	
	// Capturar valores atómicamente
	metrics["bytes_sent"] = atomic.LoadInt64(&p.metrics.BytesSent)
	metrics["bytes_received"] = atomic.LoadInt64(&p.metrics.BytesReceived)
	metrics["messages_sent"] = atomic.LoadInt64(&p.metrics.MessagesSent)
	metrics["messages_received"] = atomic.LoadInt64(&p.metrics.MessagesReceived)
	metrics["retries"] = atomic.LoadInt64(&p.metrics.Retries)
	
	// Calcular RTT
	rttSum := atomic.LoadInt64(&p.metrics.RTTSum)
	rttCount := atomic.LoadInt64(&p.metrics.RTTCount)
	if rttCount > 0 {
		metrics["avg_rtt_ns"] = rttSum / rttCount
	} else {
		metrics["avg_rtt_ns"] = 0
	}
	
	// Añadir errores
	p.metrics.mu.RLock()
	metrics["errors"] = p.metrics.Errors
	if p.metrics.LastError != nil {
		metrics["last_error"] = p.metrics.LastError.Error()
	}
	p.metrics.mu.RUnlock()
	
	// Añadir información de estado
	metrics["state"] = string(p.state)
	metrics["last_activity"] = p.lastActivity.Format(time.RFC3339)
	metrics["idle_time_sec"] = time.Since(p.lastActivity).Seconds()
	
	return metrics
}

// createClientTLSConfig crea la configuración TLS para conectar a un peer
func createClientTLSConfig(certFile, keyFile, caFile, serverName string) (*tls.Config, error) {
	// Cargar certificado de cliente
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("error cargando certificado de cliente: %w", err)
	}
	
	// Crear pool de certificados para validar al servidor
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("error cargando certificado CA: %w", err)
	}
	
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("error añadiendo certificado CA al pool")
	}
	
	// Crear configuración TLS
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		ServerName:   serverName,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		MinVersion:   tls.VersionTLS12,
	}, nil
}

// createServerTLSConfig crea la configuración TLS para aceptar conexiones de peers
func createServerTLSConfig(certFile, keyFile, caFile string) (*tls.Config, error) {
	// Cargar certificado de servidor
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("error cargando certificado de servidor: %w", err)
	}
	
	// Crear pool de certificados para validar clientes
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("error cargando certificado CA: %w", err)
	}
	
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("error añadiendo certificado CA al pool")
	}
	
	// Crear configuración TLS
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		MinVersion:   tls.VersionTLS12,
	}, nil
}
