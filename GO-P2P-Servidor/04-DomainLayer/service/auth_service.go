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
	) (*UsuarioDTO, error)

	// Login autentica un usuario en el sistema
	Login(
		email, password, ip string,
	) (*UsuarioDTO, error)

	// Logout cierra la sesión de un usuario
	Logout(
		userID uuid.UUID,
	) error
}

// MapUsuarioToDTO convierte un modelo UsuarioServidor a un DTO
func MapUsuarioToDTO(u *model.UsuarioServidor) *UsuarioDTO {
	return &UsuarioDTO{
		ID:            u.ID(),
		NombreUsuario: u.NombreUsuario(),
		Email:         u.Email(),
		FotoURL:       u.FotoURL(),
		IPRegistrada:  u.IPRegistrada(),
		FechaRegistro: u.FechaRegistro(),
		IsConnected:   u.IsConnected(),
	}
}
