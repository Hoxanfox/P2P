package repository

import (
	"context"
	
	"github.com/google/uuid"
	"model"
)

// IReplicaEventRepository define las operaciones para el repositorio de eventos de réplica
type IReplicaEventRepository interface {
    Save(ctx context.Context, e *model.ReplicaEvent) error
    Delete(ctx context.Context, id uuid.UUID) error
    FindByID(ctx context.Context, id uuid.UUID) (*model.ReplicaEvent, error)

    // Gestión de eventos pendientes
    ListPending(ctx context.Context) ([]*model.ReplicaEvent, error)
    MarkProcessed(ctx context.Context, id uuid.UUID) error
}
