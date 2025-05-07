package observer

import (
	"sync"

	"model"
)

// ReplicaNotifier implementa IReplicaObserver y almacena una lista de observadores.
// El ReplicaManager invocará sus métodos NotifyXXX tras cada operación de réplica.
type ReplicaNotifier struct {
	observers []IReplicaObserver
	mu        sync.RWMutex
}

// NewReplicaNotifier crea una nueva instancia de ReplicaNotifier
func NewReplicaNotifier() *ReplicaNotifier {
	return &ReplicaNotifier{
		observers: make([]IReplicaObserver, 0),
	}
}

// Subscribe añade un observador a la lista de suscriptores
func (n *ReplicaNotifier) Subscribe(observer IReplicaObserver) {
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
func (n *ReplicaNotifier) Unsubscribe(observer IReplicaObserver) {
	n.mu.Lock()
	defer n.mu.Unlock()
	
	filtered := make([]IReplicaObserver, 0)
	for _, o := range n.observers {
		if o != observer {
			filtered = append(filtered, o)
		}
	}
	
	n.observers = filtered
}

// NotifyUserReplicated notifica a todos los observadores que un usuario ha sido replicado
func (n *ReplicaNotifier) NotifyUserReplicated(user *model.UsuarioServidor, toPeer *model.Peer) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnUserReplicated(user, toPeer)
	}
}

// NotifyChannelReplicated notifica a todos los observadores que un canal ha sido replicado
func (n *ReplicaNotifier) NotifyChannelReplicated(channel *model.CanalServidor, toPeer *model.Peer) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnChannelReplicated(channel, toPeer)
	}
}

// NotifyMessageReplicated notifica a todos los observadores que un mensaje ha sido replicado
func (n *ReplicaNotifier) NotifyMessageReplicated(msg *model.MensajeServidor, toPeer *model.Peer) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnMessageReplicated(msg, toPeer)
	}
}

// NotifyFileReplicated notifica a todos los observadores que un archivo ha sido replicado
func (n *ReplicaNotifier) NotifyFileReplicated(file *model.ArchivoMetadata, toPeer *model.Peer) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnFileReplicated(file, toPeer)
	}
}

// ObserversCount devuelve el número de observadores registrados
func (n *ReplicaNotifier) ObserversCount() int {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	return len(n.observers)
}
