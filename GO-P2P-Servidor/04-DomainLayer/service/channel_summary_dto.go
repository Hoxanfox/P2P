package service

import (
	"time"

	"github.com/google/uuid"
)

// ChannelSummaryDTO representa un resumen de información de canal para administradores
type ChannelSummaryDTO struct {
	ID               uuid.UUID `json:"id"`
	Nombre           string    `json:"nombre"`
	NumMiembros      int       `json:"numMiembros"`
	NumMensajes      int       `json:"numMensajes"`
	UltimoMensaje    time.Time `json:"ultimoMensaje"`
	FechaCreacion    time.Time `json:"fechaCreacion"`
	TipoCanal        string    `json:"tipoCanal"`
	ArchivoCompartidos int      `json:"archivosCompartidos"`
}
