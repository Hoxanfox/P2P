package repository

import (
	"context"
	
	"github.com/google/uuid"
	"model"
)

// IMessageRepository define las operaciones para el repositorio de mensajes
type IMessageRepository interface {
    // CRUD
    Save(ctx context.Context, m *model.MensajeServidor) error
    Update(ctx context.Context, m *model.MensajeServidor) error
    Delete(ctx context.Context, id uuid.UUID) error
    FindByID(ctx context.Context, id uuid.UUID) (*model.MensajeServidor, error)

    // Consultas por canal/usuario
    FindByChannel(ctx context.Context, channelID uuid.UUID, limit, offset int) ([]*model.MensajeServidor, error)
    FindByUser(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*model.MensajeServidor, error)
    FindDirect(ctx context.Context, a, b uuid.UUID, limit, offset int) ([]*model.MensajeServidor, error)
}
