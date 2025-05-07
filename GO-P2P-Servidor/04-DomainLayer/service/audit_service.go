package service

import (
	"github.com/google/uuid"
)

// AuditService define las operaciones para registro y consulta de logs de eventos
type AuditService interface {
	// LogEvent registra un evento en el sistema de auditoría
	LogEvent(
		tipo EventoTipo,
		detalle string,
		usuarioID *uuid.UUID,
	) error

	// ListLogs lista los logs según el filtro especificado
	ListLogs(
		filtro LogFilter, // rango fechas, tipo, usuario
	) ([]LogEntryDTO, error)
}
