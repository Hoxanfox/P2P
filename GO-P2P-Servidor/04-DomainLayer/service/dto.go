package service

import (
	"time"

	"github.com/google/uuid"
)

// UsuarioDTO representa un usuario simplificado para transferencia de datos
type UsuarioDTO struct {
	ID            uuid.UUID `json:"id"`
	NombreUsuario string    `json:"nombreUsuario"`
	Email         string    `json:"email"`
	FotoURL       string    `json:"fotoURL"`
	IPRegistrada  string    `json:"ipRegistrada"`
	FechaRegistro time.Time `json:"fechaRegistro"`
	IsConnected   bool      `json:"isConnected"`
}
