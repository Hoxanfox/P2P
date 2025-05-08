package repository

import (
	"context"
	"dao"
	"github.com/google/uuid"
	"model"
)

// FileRepository implementa la interfaz IFileRepository
type FileRepository struct {
	dao *dao.ArchivoDAO
}

// NewFileRepository crea una nueva instancia del repositorio de archivos
func NewFileRepository(dao *dao.ArchivoDAO) *FileRepository {
	return &FileRepository{
		dao: dao,
	}
}

// Save persiste los metadatos de un archivo
func (r *FileRepository) Save(ctx context.Context, f *model.ArchivoMetadata) error {
	// Utiliza el método Crear del DAO
	return r.dao.Crear(f)
}

// Delete elimina un archivo por su ID
func (r *FileRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.dao.Delete(id)
}

// FindByID busca un archivo por su ID
func (r *FileRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.ArchivoMetadata, error) {
	return r.dao.GetByID(id)
}

// FindByMessage busca archivos asociados a un mensaje específico
// Nota: Esta implementación requiere extender el DAO y la tabla en la BD
// para soportar la asociación de archivos con mensajes
func (r *FileRepository) FindByMessage(ctx context.Context, messageID uuid.UUID) ([]*model.ArchivoMetadata, error) {
	// TODO: Implementar cuando el DAO soporte consultar archivos por mensaje
	// Esta funcionalidad necesita una extensión del DAO para buscar por relación con mensajes
	return []*model.ArchivoMetadata{}, nil
}
