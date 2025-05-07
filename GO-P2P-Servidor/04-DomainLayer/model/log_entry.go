package model

import (
    "errors"
    "time"

    "github.com/google/uuid"
)

// Errores de validación para LogEntry
var (
    ErrLogEntryIDNil        = errors.New("id de log inválido")
    ErrEventoTipoInvalido   = errors.New("tipo de evento inválido")
    ErrDetalleVacio         = errors.New("detalle de log vacío")
    ErrTimestampLogZero     = errors.New("timestamp de log no puede ser cero")
)

// LogEntry representa una entrada de auditoría o registro de eventos.
type LogEntry struct {
    id          uuid.UUID
    tipoEvento  EventoTipo
    detalle     string
    timestamp   time.Time
    usuarioID   uuid.UUID // Opcional: puede ser uuid.Nil si no aplica
}

// NewLogEntry crea un LogEntry validando sus invariantes.
func NewLogEntry(
    id uuid.UUID,
    tipoEvento EventoTipo,
    detalle string,
    timestamp time.Time,
    usuarioID uuid.UUID,
) (*LogEntry, error) {
    if id == uuid.Nil {
        return nil, ErrLogEntryIDNil
    }
    if !tipoEvento.Valid() {
        return nil, ErrEventoTipoInvalido
    }
    if detalle == "" {
        return nil, ErrDetalleVacio
    }
    if timestamp.IsZero() {
        return nil, ErrTimestampLogZero
    }
    return &LogEntry{
        id:         id,
        tipoEvento: tipoEvento,
        detalle:    detalle,
        timestamp:  timestamp,
        usuarioID:  usuarioID,
    }, nil
}

// Getters
func (l *LogEntry) ID() uuid.UUID         { return l.id }
func (l *LogEntry) TipoEvento() EventoTipo { return l.tipoEvento }
func (l *LogEntry) Detalle() string        { return l.detalle }
func (l *LogEntry) Timestamp() time.Time   { return l.timestamp }
func (l *LogEntry) UsuarioID() uuid.UUID   { return l.usuarioID }
