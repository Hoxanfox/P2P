package service

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// HeartbeatLogDTO representa un registro de heartbeat simplificado para transferencia de datos
type HeartbeatLogDTO struct {
	ID           uuid.UUID `json:"id"`
	PeerID       uuid.UUID `json:"peerID"`
	Timestamp    time.Time `json:"timestamp"`
	Latencia     int64     `json:"latencia"`
	Estado       string    `json:"estado"`
	IP           string    `json:"ip"`
	PuertoP2P    int       `json:"puertoP2P"`
}

// MapHeartbeatLogToDTO convierte un modelo HeartbeatLog a un DTO
// Nota: Esta función asume que los siguientes métodos existen en el modelo HeartbeatLog:
// - ID() uuid.UUID
// - PeerID() uuid.UUID
// - Timestamp() time.Time
// - Latencia() int64
// - Estado() string
// - IP() string
// - PuertoP2P() int
func MapHeartbeatLogToDTO(h *model.HeartbeatLog) *HeartbeatLogDTO {
	// NOTA: Esta función requiere implementación de los métodos en el modelo HeartbeatLog
	return &HeartbeatLogDTO{
		ID:          h.ID(),
		PeerID:      h.PeerID(),
		Timestamp:   h.Timestamp(),
		Latencia:    h.Latencia(), 
		Estado:      h.Estado(),
		IP:          h.IP(),
		PuertoP2P:   h.PuertoP2P(),
	}
}
