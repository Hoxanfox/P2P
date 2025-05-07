package model

import (
    "errors"
    "time"

    "github.com/google/uuid"
)

// Errores de validación para RoutedMessage
var (
    ErrRoutedMessageMensajeIDNil   = errors.New("mensajeId inválido")
    ErrRoutedMessageDestinoIDNil   = errors.New("nodoDestinoId inválido")
    ErrRoutedMessageEnrutaAtZero   = errors.New("enrutaAt no puede ser cero")
)

// RoutedMessage representa el enrutamiento de un mensaje a un peer destino.
type RoutedMessage struct {
    mensajeID      uuid.UUID
    nodoDestinoID  uuid.UUID
    enrutaAt       time.Time
}

// NewRoutedMessage crea un RoutedMessage validando sus invariantes:
// - mensajeID no puede ser uuid.Nil
// - nodoDestinoID no puede ser uuid.Nil
// - enrutaAt no puede ser tiempo cero
func NewRoutedMessage(
    mensajeID, nodoDestinoID uuid.UUID,
    enrutaAt time.Time,
) (*RoutedMessage, error) {
    if mensajeID == uuid.Nil {
        return nil, ErrRoutedMessageMensajeIDNil
    }
    if nodoDestinoID == uuid.Nil {
        return nil, ErrRoutedMessageDestinoIDNil
    }
    if enrutaAt.IsZero() {
        return nil, ErrRoutedMessageEnrutaAtZero
    }
    return &RoutedMessage{
        mensajeID:     mensajeID,
        nodoDestinoID: nodoDestinoID,
        enrutaAt:      enrutaAt,
    }, nil
}

// Getters
func (r *RoutedMessage) MensajeID() uuid.UUID     { return r.mensajeID }
func (r *RoutedMessage) NodoDestinoID() uuid.UUID { return r.nodoDestinoID }
func (r *RoutedMessage) EnrutaAt() time.Time      { return r.enrutaAt }
