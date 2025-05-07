package service

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// MessageDTO representa un mensaje simplificado para transferencia de datos
type MessageDTO struct {
	ID               uuid.UUID  `json:"id"`
	RemitenteID      uuid.UUID  `json:"remitenteID"`
	DestinoUsuarioID uuid.UUID  `json:"destinoUsuarioID,omitempty"`
	CanalID          uuid.UUID  `json:"canalID,omitempty"`
	ChatPrivadoID    uuid.UUID  `json:"chatPrivadoID,omitempty"`
	Contenido        string     `json:"contenido"`
	Timestamp        time.Time  `json:"timestamp"`
	ArchivoID        uuid.UUID  `json:"archivoID,omitempty"`
}

// MapMensajeToDTO convierte un modelo MensajeServidor a un DTO
func MapMensajeToDTO(m *model.MensajeServidor) *MessageDTO {
	return &MessageDTO{
		ID:               m.ID(),
		RemitenteID:      m.RemitenteID(),
		DestinoUsuarioID: m.DestinoUsuarioID(),
		CanalID:          m.CanalID(),
		ChatPrivadoID:    m.ChatPrivadoID(),
		Contenido:        m.Contenido(),
		Timestamp:        m.Timestamp(),
		ArchivoID:        m.ArchivoID(),
	}
}
