package model

import (
	"errors"

	"github.com/google/uuid"
)

// Errores de validación para ChatPrivadoUsuario
var (
	ErrChatPrivadoUsuarioChatIDNil   = errors.New("id de chat privado inválido")
	ErrChatPrivadoUsuarioUsuarioIDNil = errors.New("id de usuario inválido")
	ErrChatPrivadoUsuariosIguales     = errors.New("no se puede crear un chat privado con el mismo usuario")
)

// ChatPrivadoUsuario representa la relación entre un chat privado y un usuario (participante)
type ChatPrivadoUsuario struct {
	chatPrivadoID uuid.UUID
	usuarioID     uuid.UUID
}

// NewChatPrivadoUsuario crea una nueva relación ChatPrivadoUsuario validando sus invariantes
func NewChatPrivadoUsuario(
	chatPrivadoID, usuarioID uuid.UUID,
) (*ChatPrivadoUsuario, error) {
	if chatPrivadoID == uuid.Nil {
		return nil, ErrChatPrivadoUsuarioChatIDNil
	}
	if usuarioID == uuid.Nil {
		return nil, ErrChatPrivadoUsuarioUsuarioIDNil
	}
	return &ChatPrivadoUsuario{
		chatPrivadoID: chatPrivadoID,
		usuarioID:     usuarioID,
	}, nil
}

// Getters
func (c *ChatPrivadoUsuario) ChatPrivadoID() uuid.UUID { return c.chatPrivadoID }
func (c *ChatPrivadoUsuario) UsuarioID() uuid.UUID    { return c.usuarioID }
