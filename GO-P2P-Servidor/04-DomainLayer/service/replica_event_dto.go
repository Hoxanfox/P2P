package service

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// TipoEntidad define el tipo de entidad que se replica
type TipoEntidad string

const (
	EntidadUsuario TipoEntidad = "USUARIO"
	EntidadMensaje TipoEntidad = "MENSAJE"
	EntidadCanal   TipoEntidad = "CANAL"
	EntidadArchivo TipoEntidad = "ARCHIVO"
)

// TipoOperacion define el tipo de operación de replicación
type TipoOperacion string

const (
	OperacionCrear    TipoOperacion = "CREAR"
	OperacionActualizar TipoOperacion = "ACTUALIZAR"
	OperacionEliminar TipoOperacion = "ELIMINAR"
)

// ReplicaEventDTO representa un evento de replicación para transferencia de datos
type ReplicaEventDTO struct {
	ID             uuid.UUID     `json:"id"`
	TipoEntidad    TipoEntidad   `json:"tipoEntidad"`
	EntidadID      uuid.UUID     `json:"entidadID"`
	TipoOperacion  TipoOperacion `json:"tipoOperacion"`
	Timestamp      time.Time     `json:"timestamp"`
	PeerDestino    uuid.UUID     `json:"peerDestino,omitempty"`
	Estado         string        `json:"estado"`
	DatosSerializados string      `json:"datosSerializados,omitempty"`
}

// MapReplicaEventToDTO convierte un modelo ReplicaEvent a un DTO
func MapReplicaEventToDTO(r *model.ReplicaEvent) *ReplicaEventDTO {
	return &ReplicaEventDTO{
		ID:              r.ID(),
		TipoEntidad:     TipoEntidad(r.TipoEntidad()),
		EntidadID:       r.EntidadID(),
		TipoOperacion:   TipoOperacion(r.TipoOperacion()),
		Timestamp:       r.Timestamp(),
		PeerDestino:     r.PeerDestino(),
		Estado:          string(r.Estado()),
		DatosSerializados: r.DatosSerializados(),
	}
}
