package factory

import (
    "errors"
    "time"

    "github.com/google/uuid"

    "model"
)

// Errores de validación para ReplicaEvent
var (
    ErrEntidadTipoVacio = errors.New("tipo de entidad vacío")
    ErrEntidadIDNil     = errors.New("id de entidad inválido")
    ErrOrigenNodoIDNil  = errors.New("id de nodo de origen inválido")
)

// IReplicaFactory define la abstracción para crear ReplicaEvent
type IReplicaFactory interface {
    // Create valida los parámetros y crea un ReplicaEvent
    Create(entidadTipo string, entidadID, origenNodoID uuid.UUID) (*model.ReplicaEvent, error)
}

// ReplicaFactory es la implementación concreta de IReplicaFactory
type ReplicaFactory struct{}

// NewReplicaFactory construye una nueva instancia de ReplicaFactory
func NewReplicaFactory() IReplicaFactory {
    return &ReplicaFactory{}
}

// Create valida los invariantes y delega en el constructor de dominio
func (f *ReplicaFactory) Create(entidadTipo string, entidadID, origenNodoID uuid.UUID) (*model.ReplicaEvent, error) {
    if entidadTipo == "" {
        return nil, ErrEntidadTipoVacio
    }
    if entidadID == uuid.Nil {
        return nil, ErrEntidadIDNil
    }
    if origenNodoID == uuid.Nil {
        return nil, ErrOrigenNodoIDNil
    }
    
    // Generar ID y timestamp para el evento
    id := uuid.New()
    eventoAt := time.Now().UTC()
    
    // Llamar al constructor del modelo con todos los parámetros
    return model.NewReplicaEvent(id, entidadTipo, entidadID, eventoAt, origenNodoID)
}
