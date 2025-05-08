package service

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// HeartbeatService define las operaciones para envío y recepción de heartbeats entre nodos P2P
type HeartbeatService interface {
	// Start inicia el servicio de heartbeat
	Start() error

	// SendHeartbeat envía un heartbeat a un nodo específico
	SendHeartbeat(toPeerID uuid.UUID) error

	// ReceiveHeartbeat registra un heartbeat recibido de otro nodo
	ReceiveHeartbeat(
		fromPeerID uuid.UUID,
		timestamp time.Time,
	) error

	// ListLogs lista los logs de heartbeat para un nodo específico
	ListLogs(peerID uuid.UUID) ([]*model.HeartbeatLog, error)
}
