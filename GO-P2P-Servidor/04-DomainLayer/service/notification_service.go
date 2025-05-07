package service

import (
	"github.com/google/uuid"
	"model"
)

// NotificationService define las operaciones para creación y gestión de notificaciones
type NotificationService interface {
	// Notify envía una notificación a un usuario
	Notify(
		userID uuid.UUID,
		contenido string,
	) (*model.Notificacion, error)

	// List obtiene todas las notificaciones de un usuario
	List(userID uuid.UUID) ([]*model.Notificacion, error)
	
	// MarkRead marca una notificación como leída
	MarkRead(notificationID uuid.UUID) error
}
