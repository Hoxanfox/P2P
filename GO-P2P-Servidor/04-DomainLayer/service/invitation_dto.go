package service

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// InvitationDTO representa una invitación a un canal para transferencia de datos
type InvitationDTO struct {
	ID             uuid.UUID            `json:"id"`
	CanalID        uuid.UUID            `json:"canalID"`
	DestinatarioID uuid.UUID            `json:"destinatarioID"`
	Estado         model.EstadoInvitacion `json:"estado"`
	FechaEnvio     time.Time            `json:"fechaEnvio"`
}

// MapInvitacionToDTO convierte un modelo InvitacionCanal a un DTO
func MapInvitacionToDTO(inv *model.InvitacionCanal) *InvitationDTO {
	return &InvitationDTO{
		ID:             inv.ID(),
		CanalID:        inv.CanalID(),
		DestinatarioID: inv.DestinatarioID(),
		Estado:         inv.Estado(),
		FechaEnvio:     inv.FechaEnvio(),
	}
}
