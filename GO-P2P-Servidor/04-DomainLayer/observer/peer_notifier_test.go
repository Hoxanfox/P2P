package observer

import (
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"model"
)

// MockPeerObserver implementa la interfaz IPeerObserver para tests
type MockPeerObserver struct {
	connectedCalled      bool
	disconnectedCalled   bool
	heartbeatMissedCalled bool
	
	lastPeer *model.Peer
}

func NewMockPeerObserver() *MockPeerObserver {
	return &MockPeerObserver{}
}

func (m *MockPeerObserver) OnPeerConnected(peer *model.Peer) {
	m.connectedCalled = true
	m.lastPeer = peer
}

func (m *MockPeerObserver) OnPeerDisconnected(peer *model.Peer) {
	m.disconnectedCalled = true
	m.lastPeer = peer
}

func (m *MockPeerObserver) OnPeerHeartbeatMissed(peer *model.Peer) {
	m.heartbeatMissedCalled = true
	m.lastPeer = peer
}

// Helper para crear un peer válido para tests
func createTestPeer() *model.Peer {
	id := uuid.New()
	// Crear una dirección IP y puerto válido para el test
	addr := net.JoinHostPort("192.168.1.100", "8080")
	peer, _ := model.NewPeer(
		id,
		addr,
		model.NodoConectado,
	)
	return peer
}

// Helper para crear un hearbeat_log válido para tests
func createTestHeartbeatLog(peerID uuid.UUID) *model.HeartbeatLog {
	id := uuid.New()
	now := time.Now().UTC()
	// El recibido es 100ms después del enviado
	received := now.Add(100 * time.Millisecond)
	
	log, _ := model.NewHeartbeatLog(
		id,
		peerID,
		now,
		received,
	)
	return log
}

func TestPeerNotifier_Subscribe(t *testing.T) {
	notifier := NewPeerNotifier()
	observer := NewMockPeerObserver()
	
	// Verificar que inicialmente no hay observadores
	if count := notifier.ObserversCount(); count != 0 {
		t.Errorf("Se esperaba 0 observadores, se obtuvo %d", count)
	}
	
	// Suscribir un observador
	notifier.Subscribe(observer)
	
	// Verificar que ahora hay un observador
	if count := notifier.ObserversCount(); count != 1 {
		t.Errorf("Se esperaba 1 observador, se obtuvo %d", count)
	}
	
	// Suscribir el mismo observador nuevamente (no debería duplicarse)
	notifier.Subscribe(observer)
	
	// Verificar que sigue habiendo un observador
	if count := notifier.ObserversCount(); count != 1 {
		t.Errorf("Se esperaba 1 observador, se obtuvo %d", count)
	}
}

func TestPeerNotifier_Unsubscribe(t *testing.T) {
	notifier := NewPeerNotifier()
	observer1 := NewMockPeerObserver()
	observer2 := NewMockPeerObserver()
	
	// Suscribir dos observadores
	notifier.Subscribe(observer1)
	notifier.Subscribe(observer2)
	
	// Verificar que hay dos observadores
	if count := notifier.ObserversCount(); count != 2 {
		t.Errorf("Se esperaba 2 observadores, se obtuvo %d", count)
	}
	
	// Desuscribir un observador
	notifier.Unsubscribe(observer1)
	
	// Verificar que queda un observador
	if count := notifier.ObserversCount(); count != 1 {
		t.Errorf("Se esperaba 1 observador, se obtuvo %d", count)
	}
}

func TestPeerNotifier_NotifyPeerConnected(t *testing.T) {
	notifier := NewPeerNotifier()
	observer := NewMockPeerObserver()
	notifier.Subscribe(observer)
	
	peer := createTestPeer()
	notifier.NotifyPeerConnected(peer)
	
	// Verificar que se llamó al método correcto del observador
	if !observer.connectedCalled {
		t.Error("OnPeerConnected no fue llamado")
	}
	
	// Verificar que se pasó el peer correcto
	if observer.lastPeer != peer {
		t.Error("El peer pasado al observador no es el esperado")
	}
}

func TestPeerNotifier_NotifyPeerDisconnected(t *testing.T) {
	notifier := NewPeerNotifier()
	observer := NewMockPeerObserver()
	notifier.Subscribe(observer)
	
	peer := createTestPeer()
	notifier.NotifyPeerDisconnected(peer)
	
	// Verificar que se llamó al método correcto del observador
	if !observer.disconnectedCalled {
		t.Error("OnPeerDisconnected no fue llamado")
	}
	
	// Verificar que se pasó el peer correcto
	if observer.lastPeer != peer {
		t.Error("El peer pasado al observador no es el esperado")
	}
}

func TestPeerNotifier_NotifyPeerHeartbeatMissed(t *testing.T) {
	notifier := NewPeerNotifier()
	observer := NewMockPeerObserver()
	notifier.Subscribe(observer)
	
	peer := createTestPeer()
	notifier.NotifyPeerHeartbeatMissed(peer)
	
	// Verificar que se llamó al método correcto del observador
	if !observer.heartbeatMissedCalled {
		t.Error("OnPeerHeartbeatMissed no fue llamado")
	}
	
	// Verificar que se pasó el peer correcto
	if observer.lastPeer != peer {
		t.Error("El peer pasado al observador no es el esperado")
	}
}
