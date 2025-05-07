package service

import (
	"time"
	
	"github.com/google/uuid"
)

// MessageService define las operaciones para envío y recuperación de mensajes
type MessageService interface {
	// SendDirect envía un mensaje directo entre dos usuarios
	SendDirect(
		remitenteID, destinoID uuid.UUID,
		contenido string,
		archivoID uuid.UUID,
	) (*MessageDTO, error)

	// SendChannel envía un mensaje a un canal
	SendChannel(
		remitenteID, channelID uuid.UUID,
		contenido string,
		archivoID uuid.UUID,
	) (*MessageDTO, error)

	// ListChannelMessages lista los mensajes de un canal en un rango de tiempo
	ListChannelMessages(
		channelID uuid.UUID,
		since, until time.Time,
	) ([]MessageDTO, error)

	// ListDirectMessages lista los mensajes directos entre dos usuarios en un rango de tiempo
	ListDirectMessages(
		userA, userB uuid.UUID,
		since, until time.Time,
	) ([]MessageDTO, error)
}
