package repository

import (
	"context"
	"dao"
	"github.com/google/uuid"
	"model"
)

// ReplicaEventRepository implementa la interfaz IReplicaEventRepository
type ReplicaEventRepository struct {
	dao *dao.ReplicaEventMySQLDAO
}

// NewReplicaEventRepository crea una nueva instancia del repositorio
func NewReplicaEventRepository(dao *dao.ReplicaEventMySQLDAO) *ReplicaEventRepository {
	return &ReplicaEventRepository{
		dao: dao,
	}
}

// Save persiste un evento de réplica
func (r *ReplicaEventRepository) Save(ctx context.Context, e *model.ReplicaEvent) error {
	return r.dao.Create(e)
}

// Delete elimina un evento de réplica por su ID
func (r *ReplicaEventRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.dao.Delete(id)
}

// FindByID busca un evento de réplica por su ID
func (r *ReplicaEventRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.ReplicaEvent, error) {
	return r.dao.FindByID(id)
}

// ListPending retorna los eventos de réplica pendientes de procesamiento
// Nota: Esta implementación asume que se necesita extender el DAO y la tabla en la BD
// con una columna 'processed' para soportar esta funcionalidad
func (r *ReplicaEventRepository) ListPending(ctx context.Context) ([]*model.ReplicaEvent, error) {
	// TODO: Implementar cuando el DAO soporte consultar eventos por estado de procesamiento
	// Por ahora, retorna todos los eventos (comportamiento temporal)
	return r.dao.FindAll()
}

// MarkProcessed marca un evento de réplica como procesado
// Nota: Esta implementación asume que se necesita extender el DAO y la tabla en la BD
// con una columna 'processed' para soportar esta funcionalidad
func (r *ReplicaEventRepository) MarkProcessed(ctx context.Context, id uuid.UUID) error {
	// TODO: Implementar cuando el DAO soporte actualizar el estado de procesamiento
	// Por ahora no hace nada (comportamiento temporal)
	return nil
}
