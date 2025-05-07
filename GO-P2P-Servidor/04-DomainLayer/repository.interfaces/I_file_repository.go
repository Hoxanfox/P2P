package repository

import (
	"context"
	
	"github.com/google/uuid"
	"model"
)

// IFileRepository define las operaciones para el repositorio de archivos
type IFileRepository interface {
    Save(ctx context.Context, f *model.ArchivoMetadata) error
    Delete(ctx context.Context, id uuid.UUID) error
    FindByID(ctx context.Context, id uuid.UUID) (*model.ArchivoMetadata, error)
    FindByMessage(ctx context.Context, messageID uuid.UUID) ([]*model.ArchivoMetadata, error)
}
