package service

import (
	"github.com/google/uuid"
	"model"
)

// InvitationService define las operaciones para enviar y procesar invitaciones a canales
type InvitationService interface {
	// SendInvitation envía una invitación a un usuario para unirse a un canal
	SendInvitation(
		channelID, destinatarioID uuid.UUID,
	) (*model.InvitacionCanal, error)

	// RespondInvitation procesa la respuesta a una invitación (aceptar o rechazar)
	RespondInvitation(
		invitationID uuid.UUID,
		accept bool,
	) (*model.InvitacionCanal, error)

	// ListPending lista todas las invitaciones pendientes para un usuario
	ListPending(userID uuid.UUID) ([]*model.InvitacionCanal, error)
}
