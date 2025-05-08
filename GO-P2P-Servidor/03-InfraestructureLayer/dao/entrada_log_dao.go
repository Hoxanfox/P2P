package dao

import (
	"database/sql"
	"pool"
	"time"

	"github.com/google/uuid"
	"model"
)

// EntradaLogDAO maneja las operaciones de base de datos para entidades LogEntry (EntradaLog)
type EntradaLogDAO struct {
	dbPool *pool.DBConnectionPool
}

// NuevoEntradaLogDAO crea una nueva instancia de EntradaLogDAO
func NuevoEntradaLogDAO(dbPool *pool.DBConnectionPool) *EntradaLogDAO {
	return &EntradaLogDAO{dbPool: dbPool}
}

// Crear persiste una nueva entrada de log en la base de datos
func (dao *EntradaLogDAO) Crear(log *model.LogEntry) error {
	query := `INSERT INTO log_entries (id, tipo_evento, detalle, timestamp, usuario_id)
              VALUES (?, ?, ?, ?, ?)`

	_, err := dao.dbPool.DB().Exec(
		query,
		log.ID().String(),
		string(log.TipoEvento()), // EventoTipo es de tipo string
		log.Detalle(),
		log.Timestamp(),
		log.UsuarioID().String(),
	)

	return err
}

// BuscarPorID recupera una entrada de log por su ID
func (dao *EntradaLogDAO) BuscarPorID(id uuid.UUID) (*model.LogEntry, error) {
	query := `SELECT id, tipo_evento, detalle, timestamp, usuario_id
              FROM log_entries WHERE id = ?`

	row := dao.dbPool.DB().QueryRow(query, id.String())
	return dao.escanearEntradaLog(row)
}

// BuscarPorUsuarioID recupera todas las entradas de log para un usuario específico
func (dao *EntradaLogDAO) BuscarPorUsuarioID(usuarioID uuid.UUID) ([]*model.LogEntry, error) {
	query := `SELECT id, tipo_evento, detalle, timestamp, usuario_id
              FROM log_entries WHERE usuario_id = ?`

	rows, err := dao.dbPool.DB().Query(query, usuarioID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesEntradasLog(rows)
}

// BuscarPorTipoEvento recupera todas las entradas de log de un tipo específico
func (dao *EntradaLogDAO) BuscarPorTipoEvento(tipoEvento model.EventoTipo) ([]*model.LogEntry, error) {
	query := `SELECT id, tipo_evento, detalle, timestamp, usuario_id
              FROM log_entries WHERE tipo_evento = ?`

	rows, err := dao.dbPool.DB().Query(query, string(tipoEvento))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesEntradasLog(rows)
}

// BuscarTodos recupera todas las entradas de log de la base de datos
func (dao *EntradaLogDAO) BuscarTodos() ([]*model.LogEntry, error) {
	query := `SELECT id, tipo_evento, detalle, timestamp, usuario_id
              FROM log_entries`

	rows, err := dao.dbPool.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesEntradasLog(rows)
}

// Eliminar elimina una entrada de log de la base de datos
func (dao *EntradaLogDAO) Eliminar(id uuid.UUID) error {
	query := `DELETE FROM log_entries WHERE id = ?`
	_, err := dao.dbPool.DB().Exec(query, id.String())
	return err
}

// Método auxiliar para escanear una fila en una entrada de log
func (dao *EntradaLogDAO) escanearEntradaLog(row *sql.Row) (*model.LogEntry, error) {
	var (
		idStr, usuarioIDStr string
		tipoEventoStr       string
		detalle             string
		timestamp           time.Time
	)

	if err := row.Scan(&idStr, &tipoEventoStr, &detalle, &timestamp, &usuarioIDStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Entrada de log no encontrada
		}
		return nil, err
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	usuarioID, err := uuid.Parse(usuarioIDStr)
	if err != nil {
		return nil, err
	}

	// Convertir tipoEventoStr a EventoTipo
	tipoEvento := model.EventoTipo(tipoEventoStr)

	return model.NewLogEntry(id, tipoEvento, detalle, timestamp, usuarioID)
}

// Método auxiliar para escanear múltiples entradas de log
func (dao *EntradaLogDAO) escanearMultiplesEntradasLog(rows *sql.Rows) ([]*model.LogEntry, error) {
	var entradas []*model.LogEntry

	for rows.Next() {
		var (
			idStr, usuarioIDStr string
			tipoEventoStr       string
			detalle             string
			timestamp           time.Time
		)

		if err := rows.Scan(&idStr, &tipoEventoStr, &detalle, &timestamp, &usuarioIDStr); err != nil {
			return nil, err
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		usuarioID, err := uuid.Parse(usuarioIDStr)
		if err != nil {
			return nil, err
		}

		tipoEvento := model.EventoTipo(tipoEventoStr)

		entrada, err := model.NewLogEntry(id, tipoEvento, detalle, timestamp, usuarioID)
		if err != nil {
			return nil, err
		}

		entradas = append(entradas, entrada)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entradas, nil
}
