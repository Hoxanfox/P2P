package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Errores de validación para Notificacion
var (
	ErrNotificacionIDNil         = errors.New("id de notificación inválido")
	ErrNotificacionUsuarioIDNil  = errors.New("id de usuario receptor inválido")
	ErrNotificacionContenidoVacio = errors.New("contenido de notificación vacío")
	ErrNotificacionFechaZero     = errors.New("fecha de notificación no puede ser cero")
)

// Notificacion representa una notificación enviada a un usuario
type Notificacion struct {
	id           uuid.UUID
	usuarioID    uuid.UUID // receptor
	contenido    string
	fecha        time.Time
	leido        bool
	invitacionID uuid.UUID // Opcional: si está relacionada con una invitación
}

// NewNotificacion crea una nueva Notificacion validando sus invariantes
func NewNotificacion(
	id, usuarioID uuid.UUID,
	contenido string,
	fecha time.Time,
	invitacionID uuid.UUID,
) (*Notificacion, error) {
	if id == uuid.Nil {
		return nil, ErrNotificacionIDNil
	}
	if usuarioID == uuid.Nil {
		return nil, ErrNotificacionUsuarioIDNil
	}
	if contenido == "" {
		return nil, ErrNotificacionContenidoVacio
	}
	if fecha.IsZero() {
		return nil, ErrNotificacionFechaZero
	}
	return &Notificacion{
		id:           id,
		usuarioID:    usuarioID,
		contenido:    contenido,
		fecha:        fecha,
		leido:        false, // Por defecto, una notificación nueva no está leída
		invitacionID: invitacionID, // Puede ser uuid.Nil si no está relacionada con una invitación
	}, nil
}

// Getters
func (n *Notificacion) ID() uuid.UUID           { return n.id }
func (n *Notificacion) UsuarioID() uuid.UUID    { return n.usuarioID }
func (n *Notificacion) Contenido() string       { return n.contenido }
func (n *Notificacion) Fecha() time.Time        { return n.fecha }
func (n *Notificacion) Leido() bool             { return n.leido }
func (n *Notificacion) InvitacionID() uuid.UUID { return n.invitacionID }

// MarcarComoLeida marca la notificación como leída
func (n *Notificacion) MarcarComoLeida() {
	n.leido = true
}

// MarcarComoNoLeida marca la notificación como no leída
func (n *Notificacion) MarcarComoNoLeida() {
	n.leido = false
}
