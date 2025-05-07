package repository

import (
	"context"
	
	"github.com/google/uuid"
	"model"
)

// IPeerRepository define las operaciones para el repositorio de peers
type IPeerRepository interface {
    Save(ctx context.Context, p *model.Peer) error
    Update(ctx context.Context, p *model.Peer) error
    Delete(ctx context.Context, id uuid.UUID) error
    FindByID(ctx context.Context, id uuid.UUID) (*model.Peer, error)
    ListAll(ctx context.Context) ([]*model.Peer, error)
    ListByState(ctx context.Context, state model.NodoEstado) ([]*model.Peer, error)
}
