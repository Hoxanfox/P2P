package service

import (
	"github.com/google/uuid"
)

// ConnectionService define las operaciones para control de conexiones activas
type ConnectionService interface {
	// Disconnect desconecta forzadamente a un usuario
	Disconnect(userID uuid.UUID) error
	
	// GetConnectionLimit obtiene el límite actual de conexiones simultáneas
	GetConnectionLimit() (int, error)
	
	// SetConnectionLimit establece un nuevo límite de conexiones simultáneas
	SetConnectionLimit(max int) error
}
