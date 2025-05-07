package model

import (
	"errors"
)

// Errores de validación para EstadoInvitacion
var (
	ErrInvitacionEstadoInvalido = errors.New("estado de invitación inválido")
)

// EstadoInvitacion representa los posibles estados de una invitación a un canal
type EstadoInvitacion string

// Constantes para los estados de invitación
const (
	InvitacionPendiente EstadoInvitacion = "PENDIENTE"
	InvitacionAceptada  EstadoInvitacion = "ACEPTADA"
	InvitacionRechazada EstadoInvitacion = "RECHAZADA"
)

// Valid verifica si el valor es un estado de invitación válido
func (e EstadoInvitacion) Valid() bool {
	switch e {
	case InvitacionPendiente, InvitacionAceptada, InvitacionRechazada:
		return true
	default:
		return false
	}
}
