package service

// AdminService define las operaciones para supervisión y estadísticas agregadas
type AdminService interface {
	// ListRegisteredUsers lista todos los usuarios registrados en el sistema
	ListRegisteredUsers() ([]UsuarioDTO, error)
	
	// ListConnectedUsers lista todos los usuarios actualmente conectados
	ListConnectedUsers() ([]UsuarioDTO, error)
	
	// ListChannelsSummary obtiene un resumen de todos los canales
	ListChannelsSummary() ([]ChannelSummaryDTO, error)
	
	// ListLogs obtiene los registros de eventos del sistema
	ListLogs() ([]LogEntryDTO, error)
}
