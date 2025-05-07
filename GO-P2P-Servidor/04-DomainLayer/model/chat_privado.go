package model

import (
	"errors"

	"github.com/google/uuid"
)

// Errores de validación para ChatPrivado
var (
	ErrChatPrivadoIDNil = errors.New("id de chat privado inválido")
)

// ChatPrivado representa un chat privado entre dos usuarios
type ChatPrivado struct {
	id uuid.UUID
}

// NewChatPrivado crea un nuevo ChatPrivado validando sus invariantes
func NewChatPrivado(
	id uuid.UUID,
) (*ChatPrivado, error) {
	if id == uuid.Nil {
		return nil, ErrChatPrivadoIDNil
	}
	return &ChatPrivado{
		id: id,
	}, nil
}

// Getters
func (c *ChatPrivado) ID() uuid.UUID { return c.id }

// Nota: Las relaciones entre ChatPrivado y UsuarioServidor o MensajeServidor
// se manejan a través de tablas relacionales separadas o repositorios específicos.
// Este modelo solo representa la entidad principal.
