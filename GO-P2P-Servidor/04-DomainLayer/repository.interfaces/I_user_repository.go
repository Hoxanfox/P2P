package repository

import (
	"context"
	
	"github.com/google/uuid"
	"model"
)

// IUserRepository define las operaciones para el repositorio de usuarios
type IUserRepository interface {
	// CRUD
	Save(ctx context.Context, u *model.UsuarioServidor) error
	Update(ctx context.Context, u *model.UsuarioServidor) error
	Delete(ctx context.Context, id uuid.UUID) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.UsuarioServidor, error)
	FindAll(ctx context.Context) ([]*model.UsuarioServidor, error)

	// Consultas específicas
	FindConnected(ctx context.Context) ([]*model.UsuarioServidor, error)
	FindByEmail(ctx context.Context, email string) (*model.UsuarioServidor, error)
}
