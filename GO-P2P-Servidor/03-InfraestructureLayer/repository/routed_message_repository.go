package repository

import (
	"context"
	"dao"
	"github.com/google/uuid"
	"model"
)

// RoutedMessageRepository implementa la interfaz IRoutedMessageRepository
type RoutedMessageRepository struct {
	dao *dao.MensajeEnrutadoDAO
}

// NewRoutedMessageRepository crea una nueva instancia del repositorio
func NewRoutedMessageRepository(dao *dao.MensajeEnrutadoDAO) *RoutedMessageRepository {
	return &RoutedMessageRepository{
		dao: dao,
	}
}

// Save persiste un mensaje enrutado
func (r *RoutedMessageRepository) Save(ctx context.Context, message *model.RoutedMessage) error {
	// El DAO actual no utiliza context, simplemente pasamos el mensaje
	return r.dao.Crear(message)
}

// FindByMessage busca mensajes enrutados por ID de mensaje
// Nota: Hay una discrepancia entre la interfaz (que retorna un slice) y el DAO (que retorna un único mensaje)
// Esto se maneja devolviendo un slice que contiene el mensaje único o está vacío
func (r *RoutedMessageRepository) FindByMessage(ctx context.Context, messageID uuid.UUID) ([]*model.RoutedMessage, error) {
	message, err := r.dao.BuscarPorMensajeID(messageID)
	if err != nil {
		return nil, err
	}

	if message == nil {
		return []*model.RoutedMessage{}, nil
	}

	return []*model.RoutedMessage{message}, nil
}

// DeleteByMessage elimina un mensaje enrutado por su ID de mensaje
func (r *RoutedMessageRepository) DeleteByMessage(ctx context.Context, messageID uuid.UUID) error {
	return r.dao.Eliminar(messageID)
}
