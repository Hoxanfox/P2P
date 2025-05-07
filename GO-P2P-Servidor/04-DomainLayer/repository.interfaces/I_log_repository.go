package repository

import (
	"context"
	"time"
	
	"github.com/google/uuid"
	"model"
)

// ILogRepository define las operaciones para el repositorio de logs
type ILogRepository interface {
    Save(ctx context.Context, entry *model.LogEntry) error
    Delete(ctx context.Context, id uuid.UUID) error
    FindByID(ctx context.Context, id uuid.UUID) (*model.LogEntry, error)

    // Consultas de auditoría
    ListAll(ctx context.Context) ([]*model.LogEntry, error)
    ListByUser(ctx context.Context, userID uuid.UUID) ([]*model.LogEntry, error)
    ListByType(ctx context.Context, t model.EventoTipo) ([]*model.LogEntry, error)
    ListByDateRange(ctx context.Context, from, to time.Time) ([]*model.LogEntry, error)
}
