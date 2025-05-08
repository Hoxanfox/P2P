package pool

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockPeerConn implementa una conexión mock para pruebas
type MockPeerConn struct {
	net.Conn
	closed bool
	data   []byte
}

func NewMockPeerConn() *MockPeerConn {
	return &MockPeerConn{
		data: make([]byte, 0),
	}
}

func (m *MockPeerConn) Read(b []byte) (n int, err error) {
	if m.closed {
		return 0, fmt.Errorf("conexión cerrada")
	}
	return 0, nil
}

func (m *MockPeerConn) Write(b []byte) (n int, err error) {
	if m.closed {
		return 0, fmt.Errorf("conexión cerrada")
	}
	m.data = append(m.data, b...)
	return len(b), nil
}

func (m *MockPeerConn) Close() error {
	m.closed = true
	return nil
}

func (m *MockPeerConn) LocalAddr() net.Addr {
	return &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9000}
}

func (m *MockPeerConn) RemoteAddr() net.Addr {
	return &net.TCPAddr{IP: net.ParseIP("192.168.0.1"), Port: 9001}
}

func (m *MockPeerConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *MockPeerConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *MockPeerConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestLoadPeerPoolConfig(t *testing.T) {
	// Verificar que podemos cargar la configuración desde el archivo YAML
	config, err := LoadPeerPoolConfig("peer_config.yaml")
	require.NoError(t, err, "Debe poder cargar la configuración sin errores")
	
	// Validar valores cargados
	assert.Equal(t, 100, config.PeerPool.MaxPeers)
	assert.Equal(t, "10s", config.PeerPool.DialTimeout)
	assert.Equal(t, "5s", config.PeerPool.HandshakeTimeout)
	
	// Verificar configuración de reconexión
	assert.Equal(t, "1s", config.PeerPool.Reconnect.BaseDelay)
	assert.Equal(t, "60s", config.PeerPool.Reconnect.MaxDelay)
	assert.Equal(t, 10, config.PeerPool.Reconnect.MaxAttempts)
	assert.Equal(t, 0.2, config.PeerPool.Reconnect.JitterFactor)
	
	// Verificar configuración TLS
	assert.Equal(t, "cert/node-cert.pem", config.PeerPool.TLS.CertFile)
	assert.Equal(t, "cert/node-key.pem", config.PeerPool.TLS.KeyFile)
	assert.Equal(t, "cert/ca-cert.pem", config.PeerPool.TLS.CAFile)
}

func TestPeerConnBasicMethods(t *testing.T) {
	// Crear una conexión de peer con un mock
	mockConn := NewMockPeerConn()
	peerID := uuid.New()
	peerInfo := PeerInfo{
		ID:       peerID,
		Address:  "192.168.0.1",
		Port:     9001,
		NodeName: "peer1",
	}
	
	// Crear logger de prueba silencioso
	logger := logrus.New()
	logger.SetOutput(io.Discard)
	
	// Crear PeerConn
	peerConn := NewPeerConn(peerID, peerInfo, mockConn, 4096, 1048576, logger)
	
	// Verificar estado inicial
	assert.Equal(t, PeerStateConnected, peerConn.state)
	
	// Verificar cambio de estado
	var notifiedID uuid.UUID
	var notifiedState PeerState
	
	peerConn.SetOnStateChange(func(id uuid.UUID, state PeerState) {
		notifiedID = id
		notifiedState = state
	})
	
	peerConn.setState(PeerStateDisconnected)
	assert.Equal(t, peerID, notifiedID)
	assert.Equal(t, PeerStateDisconnected, notifiedState)
	
	// Verificar cierre
	err := peerConn.Close()
	assert.NoError(t, err)
	assert.True(t, mockConn.closed)
	
	// Verificar que no se puede enviar después de cerrar
	err = peerConn.SendFrame(FrameTypeData, []byte("test"))
	assert.Error(t, err)
}

func TestPeerFrameOperations(t *testing.T) {
	// Esta prueba solamente se puede realizar de manera simplificada
	// ya que requeriría una conexión real o un mock más sofisticado
	
	// Verificar tipos de frames
	assert.Equal(t, uint16(0x0001), FrameTypeData)
	assert.Equal(t, uint16(0x0002), FrameTypeKeepAlive)
	assert.Equal(t, uint16(0x0003), FrameTypeACK)
	assert.Equal(t, uint16(0x0004), FrameTypeNACK)
	assert.Equal(t, uint16(0x0005), FrameTypeClose)
}

func TestPeerConnectionPoolOperations(t *testing.T) {
	// Esta prueba es principalmente para validar compilación y estructura
	// ya que una prueba real requeriría certificados TLS y conexiones reales
	
	t.Skip("Esta prueba requiere certificados TLS y conexiones reales")
	
	// En un entorno real, se podría usar algo como:
	/*
	pool, err := NewPeerConnectionPool("peer_config.yaml")
	require.NoError(t, err)
	
	peer := PeerInfo{
		ID:       uuid.New(),
		Address:  "localhost",
		Port:     9001,
		NodeName: "test-peer",
	}
	
	conn, err := pool.DialAndRegister(peer)
	require.NoError(t, err)
	
	retrievedConn, exists := pool.Get(peer.ID)
	assert.True(t, exists)
	assert.Equal(t, conn, retrievedConn)
	
	// Enviar y recibir frames
	err = conn.SendFrame(FrameTypeData, []byte("hello"))
	require.NoError(t, err)
	
	// Cerrar conexión
	err = pool.Close(peer.ID)
	require.NoError(t, err)
	
	// Verificar que ya no existe
	_, exists = pool.Get(peer.ID)
	assert.False(t, exists)
	
	// Cerrar todo
	pool.CloseAll()
	*/
}
