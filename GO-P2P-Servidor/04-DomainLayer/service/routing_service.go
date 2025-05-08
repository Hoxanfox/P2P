package service

import (
	"github.com/google/uuid"
	"model"
)

// RoutingService define las operaciones para enrutamiento de mensajes entre nodos P2P
type RoutingService interface {
	// RouteMessage enruta un mensaje al nodo correspondiente si el destinatario 
	// no está en el mismo servidor
	RouteMessage(message *model.MensajeServidor) error
	
	// ListRoutes lista las rutas para un mensaje específico
	ListRoutes(messageID uuid.UUID) ([]*model.RoutedMessage, error)
}
