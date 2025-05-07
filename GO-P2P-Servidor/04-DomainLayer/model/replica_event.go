package model

import (
    "errors"
    "time"

    "github.com/google/uuid"
)

// Errores de validación para ReplicaEvent
var (
    ErrReplicaEventIDNil       = errors.New("id de evento de réplica inválido")
    ErrEntidadTipoVacio        = errors.New("entidadTipo no puede estar vacío")
    ErrEntidadIDNil            = errors.New("entidadId inválido")
    ErrReplicaEventAtZero      = errors.New("eventoAt no puede ser cero")
    ErrOrigenNodoIDNil         = errors.New("origenNodoId inválido")
)

// ReplicaEvent representa un evento de replicación en la red P2P.
type ReplicaEvent struct {
    id           uuid.UUID
    entidadTipo  string
    entidadID    uuid.UUID
    eventoAt     time.Time
    origenNodoID uuid.UUID
}

// NewReplicaEvent crea un ReplicaEvent validando sus invariantes.
func NewReplicaEvent(
    id uuid.UUID,
    entidadTipo string,
    entidadID uuid.UUID,
    eventoAt time.Time,
    origenNodoID uuid.UUID,
) (*ReplicaEvent, error) {
    if id == uuid.Nil {
        return nil, ErrReplicaEventIDNil
    }
    if entidadTipo == "" {
        return nil, ErrEntidadTipoVacio
    }
    if entidadID == uuid.Nil {
        return nil, ErrEntidadIDNil
    }
    if eventoAt.IsZero() {
        return nil, ErrReplicaEventAtZero
    }
    if origenNodoID == uuid.Nil {
        return nil, ErrOrigenNodoIDNil
    }
    return &ReplicaEvent{
        id:           id,
        entidadTipo:  entidadTipo,
        entidadID:    entidadID,
        eventoAt:     eventoAt,
        origenNodoID: origenNodoID,
    }, nil
}

// Getters
func (r *ReplicaEvent) ID() uuid.UUID         { return r.id }
func (r *ReplicaEvent) EntidadTipo() string   { return r.entidadTipo }
func (r *ReplicaEvent) EntidadID() uuid.UUID  { return r.entidadID }
func (r *ReplicaEvent) EventoAt() time.Time   { return r.eventoAt }
func (r *ReplicaEvent) OrigenNodoID() uuid.UUID { return r.origenNodoID }
