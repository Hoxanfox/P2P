package Dao

import (
	"database/sql"
	"github.com/google/uuid"
	"model"
)

// PeerMySQLDAO handles database operations for Peer entities
type PeerMySQLDAO struct {
	db *sql.DB
}

// NewPeerMySQLDAO creates a new PeerMySQLDAO instance
func NewPeerMySQLDAO(db *sql.DB) *PeerMySQLDAO {
	return &PeerMySQLDAO{db: db}
}

// Save persists a peer to the database
func (dao *PeerMySQLDAO) Save(peer *model.Peer) error {
	query := `INSERT INTO peers (id_nodo, direccion, estado)
              VALUES (?, ?, ?)`

	_, err := dao.db.Exec(
		query,
		peer.IDNodo().String(),
		peer.Direccion(),
		int(peer.Estado()),
	)

	return err
}

// FindByID retrieves a peer by its ID
func (dao *PeerMySQLDAO) FindByID(id uuid.UUID) (*model.Peer, error) {
	query := `SELECT id_nodo, direccion, estado
              FROM peers WHERE id_nodo = ?`

	row := dao.db.QueryRow(query, id.String())
	return dao.scanPeer(row)
}

// FindAll retrieves all peers from the database
func (dao *PeerMySQLDAO) FindAll() ([]*model.Peer, error) {
	query := `SELECT id_nodo, direccion, estado
              FROM peers`

	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultiplePeers(rows)
}

// Update updates an existing peer
func (dao *PeerMySQLDAO) Update(peer *model.Peer) error {
	query := `UPDATE peers SET direccion = ?, estado = ?
              WHERE id_nodo = ?`

	_, err := dao.db.Exec(
		query,
		peer.Direccion(),
		int(peer.Estado()),
		peer.IDNodo().String(),
	)

	return err
}

// Delete removes a peer from the database
func (dao *PeerMySQLDAO) Delete(id uuid.UUID) error {
	query := `DELETE FROM peers WHERE id_nodo = ?`
	_, err := dao.db.Exec(query, id.String())
	return err
}

// FindByEstado retrieves all peers with a specific state
func (dao *PeerMySQLDAO) FindByEstado(estado model.NodoEstado) ([]*model.Peer, error) {
	query := `SELECT id_nodo, direccion, estado
              FROM peers WHERE estado = ?`

	rows, err := dao.db.Query(query, int(estado))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMultiplePeers(rows)
}

// Helper method to scan a row into a peer
func (dao *PeerMySQLDAO) scanPeer(row *sql.Row) (*model.Peer, error) {
	var (
		idStr       string
		direccion   string
		estadoValue int
	)

	if err := row.Scan(&idStr, &direccion, &estadoValue); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Peer not found
		}
		return nil, err
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	estado := model.NodoEstado(estadoValue)

	return model.NewPeer(id, direccion, estado)
}

// Helper method to scan multiple peers
func (dao *PeerMySQLDAO) scanMultiplePeers(rows *sql.Rows) ([]*model.Peer, error) {
	var peers []*model.Peer

	for rows.Next() {
		var (
			idStr       string
			direccion   string
			estadoValue int
		)

		if err := rows.Scan(&idStr, &direccion, &estadoValue); err != nil {
			return nil, err
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		estado := model.NodoEstado(estadoValue)

		peer, err := model.NewPeer(id, direccion, estado)
		if err != nil {
			return nil, err
		}

		peers = append(peers, peer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return peers, nil
}
