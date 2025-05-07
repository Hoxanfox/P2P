package model

// EventoTipo categoriza los distintos tipos de eventos para logs y auditoría.
type EventoTipo string

const (
	EventoLogin   EventoTipo = "LOGIN"
	EventoMensaje EventoTipo = "MENSAJE"
	EventoArchivo EventoTipo = "ARCHIVO"
	EventoCanal   EventoTipo = "CANAL"
)

// Valid comprueba que el EventoTipo sea uno de los valores admitidos.
func (e EventoTipo) Valid() bool {
	switch e {
	case EventoLogin,
		EventoMensaje,
		EventoArchivo,
		EventoCanal:
		return true
	default:
		return false
	}
}
