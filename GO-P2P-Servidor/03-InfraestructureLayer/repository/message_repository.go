package repository

import (
	"context"
	"dao"
	"errors"
	"github.com/google/uuid"
	"model"
)

// MessageRepository implements the IMessageRepository interface using MensajeDAO
type MessageRepository struct {
	mensajeDAO *dao.MensajeDAO
}

// NewMessageRepository creates a new MessageRepository instance
func NewMessageRepository(mensajeDAO *dao.MensajeDAO) *MessageRepository {
	return &MessageRepository{mensajeDAO: mensajeDAO}
}

// Save persists a message to the database
func (r *MessageRepository) Save(ctx context.Context, m *model.MensajeServidor) error {
	return r.mensajeDAO.Crear(m)
}

// Update updates an existing message
// Note: MensajeDAO doesn't have an update method, so this is a stub
func (r *MessageRepository) Update(ctx context.Context, m *model.MensajeServidor) error {
	// This would require first deleting the message and then creating it again
	// or implementing a proper update method in MensajeDAO
	return errors.New("update operation not supported by underlying DAO")
}

// Delete removes a message from the database
func (r *MessageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.mensajeDAO.Eliminar(id)
}

// FindByID retrieves a message by its ID
func (r *MessageRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.MensajeServidor, error) {
	return r.mensajeDAO.BuscarPorID(id)
}

// FindByChannel retrieves messages from a specific channel
func (r *MessageRepository) FindByChannel(ctx context.Context, channelID uuid.UUID, limit, offset int) ([]*model.MensajeServidor, error) {
	// Note: MensajeDAO.BuscarPorCanalID doesn't support offset,
	// we'd need to enhance the DAO to support pagination properly
	return r.mensajeDAO.BuscarPorCanalID(channelID, limit)
}

// FindByUser retrieves messages sent by a specific user
func (r *MessageRepository) FindByUser(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*model.MensajeServidor, error) {
	// MensajeDAO doesn't provide a direct method to find messages by user
	// This would need to be implemented in the DAO layer
	return nil, errors.New("findByUser not implemented in the underlying DAO")
}

// FindDirect retrieves direct messages between two users
func (r *MessageRepository) FindDirect(ctx context.Context, a, b uuid.UUID, limit, offset int) ([]*model.MensajeServidor, error) {
	// Note: MensajeDAO.BuscarMensajesDirectos doesn't support offset
	return r.mensajeDAO.BuscarMensajesDirectos(a, b, limit)
}
