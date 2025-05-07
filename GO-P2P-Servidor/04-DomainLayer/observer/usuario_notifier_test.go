package observer

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"model"
)

// MockUserObserver implementa la interfaz IUserObserver para tests
type MockUserObserver struct {
	registeredCalled     bool
	loggedInCalled       bool
	loggedOutCalled      bool
	updatedCalled        bool
	invitationSentCalled bool
	invitationRespCalled bool
	
	lastUser        *model.UsuarioServidor
	lastCanal       *model.CanalServidor
	lastChanges     []string
	lastByUser      *model.UsuarioServidor
	lastInvitedUser *model.UsuarioServidor
	lastAccepted    bool
}

func NewMockUserObserver() *MockUserObserver {
	return &MockUserObserver{}
}

func (m *MockUserObserver) OnUserRegistered(user *model.UsuarioServidor) {
	m.registeredCalled = true
	m.lastUser = user
}

func (m *MockUserObserver) OnUserLoggedIn(user *model.UsuarioServidor) {
	m.loggedInCalled = true
	m.lastUser = user
}

func (m *MockUserObserver) OnUserLoggedOut(user *model.UsuarioServidor) {
	m.loggedOutCalled = true
	m.lastUser = user
}

func (m *MockUserObserver) OnUserUpdated(user *model.UsuarioServidor, changedFields []string) {
	m.updatedCalled = true
	m.lastUser = user
	m.lastChanges = changedFields
}

func (m *MockUserObserver) OnInvitationSent(canal *model.CanalServidor, invitedUser *model.UsuarioServidor, byUser *model.UsuarioServidor) {
	m.invitationSentCalled = true
	m.lastCanal = canal
	m.lastInvitedUser = invitedUser
	m.lastByUser = byUser
}

func (m *MockUserObserver) OnInvitationResponded(canal *model.CanalServidor, user *model.UsuarioServidor, accepted bool) {
	m.invitationRespCalled = true
	m.lastCanal = canal
	m.lastUser = user
	m.lastAccepted = accepted
}

// Helper para crear un usuario válido para tests
func createTestUser() *model.UsuarioServidor {
	id := uuid.New()
	now := time.Now().UTC()
	user, _ := model.NewUsuarioServidor(
		id,
		"testuser",
		"test@example.com",
		"hashedpw",
		"https://example.com/profile.jpg",
		"192.168.1.100",
		now,
	)
	return user
}

// Helper para crear un canal válido para tests
func createTestCanal() *model.CanalServidor {
	id := uuid.New()
	canal, _ := model.NewCanalServidor(
		id,
		"test-channel",
		"Canal de prueba",
		model.CanalPublico,
	)
	return canal
}

func TestUserNotifier_Subscribe(t *testing.T) {
	notifier := NewUserNotifier()
	observer := NewMockUserObserver()
	
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

func TestUserNotifier_Unsubscribe(t *testing.T) {
	notifier := NewUserNotifier()
	observer1 := NewMockUserObserver()
	observer2 := NewMockUserObserver()
	
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

func TestUserNotifier_NotifyUserRegistered(t *testing.T) {
	notifier := NewUserNotifier()
	observer := NewMockUserObserver()
	notifier.Subscribe(observer)
	
	user := createTestUser()
	notifier.NotifyUserRegistered(user)
	
	// Verificar que se llamó al método correcto del observador
	if !observer.registeredCalled {
		t.Error("OnUserRegistered no fue llamado")
	}
	
	// Verificar que se pasó el usuario correcto
	if observer.lastUser != user {
		t.Error("El usuario pasado al observador no es el esperado")
	}
}

func TestUserNotifier_NotifyUserLoggedIn(t *testing.T) {
	notifier := NewUserNotifier()
	observer := NewMockUserObserver()
	notifier.Subscribe(observer)
	
	user := createTestUser()
	notifier.NotifyUserLoggedIn(user)
	
	if !observer.loggedInCalled {
		t.Error("OnUserLoggedIn no fue llamado")
	}
	
	if observer.lastUser != user {
		t.Error("El usuario pasado al observador no es el esperado")
	}
}

func TestUserNotifier_NotifyUserLoggedOut(t *testing.T) {
	notifier := NewUserNotifier()
	observer := NewMockUserObserver()
	notifier.Subscribe(observer)
	
	user := createTestUser()
	notifier.NotifyUserLoggedOut(user)
	
	if !observer.loggedOutCalled {
		t.Error("OnUserLoggedOut no fue llamado")
	}
	
	if observer.lastUser != user {
		t.Error("El usuario pasado al observador no es el esperado")
	}
}

func TestUserNotifier_NotifyUserUpdated(t *testing.T) {
	notifier := NewUserNotifier()
	observer := NewMockUserObserver()
	notifier.Subscribe(observer)
	
	user := createTestUser()
	changedFields := []string{"nombreUsuario", "fotoURL"}
	notifier.NotifyUserUpdated(user, changedFields)
	
	if !observer.updatedCalled {
		t.Error("OnUserUpdated no fue llamado")
	}
	
	if observer.lastUser != user {
		t.Error("El usuario pasado al observador no es el esperado")
	}
	
	if len(observer.lastChanges) != len(changedFields) {
		t.Error("Los campos cambiados no coinciden")
	}
}

func TestUserNotifier_NotifyInvitationSent(t *testing.T) {
	notifier := NewUserNotifier()
	observer := NewMockUserObserver()
	notifier.Subscribe(observer)
	
	canal := createTestCanal()
	invitedUser := createTestUser()
	byUser := createTestUser()
	
	notifier.NotifyInvitationSent(canal, invitedUser, byUser)
	
	if !observer.invitationSentCalled {
		t.Error("OnInvitationSent no fue llamado")
	}
	
	if observer.lastCanal != canal {
		t.Error("El canal pasado al observador no es el esperado")
	}
	
	if observer.lastInvitedUser != invitedUser {
		t.Error("El usuario invitado pasado al observador no es el esperado")
	}
	
	if observer.lastByUser != byUser {
		t.Error("El usuario que invita pasado al observador no es el esperado")
	}
}

func TestUserNotifier_NotifyInvitationResponded(t *testing.T) {
	notifier := NewUserNotifier()
	observer := NewMockUserObserver()
	notifier.Subscribe(observer)
	
	canal := createTestCanal()
	user := createTestUser()
	accepted := true
	
	notifier.NotifyInvitationResponded(canal, user, accepted)
	
	if !observer.invitationRespCalled {
		t.Error("OnInvitationResponded no fue llamado")
	}
	
	if observer.lastCanal != canal {
		t.Error("El canal pasado al observador no es el esperado")
	}
	
	if observer.lastUser != user {
		t.Error("El usuario pasado al observador no es el esperado")
	}
	
	if observer.lastAccepted != accepted {
		t.Error("El valor de aceptación pasado al observador no es el esperado")
	}
}
