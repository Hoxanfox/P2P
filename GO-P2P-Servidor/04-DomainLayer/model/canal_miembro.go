package model

import (
	"errors"

	"github.com/google/uuid"
)

// Errores de validación para CanalMiembro
var (
	ErrCanalMiembroCanalIDNil    = errors.New("id de canal inválido")
	ErrCanalMiembroUsuarioIDNil  = errors.New("id de usuario inválido")
	ErrCanalMiembroRolVacio      = errors.New("rol de canal miembro vacío")
)

// CanalMiembro representa la relación entre un canal y un usuario
type CanalMiembro struct {
	canalID   uuid.UUID
	usuarioID uuid.UUID
	rol       string // "owner", "member", etc.
}

// NewCanalMiembro crea un nuevo CanalMiembro validando sus invariantes
func NewCanalMiembro(
	canalID, usuarioID uuid.UUID,
	rol string,
) (*CanalMiembro, error) {
	if canalID == uuid.Nil {
		return nil, ErrCanalMiembroCanalIDNil
	}
	if usuarioID == uuid.Nil {
		return nil, ErrCanalMiembroUsuarioIDNil
	}
	if rol == "" {
		return nil, ErrCanalMiembroRolVacio
	}
	return &CanalMiembro{
		canalID:   canalID,
		usuarioID: usuarioID,
		rol:       rol,
	}, nil
}

// Getters
func (c *CanalMiembro) CanalID() uuid.UUID   { return c.canalID }
func (c *CanalMiembro) UsuarioID() uuid.UUID { return c.usuarioID }
func (c *CanalMiembro) Rol() string          { return c.rol }
