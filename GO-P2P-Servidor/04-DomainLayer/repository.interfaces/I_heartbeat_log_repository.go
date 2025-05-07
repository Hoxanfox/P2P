package repository

import (
	"context"
	"time"
	
	"github.com/google/uuid"
	"model"
)

// IHeartbeatLogRepository define las operaciones para el repositorio de logs de heartbeat
type IHeartbeatLogRepository interface {
    Save(ctx context.Context, h *model.HeartbeatLog) error
    FindByPeer(ctx context.Context, peerID uuid.UUID) ([]*model.HeartbeatLog, error)
    PruneOlderThan(ctx context.Context, cutoff time.Time) error
}
