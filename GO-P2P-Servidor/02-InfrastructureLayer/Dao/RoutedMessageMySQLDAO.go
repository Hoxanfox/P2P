package Dao

import (
	"database/sql"
	"github.com/Hoxanfox/P2P/GO-P2P-Servidor/04-DomainLayer/model"
	"github.com/google/uuid"
	"time"
)

// RoutedMessageMySQLDAO handles database operations for RoutedMessage entities
type RoutedMessageMySQLDAO struct {
	db *sql.DB
}

// NewRoutedMessageMySQLDAO creates a new RoutedMessageMySQLDAO instance
func NewRoutedMessageMySQLDAO(db *sql.DB) *RoutedMessageMySQLDAO {
	return &RoutedMessageMySQLDAO{db: db}
}

// Create persists a new routed message to the database
func (dao *RoutedMessageMySQLDAO) Create(message *model.RoutedMessage) error {
	query := `INSERT INTO routed_messages (mensaje_id, nodo_destino_id, enruta_at)
              VALUES (?, ?, ?)`

	_, err := dao.db.Exec(
		query,
		message.MensajeID().String(),
		message.NodoDestinoID().String(),
		message.EnrutaAt(),
	)

	return err
}

// FindByMensajeID retrieves a routed message by its mensaje ID
func (dao *RoutedMessageMySQLDAO) FindByMensajeID(mensajeID uuid.UUID) (*model.RoutedMessage, error) {
	query := `SELECT mensaje_id, nodo_destino_id, enruta_at
              FROM routed_messages WHERE mensaje_id = ?`

	row := dao.db.QueryRow(query, mensajeID.String())
	return dao.scanRoutedMessage(row)
}

// FindAll retrieves all routed messages from the database
func (dao *RoutedMessageMySQLDAO) FindAll() ([]*model.RoutedMessage, error) {
	query := `SELECT mensaje_id, nodo_destino_id, enruta_at
              FROM routed_messages`

	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleRoutedMessages(rows)
}

// FindByNodoDestinoID retrieves all routed messages for a specific destination node
func (dao *RoutedMessageMySQLDAO) FindByNodoDestinoID(nodoDestinoID uuid.UUID) ([]*model.RoutedMessage, error) {
	query := `SELECT mensaje_id, nodo_destino_id, enruta_at
              FROM routed_messages WHERE nodo_destino_id = ?`

	rows, err := dao.db.Query(query, nodoDestinoID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleRoutedMessages(rows)
}

// FindByTimeRange retrieves routed messages within a specific time range
func (dao *RoutedMessageMySQLDAO) FindByTimeRange(start, end time.Time) ([]*model.RoutedMessage, error) {
	query := `SELECT mensaje_id, nodo_destino_id, enruta_at
              FROM routed_messages
              WHERE enruta_at BETWEEN ? AND ?`

	rows, err := dao.db.Query(query, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultipleRoutedMessages(rows)
}

// Delete removes a routed message from the database
func (dao *RoutedMessageMySQLDAO) Delete(mensajeID uuid.UUID) error {
	query := `DELETE FROM routed_messages WHERE mensaje_id = ?`
	_, err := dao.db.Exec(query, mensajeID.String())
	return err
}

// Helper method to scan a row into a routed message
func (dao *RoutedMessageMySQLDAO) scanRoutedMessage(row *sql.Row) (*model.RoutedMessage, error) {
	var (
		mensajeIDStr, nodoDestinoIDStr string
		enrutaAt                       time.Time
	)

	if err := row.Scan(&mensajeIDStr, &nodoDestinoIDStr, &enrutaAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Routed message not found
		}
		return nil, err
	}

	mensajeID, err := uuid.Parse(mensajeIDStr)
	if err != nil {
		return nil, err
	}

	nodoDestinoID, err := uuid.Parse(nodoDestinoIDStr)
	if err != nil {
		return nil, err
	}

	return model.NewRoutedMessage(mensajeID, nodoDestinoID, enrutaAt)
}

// Helper method to scan multiple routed messages
func (dao *RoutedMessageMySQLDAO) scanMultipleRoutedMessages(rows *sql.Rows) ([]*model.RoutedMessage, error) {
	var messages []*model.RoutedMessage

	for rows.Next() {
		var (
			mensajeIDStr, nodoDestinoIDStr string
			enrutaAt                       time.Time
		)

		if err := rows.Scan(&mensajeIDStr, &nodoDestinoIDStr, &enrutaAt); err != nil {
			return nil, err
		}

		mensajeID, err := uuid.Parse(mensajeIDStr)
		if err != nil {
			return nil, err
		}

		nodoDestinoID, err := uuid.Parse(nodoDestinoIDStr)
		if err != nil {
			return nil, err
		}

		message, err := model.NewRoutedMessage(mensajeID, nodoDestinoID, enrutaAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
