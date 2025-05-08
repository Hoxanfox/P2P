package service

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// LogFilter define los criterios de filtrado para consulta de logs
type LogFilter struct {
	Desde      time.Time       // Fecha de inicio para el filtro
	Hasta      time.Time       // Fecha de fin para el filtro
	TiposEvento []model.EventoTipo // Tipos de evento a incluir (opcional)
	UsuarioID   *uuid.UUID       // ID del usuario para filtrar (opcional)
}

// AuditService define las operaciones para registro y consulta de logs de eventos
type AuditService interface {
	// LogEvent registra un evento en el sistema de auditoría
	LogEvent(
		tipo model.EventoTipo,
		detalle string,
		usuarioID *uuid.UUID,
	) error

	// ListLogs lista los logs según el filtro especificado
	ListLogs(
		filtro LogFilter, // rango fechas, tipo, usuario
	) ([]*model.LogEntry, error)
}
