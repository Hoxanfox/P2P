package pool

import (
	"bytes"
	"io"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// MockConn implementa la interfaz net.Conn para pruebas
type MockConn struct {
	ReadBuffer  bytes.Buffer
	WriteBuffer bytes.Buffer
	IsClosed    bool
	RemoteAddress net.Addr
}

func (m *MockConn) Read(b []byte) (n int, err error) {
	if m.IsClosed {
		return 0, io.EOF
	}
	return m.ReadBuffer.Read(b)
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	if m.IsClosed {
		return 0, io.EOF
	}
	return m.WriteBuffer.Write(b)
}

func (m *MockConn) Close() error {
	m.IsClosed = true
	return nil
}

func (m *MockConn) LocalAddr() net.Addr {
	return &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080}
}

func (m *MockConn) RemoteAddr() net.Addr {
	if m.RemoteAddress != nil {
		return m.RemoteAddress
	}
	return &net.TCPAddr{IP: net.ParseIP("192.168.1.1"), Port: 45678}
}

func (m *MockConn) SetDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestLoadSocketConfig(t *testing.T) {
	// Verificar que podemos cargar la configuración desde el archivo YAML
	config, err := LoadSocketConfig("socket_config.yaml")
	require.NoError(t, err, "Debe poder cargar la configuración sin errores")
	
	// Validar valores cargados
	assert.Equal(t, 1000, config.SocketPool.MaxConnections)
	assert.Equal(t, 300, config.SocketPool.InactiveTimeout)
	assert.Equal(t, 60, config.SocketPool.HealthCheckInterval)
	assert.Equal(t, 4096, config.SocketPool.BufferSize)
	assert.Equal(t, int64(5000), config.SocketPool.WriteTimeout)
}

func TestSocketPoolBasicOperations(t *testing.T) {
	// Crear un pool de sockets para pruebas
	pool, err := NewSocketPool("socket_config.yaml")
	require.NoError(t, err, "Debe crear el pool sin errores")
	defer pool.Close()
	
	// Crear un cliente de prueba
	clientID := uuid.New()
	conn := &MockConn{}
	
	// Registrar el cliente
	err = pool.Register(clientID, conn)
	assert.NoError(t, err, "Debe registrar el cliente sin errores")
	
	// Verificar que podemos recuperar la conexión
	retrievedConn, exists := pool.Get(clientID)
	assert.True(t, exists, "La conexión debe existir en el pool")
	assert.Equal(t, conn, retrievedConn, "Debe devolver la misma conexión")
	
	// Verificar métricas
	metrics := pool.GetMetrics()
	assert.Equal(t, 1, metrics["active_connections"].(int), "Debe tener 1 conexión activa")
	
	// Liberar la conexión
	pool.Release(clientID)
	
	// Verificar que ya no existe
	_, exists = pool.Get(clientID)
	assert.False(t, exists, "La conexión no debe existir después de liberarla")
	assert.True(t, conn.IsClosed, "La conexión debe estar cerrada")
}

func TestSocketPoolBroadcast(t *testing.T) {
	// Crear un pool de sockets para pruebas
	pool, err := NewSocketPool("socket_config.yaml")
	require.NoError(t, err, "Debe crear el pool sin errores")
	defer pool.Close()
	
	// Crear varios clientes de prueba
	clientIDs := []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}
	conns := make([]*MockConn, len(clientIDs))
	
	// Registrar los clientes
	for i, id := range clientIDs {
		conns[i] = &MockConn{}
		err = pool.Register(id, conns[i])
		assert.NoError(t, err, "Debe registrar el cliente sin errores")
	}
	
	// Enviar un mensaje de broadcast
	message := []byte("Este es un mensaje de prueba")
	err = pool.Broadcast(clientIDs, message)
	assert.NoError(t, err, "El broadcast debe funcionar sin errores")
	
	// Verificar que todos recibieron el mensaje
	for _, conn := range conns {
		assert.Equal(t, message, conn.WriteBuffer.Bytes(), "Todos los clientes deben recibir el mensaje")
	}
	
	// Probar broadcast con un ID inválido
	invalidIDs := append(clientIDs, uuid.New())
	err = pool.Broadcast(invalidIDs, message)
	assert.Error(t, err, "Debe dar error al hacer broadcast a un cliente inexistente")
}

func TestSocketPoolMaxConnections(t *testing.T) {
	// Crear un pool de sockets con configuración personalizada para la prueba
	// Modificamos la configuración para un límite más bajo
	config := &SocketConfig{}
	config.SocketPool.MaxConnections = 2 // Límite bajo para facilitar la prueba
	
	// Crear un logger de prueba que no haga nada
	logger := logrus.New()
	logger.SetOutput(io.Discard) // No queremos la salida del logger en las pruebas
	
	pool := &SocketPool{
		connections: make(map[uuid.UUID]*ClientConnection),
		config:      config,
		log:         logger,
	}
	
	// Registrar clientes hasta llegar al límite
	clientID1 := uuid.New()
	clientID2 := uuid.New()
	clientID3 := uuid.New()
	
	conn1 := &MockConn{}
	conn2 := &MockConn{}
	conn3 := &MockConn{}
	
	// Las primeras 2 conexiones deben aceptarse
	err := pool.Register(clientID1, conn1)
	assert.NoError(t, err, "Debe registrar la primera conexión sin errores")
	
	err = pool.Register(clientID2, conn2)
	assert.NoError(t, err, "Debe registrar la segunda conexión sin errores")
	
	// La tercera conexión debe ser rechazada
	err = pool.Register(clientID3, conn3)
	assert.Error(t, err, "Debe rechazar conexiones por encima del límite")
	assert.Contains(t, err.Error(), "se alcanzó el límite máximo de conexiones", "Debe indicar el motivo del rechazo")
}

func TestSocketPoolConnOverride(t *testing.T) {
	// Crear un pool de sockets para pruebas
	pool, err := NewSocketPool("socket_config.yaml")
	require.NoError(t, err, "Debe crear el pool sin errores")
	defer pool.Close()
	
	// Crear un cliente de prueba
	clientID := uuid.New()
	conn1 := &MockConn{}
	
	// Registrar el cliente
	err = pool.Register(clientID, conn1)
	assert.NoError(t, err, "Debe registrar el cliente sin errores")
	
	// Registrar otro con el mismo ID (debe reemplazar al anterior)
	conn2 := &MockConn{}
	err = pool.Register(clientID, conn2)
	assert.NoError(t, err, "Debe registrar el cliente sin errores")
	
	// Verificar que obtenemos la nueva conexión
	retrievedConn, exists := pool.Get(clientID)
	assert.True(t, exists, "La conexión debe existir en el pool")
	assert.Equal(t, conn2, retrievedConn, "Debe devolver la nueva conexión")
	
	// La conexión anterior debe estar cerrada
	assert.True(t, conn1.IsClosed, "La conexión anterior debe estar cerrada")
}
