package service

import (
	"github.com/google/uuid"
	"model"
)

// AuthService define las operaciones para autenticación y registro de usuarios
type AuthService interface {
	// Register crea un nuevo usuario en el sistema
	Register(
		nombre, email, password, foto, ip string,
	) (*model.UsuarioServidor, error)

	// Login autentica un usuario en el sistema
	Login(
		email, password, ip string,
	) (*model.UsuarioServidor, error)

	// Logout cierra la sesión de un usuario
	Logout(
		userID uuid.UUID,
	) error
}
