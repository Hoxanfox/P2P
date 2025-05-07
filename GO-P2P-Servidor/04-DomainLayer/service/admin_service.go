package service

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// ChannelSummary representa un resumen de información de canal para administradores
type ChannelSummary struct {
	ID               uuid.UUID // ID del canal
	Nombre           string    // Nombre del canal
	NumMiembros      int       // Número de miembros en el canal
	NumMensajes      int       // Número de mensajes en el canal
	UltimoMensaje    time.Time // Fecha del último mensaje
	FechaCreacion    time.Time // Fecha de creación del canal
	TipoCanal        string    // Tipo de canal (público, privado, etc)
	ArchivosCompartidos int     // Número de archivos compartidos
}

// AdminService define las operaciones para supervisión y estadísticas agregadas
type AdminService interface {
	// ListRegisteredUsers lista todos los usuarios registrados en el sistema
	ListRegisteredUsers() ([]*model.UsuarioServidor, error)
	
	// ListConnectedUsers lista todos los usuarios actualmente conectados
	ListConnectedUsers() ([]*model.UsuarioServidor, error)
	
	// ListChannelsSummary obtiene un resumen de todos los canales
	ListChannelsSummary() ([]ChannelSummary, error)
	
	// ListLogs obtiene los registros de eventos del sistema
	ListLogs() ([]*model.LogEntry, error)
}
