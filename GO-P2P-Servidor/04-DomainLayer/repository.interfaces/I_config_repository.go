package repository

import (
	"context"
	"model"
)

// IConfigRepository define las operaciones para el repositorio de configuración del servidor
type IConfigRepository interface {
    Get(ctx context.Context) (*model.ConfiguracionServidor, error)
    Update(ctx context.Context, cfg *model.ConfiguracionServidor) error
}
