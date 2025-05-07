package service

import (
	"github.com/google/uuid"
)

// UserService define las operaciones para gestión de perfiles y listado de usuarios
type UserService interface {
	// GetAll obtiene todos los usuarios del sistema
	GetAll() ([]UsuarioDTO, error)
	
	// GetByID obtiene un usuario por su ID
	GetByID(userID uuid.UUID) (*UsuarioDTO, error)
	
	// UpdateProfile actualiza el perfil de un usuario
	UpdateProfile(
		userID uuid.UUID,
		nombre, email, foto string,
	) (*UsuarioDTO, error)
}
