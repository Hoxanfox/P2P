package observer

import (
	"sync"

	"model"
)

// PeerNotifier implementa un publisher para los eventos relacionados con peers.
// Mantiene una lista de suscriptores ([]IPeerObserver) y los notifica cuando ocurren eventos.
// Los servicios de red (sobre todo HeartbeatService) lo usarán para emitir eventos.
type PeerNotifier struct {
	observers []IPeerObserver
	mu        sync.RWMutex
}

// NewPeerNotifier crea una nueva instancia de PeerNotifier
func NewPeerNotifier() *PeerNotifier {
	return &PeerNotifier{
		observers: make([]IPeerObserver, 0),
	}
}

// Subscribe añade un observador a la lista de suscriptores
func (n *PeerNotifier) Subscribe(observer IPeerObserver) {
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
func (n *PeerNotifier) Unsubscribe(observer IPeerObserver) {
	n.mu.Lock()
	defer n.mu.Unlock()
	
	filtered := make([]IPeerObserver, 0)
	for _, o := range n.observers {
		if o != observer {
			filtered = append(filtered, o)
		}
	}
	
	n.observers = filtered
}

// NotifyPeerConnected notifica a todos los observadores que un peer se ha conectado
func (n *PeerNotifier) NotifyPeerConnected(peer *model.Peer) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnPeerConnected(peer)
	}
}

// NotifyPeerDisconnected notifica a todos los observadores que un peer se ha desconectado
func (n *PeerNotifier) NotifyPeerDisconnected(peer *model.Peer) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnPeerDisconnected(peer)
	}
}

// NotifyPeerHeartbeatMissed notifica a todos los observadores que se han perdido
// heartbeats de un peer
func (n *PeerNotifier) NotifyPeerHeartbeatMissed(peer *model.Peer) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	for _, o := range n.observers {
		o.OnPeerHeartbeatMissed(peer)
	}
}

// ObserversCount devuelve el número de observadores registrados
func (n *PeerNotifier) ObserversCount() int {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	return len(n.observers)
}
