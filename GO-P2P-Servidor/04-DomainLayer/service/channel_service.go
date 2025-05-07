package service

import (
	"github.com/google/uuid"
	"model"
)

// ChannelService define las operaciones para la gestión de canales
type ChannelService interface {
	// Create crea un nuevo canal con miembros iniciales
	Create(
		nombre, descripcion string,
		tipo model.CanalTipo,
		miembrosIniciales []uuid.UUID,
	) (*model.CanalServidor, error)

	// ListAll lista todos los canales a los que pertenece un usuario
	ListAll(userID uuid.UUID) ([]*model.CanalServidor, error)
	
	// GetByID obtiene un canal por su ID
	GetByID(channelID uuid.UUID) (*model.CanalServidor, error)
}
