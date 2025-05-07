package factory

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// INotificacionFactory define la creación de Notificacion.
type INotificacionFactory interface {
	// Create genera una Notificacion básica para un usuario.
	Create(usuarioID uuid.UUID, contenido string) (*model.Notificacion, error)
	
	// CreateForInvitacion genera una Notificacion asociada a una invitación.
	CreateForInvitacion(usuarioID, invitacionID uuid.UUID, contenido string) (*model.Notificacion, error)
}

// notificacionFactory es la implementación de INotificacionFactory.
type notificacionFactory struct{}

// NewNotificacionFactory devuelve una instancia de INotificacionFactory.
func NewNotificacionFactory() INotificacionFactory {
	return &notificacionFactory{}
}

// Create implementa INotificacionFactory: crea una Notificacion básica para un usuario.
func (f *notificacionFactory) Create(usuarioID uuid.UUID, contenido string) (*model.Notificacion, error) {
	id := uuid.New()
	ahora := time.Now().UTC()
	return model.NewNotificacion(id, usuarioID, contenido, ahora, uuid.Nil)
}

// CreateForInvitacion implementa INotificacionFactory: crea una Notificacion asociada a una invitación.
func (f *notificacionFactory) CreateForInvitacion(
	usuarioID, invitacionID uuid.UUID, 
	contenido string,
) (*model.Notificacion, error) {
	id := uuid.New()
	ahora := time.Now().UTC()
	return model.NewNotificacion(id, usuarioID, contenido, ahora, invitacionID)
}
