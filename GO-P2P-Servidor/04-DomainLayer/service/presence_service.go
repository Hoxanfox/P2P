package service

import (
	"github.com/google/uuid"
)

// PresenceService define las operaciones para gestionar la presencia de usuarios en tiempo real
type PresenceService interface {
	// MarkConnected marca a un usuario como conectado
	MarkConnected(userID uuid.UUID) error
	
	// MarkDisconnected marca a un usuario como desconectado
	MarkDisconnected(userID uuid.UUID) error
	
	// ListConnected lista todos los usuarios conectados actualmente
	ListConnected() ([]UsuarioDTO, error)
}
