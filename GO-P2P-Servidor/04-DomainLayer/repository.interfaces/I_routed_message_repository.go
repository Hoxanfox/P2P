package repository

import (
	"context"
	
	"github.com/google/uuid"
	"model"
)

// IRoutedMessageRepository define las operaciones para el repositorio de mensajes ruteados
type IRoutedMessageRepository interface {
    Save(ctx context.Context, r *model.RoutedMessage) error
    FindByMessage(ctx context.Context, messageID uuid.UUID) ([]*model.RoutedMessage, error)
    DeleteByMessage(ctx context.Context, messageID uuid.UUID) error
}
