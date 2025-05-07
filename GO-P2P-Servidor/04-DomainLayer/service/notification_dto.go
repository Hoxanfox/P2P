package service

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// NotificationDTO representa una notificación simplificada para transferencia de datos
type NotificationDTO struct {
	ID            uuid.UUID  `json:"id"`
	UsuarioID     uuid.UUID  `json:"usuarioID"`
	Contenido     string     `json:"contenido"`
	Fecha         time.Time  `json:"fecha"`
	Leido         bool       `json:"leido"`
	InvitacionID  uuid.UUID  `json:"invitacionID,omitempty"`
}

// MapNotificacionToDTO convierte un modelo Notificacion a un DTO
func MapNotificacionToDTO(n *model.Notificacion) *NotificationDTO {
	return &NotificationDTO{
		ID:           n.ID(),
		UsuarioID:    n.UsuarioID(),
		Contenido:    n.Contenido(),
		Fecha:        n.Fecha(),
		Leido:        n.Leido(),
		InvitacionID: n.InvitacionID(),
	}
}
