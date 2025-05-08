package Dao

import (
	"database/sql"
	"github.com/google/uuid"
	"model"
	"time"
)

// ReplicaEventMySQLDAO handles database operations for ReplicaEvent entities
type ReplicaEventMySQLDAO struct {
	db *sql.DB
}

// NewReplicaEventMySQLDAO creates a new ReplicaEventMySQLDAO instance
func NewReplicaEventMySQLDAO(db *sql.DB) *ReplicaEventMySQLDAO {
	return &ReplicaEventMySQLDAO{db: db}
}

// Create persists a new replica event to the database
func (dao *ReplicaEventMySQLDAO) Create(event *model.ReplicaEvent) error {
	query := `INSERT INTO replica_events (id, entidad_tipo, entidad_id, evento_at, origen_nodo_id)
              VALUES (?, ?, ?, ?, ?)`

	_, err := dao.db.Exec(
		query,
		event.ID().String(),
		event.EntidadTipo(),
		event.EntidadID().String(),
		event.EventoAt(),
		event.OrigenNodoID().String(),
	)

	return err
}

// FindByID retrieves a replica event by its ID
func (dao *ReplicaEventMySQLDAO) FindByID(id uuid.UUID) (*model.ReplicaEvent, error) {
	query := `SELECT id, entidad_tipo, entidad_id, evento_at, origen_nodo_id
              FROM replica_events WHERE id = ?`

	row := dao.db.QueryRow(query, id.String())
	return dao.scanReplicaEvent(row)
}

// FindAll retrieves all replica events from the database
func (dao *ReplicaEventMySQLDAO) FindAll() ([]*model.ReplicaEvent, error) {
	query := `SELECT id, entidad_tipo, entidad_id, evento_at, origen_nodo_id
              FROM replica_events`

	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleReplicaEvents(rows)
}

// FindByEntidadID retrieves all replica events for a specific entity
func (dao *ReplicaEventMySQLDAO) FindByEntidadID(entidadID uuid.UUID) ([]*model.ReplicaEvent, error) {
	query := `SELECT id, entidad_tipo, entidad_id, evento_at, origen_nodo_id
              FROM replica_events WHERE entidad_id = ?`

	rows, err := dao.db.Query(query, entidadID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleReplicaEvents(rows)
}

// FindByOrigenNodoID retrieves all replica events from a specific origin node
func (dao *ReplicaEventMySQLDAO) FindByOrigenNodoID(origenNodoID uuid.UUID) ([]*model.ReplicaEvent, error) {
	query := `SELECT id, entidad_tipo, entidad_id, evento_at, origen_nodo_id
              FROM replica_events WHERE origen_nodo_id = ?`

	rows, err := dao.db.Query(query, origenNodoID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleReplicaEvents(rows)
}

// FindByEntidadTipo retrieves all replica events of a specific entity type
func (dao *ReplicaEventMySQLDAO) FindByEntidadTipo(entidadTipo string) ([]*model.ReplicaEvent, error) {
	query := `SELECT id, entidad_tipo, entidad_id, evento_at, origen_nodo_id
              FROM replica_events WHERE entidad_tipo = ?`

	rows, err := dao.db.Query(query, entidadTipo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleReplicaEvents(rows)
}

// Delete removes a replica event from the database
func (dao *ReplicaEventMySQLDAO) Delete(id uuid.UUID) error {
	query := `DELETE FROM replica_events WHERE id = ?`
	_, err := dao.db.Exec(query, id.String())
	return err
}

// Helper method to scan a row into a replica event
func (dao *ReplicaEventMySQLDAO) scanReplicaEvent(row *sql.Row) (*model.ReplicaEvent, error) {
	var (
		idStr, entidadIDStr, origenNodoIDStr string
		entidadTipo                          string
		eventoAt                             time.Time
	)

	if err := row.Scan(&idStr, &entidadTipo, &entidadIDStr, &eventoAt, &origenNodoIDStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Replica event not found
		}
		return nil, err
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	entidadID, err := uuid.Parse(entidadIDStr)
	if err != nil {
		return nil, err
	}

	origenNodoID, err := uuid.Parse(origenNodoIDStr)
	if err != nil {
		return nil, err
	}

	return model.NewReplicaEvent(id, entidadTipo, entidadID, eventoAt, origenNodoID)
}

// Helper method to scan multiple replica events
func (dao *ReplicaEventMySQLDAO) scanMultipleReplicaEvents(rows *sql.Rows) ([]*model.ReplicaEvent, error) {
	var events []*model.ReplicaEvent

	for rows.Next() {
		var (
			idStr, entidadIDStr, origenNodoIDStr string
			entidadTipo                          string
			eventoAt                             time.Time
		)

		if err := rows.Scan(&idStr, &entidadTipo, &entidadIDStr, &eventoAt, &origenNodoIDStr); err != nil {
			return nil, err
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		entidadID, err := uuid.Parse(entidadIDStr)
		if err != nil {
			return nil, err
		}

		origenNodoID, err := uuid.Parse(origenNodoIDStr)
		if err != nil {
			return nil, err
		}

		event, err := model.NewReplicaEvent(id, entidadTipo, entidadID, eventoAt, origenNodoID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
