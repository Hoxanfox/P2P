package Dao

import (
	"database/sql"
	"github.com/google/uuid"
	"model"
	"time"
)

// HeartbeatLogMySQLDAO handles database operations for HeartbeatLog entities
type HeartbeatLogMySQLDAO struct {
	db *sql.DB
}

// NewHeartbeatLogMySQLDAO creates a new HeartbeatLogMySQLDAO instance
func NewHeartbeatLogMySQLDAO(db *sql.DB) *HeartbeatLogMySQLDAO {
	return &HeartbeatLogMySQLDAO{db: db}
}

// Save persists a heartbeat log to the database
func (dao *HeartbeatLogMySQLDAO) Save(log *model.HeartbeatLog) error {
	query := `INSERT INTO heartbeat_logs (id, nodo_id, enviado_at, recibido_at)
              VALUES (?, ?, ?, ?)`

	_, err := dao.db.Exec(
		query,
		log.ID().String(),
		log.NodoID().String(),
		log.EnviadoAt(),
		log.RecibidoAt(),
	)

	return err
}

// FindByID retrieves a heartbeat log by its ID
func (dao *HeartbeatLogMySQLDAO) FindByID(id uuid.UUID) (*model.HeartbeatLog, error) {
	query := `SELECT id, nodo_id, enviado_at, recibido_at
              FROM heartbeat_logs WHERE id = ?`

	row := dao.db.QueryRow(query, id.String())
	return dao.scanHeartbeatLog(row)
}

// FindAll retrieves all heartbeat logs from the database
func (dao *HeartbeatLogMySQLDAO) FindAll() ([]*model.HeartbeatLog, error) {
	query := `SELECT id, nodo_id, enviado_at, recibido_at
              FROM heartbeat_logs`

	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleHeartbeatLogs(rows)
}

// FindByNodoID retrieves all heartbeat logs for a specific node
func (dao *HeartbeatLogMySQLDAO) FindByNodoID(nodoID uuid.UUID) ([]*model.HeartbeatLog, error) {
	query := `SELECT id, nodo_id, enviado_at, recibido_at
              FROM heartbeat_logs WHERE nodo_id = ?`

	rows, err := dao.db.Query(query, nodoID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleHeartbeatLogs(rows)
}

// FindByTimeRange retrieves heartbeat logs within a specific time range
func (dao *HeartbeatLogMySQLDAO) FindByTimeRange(start, end time.Time) ([]*model.HeartbeatLog, error) {
	query := `SELECT id, nodo_id, enviado_at, recibido_at
              FROM heartbeat_logs 
              WHERE enviado_at BETWEEN ? AND ?`

	rows, err := dao.db.Query(query, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleHeartbeatLogs(rows)
}

// Delete removes a heartbeat log from the database
func (dao *HeartbeatLogMySQLDAO) Delete(id uuid.UUID) error {
	query := `DELETE FROM heartbeat_logs WHERE id = ?`
	_, err := dao.db.Exec(query, id.String())
	return err
}

// Helper method to scan a row into a heartbeat log
func (dao *HeartbeatLogMySQLDAO) scanHeartbeatLog(row *sql.Row) (*model.HeartbeatLog, error) {
	var (
		idStr, nodoIDStr      string
		enviadoAt, recibidoAt time.Time
	)

	if err := row.Scan(&idStr, &nodoIDStr, &enviadoAt, &recibidoAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Heartbeat log not found
		}
		return nil, err
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	nodoID, err := uuid.Parse(nodoIDStr)
	if err != nil {
		return nil, err
	}

	return model.NewHeartbeatLog(id, nodoID, enviadoAt, recibidoAt)
}

// Helper method to scan multiple heartbeat logs
func (dao *HeartbeatLogMySQLDAO) scanMultipleHeartbeatLogs(rows *sql.Rows) ([]*model.HeartbeatLog, error) {
	var logs []*model.HeartbeatLog

	for rows.Next() {
		var (
			idStr, nodoIDStr      string
			enviadoAt, recibidoAt time.Time
		)

		if err := rows.Scan(&idStr, &nodoIDStr, &enviadoAt, &recibidoAt); err != nil {
			return nil, err
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		nodoID, err := uuid.Parse(nodoIDStr)
		if err != nil {
			return nil, err
		}

		log, err := model.NewHeartbeatLog(id, nodoID, enviadoAt, recibidoAt)
		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}
