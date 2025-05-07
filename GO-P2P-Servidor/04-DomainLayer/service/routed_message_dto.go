package service

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// RoutedMessageDTO representa un mensaje ruteado para transferencia de datos
type RoutedMessageDTO struct {
	ID              uuid.UUID `json:"id"`
	MensajeID       uuid.UUID `json:"mensajeID"`
	PeerOrigen      uuid.UUID `json:"peerOrigen"`
	PeerDestino     uuid.UUID `json:"peerDestino"`
	FechaEnvio      time.Time `json:"fechaEnvio"`
	FechaRecepcion  time.Time `json:"fechaRecepcion,omitempty"`
	Estado          string    `json:"estado"`
	Intentos        int       `json:"intentos"`
	UltimoError     string    `json:"ultimoError,omitempty"`
}

// MapRoutedMessageToDTO convierte un modelo RoutedMessage a un DTO
func MapRoutedMessageToDTO(r *model.RoutedMessage) *RoutedMessageDTO {
	return &RoutedMessageDTO{
		ID:             r.ID(),
		MensajeID:      r.MensajeID(),
		PeerOrigen:     r.PeerOrigen(),
		PeerDestino:    r.PeerDestino(),
		FechaEnvio:     r.FechaEnvio(),
		FechaRecepcion: r.FechaRecepcion(),
		Estado:         string(r.Estado()),
		Intentos:       r.Intentos(),
		UltimoError:    r.UltimoError(),
	}
}
