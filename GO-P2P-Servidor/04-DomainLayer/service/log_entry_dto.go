package service

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// EventoTipo define los tipos de eventos para el sistema de auditoría
type EventoTipo string

const (
	EventoLogin       EventoTipo = "LOGIN"
	EventoLogout      EventoTipo = "LOGOUT"
	EventoRegistro    EventoTipo = "REGISTRO"
	EventoMensaje     EventoTipo = "MENSAJE"
	EventoArchivo     EventoTipo = "ARCHIVO"
	EventoCanal       EventoTipo = "CANAL"
	EventoInvitacion  EventoTipo = "INVITACION"
	EventoSistema     EventoTipo = "SISTEMA"
)

// LogFilter define los criterios de filtrado para consulta de logs
type LogFilter struct {
	Desde      time.Time  `json:"desde"`
	Hasta      time.Time  `json:"hasta"`
	TiposEvento []EventoTipo `json:"tiposEvento,omitempty"`
	UsuarioID   *uuid.UUID  `json:"usuarioID,omitempty"`
}

// LogEntryDTO representa una entrada de log simplificada para transferencia de datos
type LogEntryDTO struct {
	ID        uuid.UUID  `json:"id"`
	Timestamp time.Time  `json:"timestamp"`
	Tipo      EventoTipo `json:"tipo"`
	Detalle   string     `json:"detalle"`
	UsuarioID *uuid.UUID `json:"usuarioID,omitempty"`
	IP        string     `json:"ip,omitempty"`
}

// MapLogEntryToDTO convierte un modelo LogEntry a un DTO
func MapLogEntryToDTO(l *model.LogEntry) *LogEntryDTO {
	return &LogEntryDTO{
		ID:        l.ID(),
		Timestamp: l.Timestamp(),
		Tipo:      EventoTipo(l.Tipo()),
		Detalle:   l.Detalle(),
		UsuarioID: l.UsuarioID(),
		IP:        l.IP(),
	}
}
