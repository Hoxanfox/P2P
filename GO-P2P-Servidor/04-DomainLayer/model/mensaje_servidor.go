package model

import (
    "errors"
    "time"

    "github.com/google/uuid"
)

// Errores de validación para MensajeServidor
var (
    ErrMensajeIDNil           = errors.New("id de mensaje inválido")
    ErrRemitenteIDNil         = errors.New("id de remitente inválido")
    ErrContenidoVacio         = errors.New("contenido de mensaje vacío")
    ErrTimestampZero          = errors.New("timestamp de mensaje no puede ser cero")
    ErrDestinoYCanalAmbos     = errors.New("debe elegir un destino de usuario O un canal, no ambos")
    ErrSinDestinoNiCanal      = errors.New("debe proporcionar un destino de usuario O un canal")
    ErrMensajeChatPrivadoIDNil = errors.New("id de chat privado inválido")
)

// MensajeServidor representa un mensaje de texto en el dominio del servidor.
type MensajeServidor struct {
    id               uuid.UUID
    remitenteID      uuid.UUID
    destinoUsuarioID uuid.UUID // Opcional: si es mensaje directo
    canalID          uuid.UUID // Opcional: si es mensaje de canal
    chatPrivadoID    uuid.UUID // Opcional: si es mensaje de chat privado (1-a-1)
    contenido        string
    timestamp        time.Time
    archivoID        uuid.UUID // Opcional: si lleva adjunto
}

// NewMensajeDirecto crea un MensajeServidor para mensajes 1:1.
// destinoUsuarioID debe ser distinto de uuid.Nil, canalID debe ser uuid.Nil.
func NewMensajeDirecto(
    id, remitenteID, destinoUsuarioID uuid.UUID,
    contenido string,
    timestamp time.Time,
    archivoID uuid.UUID,
) (*MensajeServidor, error) {
    // Validar explícitamente que no se está pasando un destinoUsuarioID y un canalID
    if destinoUsuarioID == uuid.Nil {
        return nil, ErrSinDestinoNiCanal
    }
    return newMensaje(id, remitenteID, destinoUsuarioID, uuid.Nil, uuid.Nil, contenido, timestamp, archivoID)
}

// NewMensajeCanal crea un MensajeServidor para mensajes de canal.
// canalID debe ser distinto de uuid.Nil, destinoUsuarioID debe ser uuid.Nil.
func NewMensajeCanal(
    id, remitenteID, canalID uuid.UUID,
    contenido string,
    timestamp time.Time,
    archivoID uuid.UUID,
) (*MensajeServidor, error) {
    // Validar explícitamente que se está pasando un canalID válido
    if canalID == uuid.Nil {
        return nil, ErrSinDestinoNiCanal
    }
    return newMensaje(id, remitenteID, uuid.Nil, canalID, uuid.Nil, contenido, timestamp, archivoID)
}

// NewMensajeChatPrivado crea un MensajeServidor para mensajes en un chat privado.
// chatPrivadoID debe ser distinto de uuid.Nil, tanto destinoUsuarioID como canalID deben ser uuid.Nil.
func NewMensajeChatPrivado(
    id, remitenteID, chatPrivadoID uuid.UUID,
    contenido string,
    timestamp time.Time,
    archivoID uuid.UUID,
) (*MensajeServidor, error) {
    // Validar explícitamente que se está pasando un chatPrivadoID válido
    if chatPrivadoID == uuid.Nil {
        return nil, ErrMensajeChatPrivadoIDNil
    }
    return newMensaje(id, remitenteID, uuid.Nil, uuid.Nil, chatPrivadoID, contenido, timestamp, archivoID)
}

// newMensaje valida invariantes comunes y construye el MensajeServidor.
func newMensaje(
    id, remitenteID, destinoUsuarioID, canalID, chatPrivadoID uuid.UUID,
    contenido string,
    timestamp time.Time,
    archivoID uuid.UUID,
) (*MensajeServidor, error) {
    if id == uuid.Nil {
        return nil, ErrMensajeIDNil
    }
    if remitenteID == uuid.Nil {
        return nil, ErrRemitenteIDNil
    }
    if contenido == "" {
        return nil, ErrContenidoVacio
    }
    if timestamp.IsZero() {
        return nil, ErrTimestampZero
    }
    hasDestino := destinoUsuarioID != uuid.Nil
    hasCanal := canalID != uuid.Nil
    hasChatPrivado := chatPrivadoID != uuid.Nil
    
    // Si es un mensaje de chat privado, no necesitamos destino ni canal
    if hasChatPrivado {
        // Pero no debería tener destino ni canal al mismo tiempo
        if hasDestino || hasCanal {
            return nil, ErrDestinoYCanalAmbos
        }
    } else {
        // Si no es chat privado, debe tener destino o canal pero no ambos
        switch {
        case hasDestino && hasCanal:
            return nil, ErrDestinoYCanalAmbos
        case !hasDestino && !hasCanal:
            return nil, ErrSinDestinoNiCanal
        }
    }
    return &MensajeServidor{
        id:               id,
        remitenteID:      remitenteID,
        destinoUsuarioID: destinoUsuarioID,
        canalID:          canalID,
        chatPrivadoID:    chatPrivadoID,
        contenido:        contenido,
        timestamp:        timestamp,
        archivoID:        archivoID,
    }, nil
}

// Getters
func (m *MensajeServidor) ID() uuid.UUID           { return m.id }
func (m *MensajeServidor) RemitenteID() uuid.UUID { return m.remitenteID }
func (m *MensajeServidor) DestinoUsuarioID() uuid.UUID { return m.destinoUsuarioID }
func (m *MensajeServidor) CanalID() uuid.UUID          { return m.canalID }
func (m *MensajeServidor) ChatPrivadoID() uuid.UUID    { return m.chatPrivadoID }
func (m *MensajeServidor) Contenido() string           { return m.contenido }
func (m *MensajeServidor) Timestamp() time.Time        { return m.timestamp }
func (m *MensajeServidor) ArchivoID() uuid.UUID        { return m.archivoID }
