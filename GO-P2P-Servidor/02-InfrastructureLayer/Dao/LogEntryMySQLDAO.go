package Dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"model"
)

// LogEntryMySQLDAO handles database operations for LogEntry entities
type LogEntryMySQLDAO struct {
	db *sql.DB
}

// NewLogEntryMySQLDAO creates a new LogEntryMySQLDAO instance
func NewLogEntryMySQLDAO(db *sql.DB) *LogEntryMySQLDAO {
	return &LogEntryMySQLDAO{db: db}
}

// Create persists a new log entry to the database
func (dao *LogEntryMySQLDAO) Create(log *model.LogEntry) error {
	query := `INSERT INTO log_entries (id, tipo_evento, detalle, timestamp, usuario_id)
              VALUES (?, ?, ?, ?, ?)`

	_, err := dao.db.Exec(
		query,
		log.ID().String(),
		int(log.TipoEvento()), // Assuming EventoTipo is an integer type
		log.Detalle(),
		log.Timestamp(),
		log.UsuarioID().String(),
	)

	return err
}

// FindByID retrieves a log entry by its ID
func (dao *LogEntryMySQLDAO) FindByID(id uuid.UUID) (*model.LogEntry, error) {
	query := `SELECT id, tipo_evento, detalle, timestamp, usuario_id
              FROM log_entries WHERE id = ?`

	row := dao.db.QueryRow(query, id.String())
	return dao.scanLogEntry(row)
}

// FindByUsuarioID retrieves all log entries for a specific user
func (dao *LogEntryMySQLDAO) FindByUsuarioID(usuarioID uuid.UUID) ([]*model.LogEntry, error) {
	query := `SELECT id, tipo_evento, detalle, timestamp, usuario_id
              FROM log_entries WHERE usuario_id = ?`

	rows, err := dao.db.Query(query, usuarioID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleLogEntries(rows)
}

// FindByTipoEvento retrieves all log entries of a specific event type
func (dao *LogEntryMySQLDAO) FindByTipoEvento(tipoEvento model.EventoTipo) ([]*model.LogEntry, error) {
	query := `SELECT id, tipo_evento, detalle, timestamp, usuario_id
              FROM log_entries WHERE tipo_evento = ?`

	rows, err := dao.db.Query(query, int(tipoEvento))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleLogEntries(rows)
}

// FindAll retrieves all log entries from the database
func (dao *LogEntryMySQLDAO) FindAll() ([]*model.LogEntry, error) {
	query := `SELECT id, tipo_evento, detalle, timestamp, usuario_id
              FROM log_entries`

	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleLogEntries(rows)
}

// Delete removes a log entry from the database
func (dao *LogEntryMySQLDAO) Delete(id uuid.UUID) error {
	query := `DELETE FROM log_entries WHERE id = ?`
	_, err := dao.db.Exec(query, id.String())
	return err
}

// Helper method to scan a row into a log entry
func (dao *LogEntryMySQLDAO) scanLogEntry(row *sql.Row) (*model.LogEntry, error) {
	var (
		idStr, usuarioIDStr string
		tipoEventoInt       int
		detalle             string
		timestamp           time.Time
	)

	if err := row.Scan(&idStr, &tipoEventoInt, &detalle, &timestamp, &usuarioIDStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Log entry not found
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

	// Convert tipoEventoInt to EventoTipo
	tipoEvento := model.EventoTipo(tipoEventoInt)

	return model.NewLogEntry(id, tipoEvento, detalle, timestamp, usuarioID)
}

// Helper method to scan multiple log entries
func (dao *LogEntryMySQLDAO) scanMultipleLogEntries(rows *sql.Rows) ([]*model.LogEntry, error) {
	var entries []*model.LogEntry

	for rows.Next() {
		var (
			idStr, usuarioIDStr string
			tipoEventoInt       int
			detalle             string
			timestamp           time.Time
		)

		if err := rows.Scan(&idStr, &tipoEventoInt, &detalle, &timestamp, &usuarioIDStr); err != nil {
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

		tipoEvento := model.EventoTipo(tipoEventoInt)

		entry, err := model.NewLogEntry(id, tipoEvento, detalle, timestamp, usuarioID)
		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
