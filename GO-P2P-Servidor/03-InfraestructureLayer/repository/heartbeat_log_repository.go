package repository

import (
	"context"
	"database/sql"
	"time"

	"dao"
	"github.com/google/uuid"
	"model"
)

// HeartbeatLogRepository implementa IHeartbeatLogRepository
type HeartbeatLogRepository struct {
	dao *dao.HeartbeatLogMySQLDAO
}

// NewHeartbeatLogRepository crea una nueva instancia del repositorio
func NewHeartbeatLogRepository(db *sql.DB) *HeartbeatLogRepository {
	return &HeartbeatLogRepository{
		dao: dao.NewHeartbeatLogMySQLDAO(db),
	}
}

// Save guarda un registro de heartbeat
func (r *HeartbeatLogRepository) Save(ctx context.Context, h *model.HeartbeatLog) error {
	return r.dao.Save(h)
}

// FindByPeer encuentra logs de heartbeat por ID del nodo
func (r *HeartbeatLogRepository) FindByPeer(ctx context.Context, peerID uuid.UUID) ([]*model.HeartbeatLog, error) {
	return r.dao.FindByNodoID(peerID)
}

// PruneOlderThan elimina logs más antiguos que una fecha determinada
func (r *HeartbeatLogRepository) PruneOlderThan(ctx context.Context, cutoff time.Time) error {
	logs, err := r.dao.FindByTimeRange(time.Time{}, cutoff)
	if err != nil {
		return err
	}

	for _, log := range logs {
		if err := r.dao.Delete(log.ID()); err != nil {
			return err
		}
	}

	return nil
}
