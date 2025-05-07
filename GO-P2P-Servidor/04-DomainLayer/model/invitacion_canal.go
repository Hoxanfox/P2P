package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Errores de validación para InvitacionCanal
var (
	ErrInvitacionIDNil            = errors.New("id de invitación inválido")
	ErrInvitacionCanalIDNil       = errors.New("id de canal inválido")
	ErrInvitacionDestinatarioIDNil = errors.New("id de destinatario inválido")
	ErrInvitacionFechaEnvioZero   = errors.New("fecha de envío no puede ser cero")
)

// InvitacionCanal representa una invitación a un canal
type InvitacionCanal struct {
	id             uuid.UUID
	canalID        uuid.UUID
	destinatarioID uuid.UUID
	estado         EstadoInvitacion
	fechaEnvio     time.Time
}

// NewInvitacionCanal crea una nueva InvitacionCanal validando sus invariantes
func NewInvitacionCanal(
	id, canalID, destinatarioID uuid.UUID,
	estado EstadoInvitacion,
	fechaEnvio time.Time,
) (*InvitacionCanal, error) {
	if id == uuid.Nil {
		return nil, ErrInvitacionIDNil
	}
	if canalID == uuid.Nil {
		return nil, ErrInvitacionCanalIDNil
	}
	if destinatarioID == uuid.Nil {
		return nil, ErrInvitacionDestinatarioIDNil
	}
	if !estado.Valid() {
		return nil, ErrInvitacionEstadoInvalido
	}
	if fechaEnvio.IsZero() {
		return nil, ErrInvitacionFechaEnvioZero
	}
	return &InvitacionCanal{
		id:             id,
		canalID:        canalID,
		destinatarioID: destinatarioID,
		estado:         estado,
		fechaEnvio:     fechaEnvio,
	}, nil
}

// Getters
func (i *InvitacionCanal) ID() uuid.UUID              { return i.id }
func (i *InvitacionCanal) CanalID() uuid.UUID         { return i.canalID }
func (i *InvitacionCanal) DestinatarioID() uuid.UUID  { return i.destinatarioID }
func (i *InvitacionCanal) Estado() EstadoInvitacion   { return i.estado }
func (i *InvitacionCanal) FechaEnvio() time.Time      { return i.fechaEnvio }

// CambiarEstado actualiza el estado de la invitación
func (i *InvitacionCanal) CambiarEstado(nuevoEstado EstadoInvitacion) error {
	if !nuevoEstado.Valid() {
		return ErrInvitacionEstadoInvalido
	}
	i.estado = nuevoEstado
	return nil
}
