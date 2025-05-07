package service

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// ChannelDTO representa un canal simplificado para transferencia de datos
type ChannelDTO struct {
	ID          uuid.UUID      `json:"id"`
	Nombre      string         `json:"nombre"`
	Descripcion string         `json:"descripcion"`
	Tipo        model.CanalTipo `json:"tipo"`
	CreadoEn    time.Time      `json:"creadoEn"`
	Miembros    []uuid.UUID    `json:"miembros"`
}

// MapCanalToDTO convierte un modelo CanalServidor a un DTO
func MapCanalToDTO(c *model.CanalServidor, miembros []uuid.UUID) *ChannelDTO {
	return &ChannelDTO{
		ID:          c.ID(),
		Nombre:      c.Nombre(),
		Descripcion: c.Descripcion(),
		Tipo:        c.Tipo(),
		CreadoEn:    c.CreadoEn(),
		Miembros:    miembros,
	}
}
