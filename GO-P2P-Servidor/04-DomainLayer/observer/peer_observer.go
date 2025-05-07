package observer

import (
	"model"
)

// IPeerObserver define los callbacks para eventos de los nodos P2P:
// conexión inicial, desconexión o fallo de heartbeats.
type IPeerObserver interface {
	// OnPeerConnected se invoca cuando se recibe el primer heartbeat de un peer
	OnPeerConnected(peer *model.Peer)

	// OnPeerDisconnected se invoca cuando se detecta que un peer ha caído
	OnPeerDisconnected(peer *model.Peer)

	// OnPeerHeartbeatMissed se invoca tras X heartbeats fallidos
	OnPeerHeartbeatMissed(peer *model.Peer)
}
