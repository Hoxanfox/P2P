package Dao
<<<<<<< HEAD
=======

import (
	"database/sql"

	"github.com/P2P/GO-P2P-SERVIDOR/04-DomainLayer/model"
	"github.com/google/uuid"
)

// ChannelMySQLDAO handles database operations for CanalServidor entities
type ChannelMySQLDAO struct {
	db *sql.DB
}

// NewChannelMySQLDAO creates a new ChannelMySQLDAO instance
func NewChannelMySQLDAO(db *sql.DB) *ChannelMySQLDAO {
	return &ChannelMySQLDAO{db: db}
}

// Create persists a new channel to the database
func (dao *ChannelMySQLDAO) Create(channel *model.CanalServidor) error {
	query := `INSERT INTO canales (id, nombre, descripcion, tipo)
              VALUES (?, ?, ?, ?)`

	_, err := dao.db.Exec(
		query,
		channel.ID().String(),
		channel.Nombre(),
		channel.Descripcion(),
		channel.Tipo().String(), // Assuming CanalTipo has a String() method
	)

	return err
}

// FindByID retrieves a channel by its ID
func (dao *ChannelMySQLDAO) FindByID(id uuid.UUID) (*model.CanalServidor, error) {
	query := `SELECT id, nombre, descripcion, tipo
              FROM canales WHERE id = ?`

	row := dao.db.QueryRow(query, id.String())
	return dao.scanChannel(row)
}

// Update updates an existing channel in the database
func (dao *ChannelMySQLDAO) Update(channel *model.CanalServidor) error {
	query := `UPDATE canales 
              SET nombre = ?, descripcion = ?, tipo = ?
              WHERE id = ?`

	_, err := dao.db.Exec(
		query,
		channel.Nombre(),
		channel.Descripcion(),
		channel.Tipo().String(), // Assuming CanalTipo has a String() method
		channel.ID().String(),
	)

	return err
}

// Delete removes a channel from the database
func (dao *ChannelMySQLDAO) Delete(id uuid.UUID) error {
	query := `DELETE FROM canales WHERE id = ?`
	_, err := dao.db.Exec(query, id.String())
	return err
}

// FindAll retrieves all channels from the database
func (dao *ChannelMySQLDAO) FindAll() ([]*model.CanalServidor, error) {
	query := `SELECT id, nombre, descripcion, tipo 
              FROM canales`

	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var channels []*model.CanalServidor
	for rows.Next() {
		channel, err := dao.scanChannelRow(rows)
		if err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return channels, nil
}

// Helper method to scan a row into a channel
func (dao *ChannelMySQLDAO) scanChannel(row *sql.Row) (*model.CanalServidor, error) {
	var (
		idStr, nombre, descripcion, tipoStr string
	)

	if err := row.Scan(&idStr, &nombre, &descripcion, &tipoStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Channel not found
		}
		return nil, err
	}

	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	tipo, err := model.ParseCanalTipo(tipoStr) // Assuming this function exists
	if err != nil {
		return nil, err
	}

	channel, err := model.NewCanalServidor(parsedID, nombre, descripcion, tipo)
	if err != nil {
		return nil, err
	}

	return channel, nil
}

// Helper method to scan a row from rows.Next()
func (dao *ChannelMySQLDAO) scanChannelRow(rows *sql.Rows) (*model.CanalServidor, error) {
	var (
		idStr, nombre, descripcion, tipoStr string
	)

	if err := rows.Scan(&idStr, &nombre, &descripcion, &tipoStr); err != nil {
		return nil, err
	}

	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	tipo, err := model.ParseCanalTipo(tipoStr) // Assuming this function exists
	if err != nil {
		return nil, err
	}

	return model.NewCanalServidor(parsedID, nombre, descripcion, tipo)
}
>>>>>>> main
