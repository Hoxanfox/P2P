package service

import (
	"github.com/google/uuid"
	"model"
)

// MembershipService define las operaciones para la gestión de membresías en canales
type MembershipService interface {
	// JoinChannel permite a un usuario unirse a un canal
	JoinChannel(userID, channelID uuid.UUID) error
	
	// LeaveChannel permite a un usuario salir de un canal
	LeaveChannel(userID, channelID uuid.UUID) error
	
	// ListMembers obtiene todos los miembros de un canal
	ListMembers(channelID uuid.UUID) ([]*model.UsuarioServidor, error)
}
