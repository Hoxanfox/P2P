package factory

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// IInvitacionCanalFactory define la creación de InvitacionCanal.
type IInvitacionCanalFactory interface {
	// Create genera una InvitacionCanal con estado PENDIENTE.
	Create(canalID, destinatarioID uuid.UUID) (*model.InvitacionCanal, error)

	// CreateWithStatus genera una InvitacionCanal con el estado especificado.
	CreateWithStatus(canalID, destinatarioID uuid.UUID, estado model.EstadoInvitacion) (*model.InvitacionCanal, error)
}

// invitacionCanalFactory es la implementación de IInvitacionCanalFactory.
type invitacionCanalFactory struct{}

// NewInvitacionCanalFactory devuelve una instancia de IInvitacionCanalFactory.
func NewInvitacionCanalFactory() IInvitacionCanalFactory {
	return &invitacionCanalFactory{}
}

// Create implementa IInvitacionCanalFactory: crea una InvitacionCanal con estado PENDIENTE.
func (f *invitacionCanalFactory) Create(canalID, destinatarioID uuid.UUID) (*model.InvitacionCanal, error) {
	return f.CreateWithStatus(canalID, destinatarioID, model.InvitacionPendiente)
}

// CreateWithStatus implementa IInvitacionCanalFactory: crea una InvitacionCanal con el estado especificado.
func (f *invitacionCanalFactory) CreateWithStatus(
	canalID, destinatarioID uuid.UUID, 
	estado model.EstadoInvitacion,
) (*model.InvitacionCanal, error) {
	id := uuid.New()
	ahora := time.Now().UTC()
	return model.NewInvitacionCanal(id, canalID, destinatarioID, estado, ahora)
}
