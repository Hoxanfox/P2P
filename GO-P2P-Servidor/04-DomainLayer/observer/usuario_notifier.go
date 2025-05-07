package observer

import (
	"sync"

	"model"
)

// UserNotifier implementa un publisher para los eventos relacionados con usuarios.
// Mantiene una lista de suscriptores ([]IUserObserver) y los notifica cuando ocurren eventos.
type UserNotifier struct {
	observers []IUserObserver
	mu        sync.RWMutex
}

// NewUserNotifier crea una nueva instancia de UserNotifier
func NewUserNotifier() *UserNotifier {
	return &UserNotifier{
		observers: make([]IUserObserver, 0),
	}
}

// Subscribe añade un observador a la lista de suscriptores
func (n *UserNotifier) Subscribe(observer IUserObserver) {
	n.mu.Lock()
	defer n.mu.Unlock()
	
	// Verificar que el observador no esté ya en la lista
	for _, o := range n.observers {
		if o == observer {
			return
		}
	}
	
	n.observers = append(n.observers, observer)
}

// Unsubscribe elimina un observador de la lista de suscriptores
func (n *UserNotifier) Unsubscribe(observer IUserObserver) {
	n.mu.Lock()
	defer n.mu.Unlock()
	
	filtered := make([]IUserObserver, 0)
	for _, o := range n.observers {
		if o != observer {
			filtered = append(filtered, o)
		}
	}
	
	n.observers = filtered
}

// NotifyUserRegistered notifica a todos los observadores que un usuario ha sido registrado
func (n *UserNotifier) NotifyUserRegistered(user *model.UsuarioServidor) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnUserRegistered(user)
	}
}

// NotifyUserLoggedIn notifica a todos los observadores que un usuario ha iniciado sesión
func (n *UserNotifier) NotifyUserLoggedIn(user *model.UsuarioServidor) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnUserLoggedIn(user)
	}
}

// NotifyUserLoggedOut notifica a todos los observadores que un usuario ha cerrado sesión
func (n *UserNotifier) NotifyUserLoggedOut(user *model.UsuarioServidor) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnUserLoggedOut(user)
	}
}

// NotifyUserUpdated notifica a todos los observadores que un usuario ha actualizado su perfil
func (n *UserNotifier) NotifyUserUpdated(user *model.UsuarioServidor, changedFields []string) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnUserUpdated(user, changedFields)
	}
}

// NotifyInvitationSent notifica a todos los observadores que se ha enviado una invitación
func (n *UserNotifier) NotifyInvitationSent(
	canal *model.CanalServidor, 
	invitedUser *model.UsuarioServidor,
	byUser *model.UsuarioServidor,
) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnInvitationSent(canal, invitedUser, byUser)
	}
}

// NotifyInvitationResponded notifica a todos los observadores que un usuario ha respondido a una invitación
func (n *UserNotifier) NotifyInvitationResponded(
	canal *model.CanalServidor, 
	user *model.UsuarioServidor, 
	accepted bool,
) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnInvitationResponded(canal, user, accepted)
	}
}

// ObserversCount devuelve el número de observadores registrados
func (n *UserNotifier) ObserversCount() int {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	return len(n.observers)
}
