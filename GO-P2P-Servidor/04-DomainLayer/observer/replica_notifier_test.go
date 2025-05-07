package observer

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"model"
)

// MockReplicaObserver implementa la interfaz IReplicaObserver para tests
type MockReplicaObserver struct {
	userReplicatedCalled     bool
	channelReplicatedCalled  bool
	messageReplicatedCalled  bool
	fileReplicatedCalled     bool
	
	lastUser    *model.UsuarioServidor
	lastChannel *model.CanalServidor
	lastMessage *model.MensajeServidor
	lastFile    *model.ArchivoMetadata
	lastPeer    *model.Peer
}

func NewMockReplicaObserver() *MockReplicaObserver {
	return &MockReplicaObserver{}
}

func (m *MockReplicaObserver) OnUserReplicated(user *model.UsuarioServidor, toPeer *model.Peer) {
	m.userReplicatedCalled = true
	m.lastUser = user
	m.lastPeer = toPeer
}

func (m *MockReplicaObserver) OnChannelReplicated(channel *model.CanalServidor, toPeer *model.Peer) {
	m.channelReplicatedCalled = true
	m.lastChannel = channel
	m.lastPeer = toPeer
}

func (m *MockReplicaObserver) OnMessageReplicated(msg *model.MensajeServidor, toPeer *model.Peer) {
	m.messageReplicatedCalled = true
	m.lastMessage = msg
	m.lastPeer = toPeer
}

func (m *MockReplicaObserver) OnFileReplicated(file *model.ArchivoMetadata, toPeer *model.Peer) {
	m.fileReplicatedCalled = true
	m.lastFile = file
	m.lastPeer = toPeer
}

// Helper para crear un archivo válido para tests
func createTestArchivoMetadata() *model.ArchivoMetadata {
	id := uuid.New()
	subidoPor := uuid.New()
	now := time.Now().UTC()
	
	file, _ := model.NewArchivoMetadata(
		id,
		"test-file.txt",
		1024,
		"/path/to/file.txt",
		subidoPor,
		now,
	)
	return file
}

// Helper para crear un mensaje válido para tests
func createTestMensaje() *model.MensajeServidor {
	id := uuid.New()
	remitente := uuid.New()
	destino := uuid.New()
	now := time.Now().UTC()
	
	mensaje, _ := model.NewMensajeDirecto(
		id,
		remitente,
		destino,
		"Test message content",
		now,
		uuid.Nil, // Sin archivo adjunto
	)
	return mensaje
}

func TestReplicaNotifier_Subscribe(t *testing.T) {
	notifier := NewReplicaNotifier()
	observer := NewMockReplicaObserver()
	
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

func TestReplicaNotifier_Unsubscribe(t *testing.T) {
	notifier := NewReplicaNotifier()
	observer1 := NewMockReplicaObserver()
	observer2 := NewMockReplicaObserver()
	
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

func TestReplicaNotifier_NotifyUserReplicated(t *testing.T) {
	notifier := NewReplicaNotifier()
	observer := NewMockReplicaObserver()
	notifier.Subscribe(observer)
	
	user := createTestUser()
	peer := createTestPeer()
	notifier.NotifyUserReplicated(user, peer)
	
	// Verificar que se llamó al método correcto del observador
	if !observer.userReplicatedCalled {
		t.Error("OnUserReplicated no fue llamado")
	}
	
	// Verificar que se pasaron los objetos correctos
	if observer.lastUser != user {
		t.Error("El usuario pasado al observador no es el esperado")
	}
	
	if observer.lastPeer != peer {
		t.Error("El peer pasado al observador no es el esperado")
	}
}

func TestReplicaNotifier_NotifyChannelReplicated(t *testing.T) {
	notifier := NewReplicaNotifier()
	observer := NewMockReplicaObserver()
	notifier.Subscribe(observer)
	
	channel := createTestCanal()
	peer := createTestPeer()
	notifier.NotifyChannelReplicated(channel, peer)
	
	// Verificar que se llamó al método correcto del observador
	if !observer.channelReplicatedCalled {
		t.Error("OnChannelReplicated no fue llamado")
	}
	
	// Verificar que se pasaron los objetos correctos
	if observer.lastChannel != channel {
		t.Error("El canal pasado al observador no es el esperado")
	}
	
	if observer.lastPeer != peer {
		t.Error("El peer pasado al observador no es el esperado")
	}
}

func TestReplicaNotifier_NotifyMessageReplicated(t *testing.T) {
	notifier := NewReplicaNotifier()
	observer := NewMockReplicaObserver()
	notifier.Subscribe(observer)
	
	message := createTestMensaje()
	peer := createTestPeer()
	notifier.NotifyMessageReplicated(message, peer)
	
	// Verificar que se llamó al método correcto del observador
	if !observer.messageReplicatedCalled {
		t.Error("OnMessageReplicated no fue llamado")
	}
	
	// Verificar que se pasaron los objetos correctos
	if observer.lastMessage != message {
		t.Error("El mensaje pasado al observador no es el esperado")
	}
	
	if observer.lastPeer != peer {
		t.Error("El peer pasado al observador no es el esperado")
	}
}

func TestReplicaNotifier_NotifyFileReplicated(t *testing.T) {
	notifier := NewReplicaNotifier()
	observer := NewMockReplicaObserver()
	notifier.Subscribe(observer)
	
	file := createTestArchivoMetadata()
	peer := createTestPeer()
	notifier.NotifyFileReplicated(file, peer)
	
	// Verificar que se llamó al método correcto del observador
	if !observer.fileReplicatedCalled {
		t.Error("OnFileReplicated no fue llamado")
	}
	
	// Verificar que se pasaron los objetos correctos
	if observer.lastFile != file {
		t.Error("El archivo pasado al observador no es el esperado")
	}
	
	if observer.lastPeer != peer {
		t.Error("El peer pasado al observador no es el esperado")
	}
}
