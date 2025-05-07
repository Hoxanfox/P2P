package model

import (
    "errors"
    "time"

    "github.com/google/uuid"
)

// Errores de validación para HeartbeatLog
var (
    ErrHeartbeatLogIDNil       = errors.New("id de heartbeat inválido")
    ErrHeartbeatNodoIDNil      = errors.New("nodoId de heartbeat inválido")
    ErrHeartbeatEnviadoZero    = errors.New("enviadoAt no puede ser cero")
    ErrHeartbeatRecibidoZero   = errors.New("recibidoAt no puede ser cero")
    ErrHeartbeatRecibidoAntes  = errors.New("recibidoAt no puede ser anterior a enviadoAt")
)

// HeartbeatLog representa el registro de un pulso (heartbeat) entre nodos.
type HeartbeatLog struct {
    id         uuid.UUID
    nodoID     uuid.UUID
    enviadoAt  time.Time
    recibidoAt time.Time
}

// NewHeartbeatLog crea un HeartbeatLog validando sus invariantes:
// - id y nodoID no pueden ser Nil
// - enviadoAt y recibidoAt no pueden ser cero
// - recibidoAt no puede ser antes de enviadoAt
func NewHeartbeatLog(
    id, nodoID uuid.UUID,
    enviadoAt, recibidoAt time.Time,
) (*HeartbeatLog, error) {
    if id == uuid.Nil {
        return nil, ErrHeartbeatLogIDNil
    }
    if nodoID == uuid.Nil {
        return nil, ErrHeartbeatNodoIDNil
    }
    if enviadoAt.IsZero() {
        return nil, ErrHeartbeatEnviadoZero
    }
    if recibidoAt.IsZero() {
        return nil, ErrHeartbeatRecibidoZero
    }
    if recibidoAt.Before(enviadoAt) {
        return nil, ErrHeartbeatRecibidoAntes
    }
    return &HeartbeatLog{
        id:         id,
        nodoID:     nodoID,
        enviadoAt:  enviadoAt,
        recibidoAt: recibidoAt,
    }, nil
}

// Getters
func (h *HeartbeatLog) ID() uuid.UUID        { return h.id }
func (h *HeartbeatLog) NodoID() uuid.UUID    { return h.nodoID }
func (h *HeartbeatLog) EnviadoAt() time.Time { return h.enviadoAt }
func (h *HeartbeatLog) RecibidoAt() time.Time { return h.recibidoAt }
