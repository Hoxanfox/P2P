package repository

import (
	"context"
	"dao"
	"github.com/google/uuid"
	"model"
)

// PeerRepository implements the IPeerRepository interface using NodoDAO
type PeerRepository struct {
	nodoDAO *dao.NodoDAO
}

// NewPeerRepository creates a new PeerRepository instance
func NewPeerRepository(nodoDAO *dao.NodoDAO) *PeerRepository {
	return &PeerRepository{nodoDAO: nodoDAO}
}

// Save persists a peer to the database
func (r *PeerRepository) Save(ctx context.Context, p *model.Peer) error {
	// Context isn't used in the DAO, but we could add that functionality later
	return r.nodoDAO.Guardar(p)
}

// Update updates an existing peer in the database
func (r *PeerRepository) Update(ctx context.Context, p *model.Peer) error {
	return r.nodoDAO.Actualizar(p)
}

// Delete removes a peer from the database by ID
func (r *PeerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.nodoDAO.Eliminar(id)
}

// FindByID retrieves a peer by its ID
func (r *PeerRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.Peer, error) {
	return r.nodoDAO.BuscarPorID(id)
}

// ListAll retrieves all peers from the database
func (r *PeerRepository) ListAll(ctx context.Context) ([]*model.Peer, error) {
	return r.nodoDAO.BuscarTodos()
}

// ListByState retrieves all peers with a specific state
func (r *PeerRepository) ListByState(ctx context.Context, state model.NodoEstado) ([]*model.Peer, error) {
	return r.nodoDAO.BuscarPorEstado(state)
}
