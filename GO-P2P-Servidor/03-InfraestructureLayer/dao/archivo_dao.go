package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"model"
	"pool"
)

// ArchivoDAO maneja las operaciones de base de datos para archivos
type ArchivoDAO struct {
	dbPool *pool.DBConnectionPool
}

// NuevoArchivoDAO crea una nueva instancia de ArchivoDAO
func NuevoArchivoDAO(dbPool *pool.DBConnectionPool) *ArchivoDAO {
	return &ArchivoDAO{
		dbPool: dbPool,
	}
}

// Crear inserta un nuevo archivo en la base de datos
func (dao *ArchivoDAO) Crear(archivo *model.ArchivoMetadata) error {
	query := `INSERT INTO archivos (id, nombre_original, tamano_bytes, ruta, subido_por, fecha_subida) 
              VALUES (?, ?, ?, ?, ?, ?)`

	_, err := dao.dbPool.DB().Exec(
		query,
		archivo.ID().String(),
		archivo.NombreOriginal(),
		archivo.TamanoBytes(),
		archivo.Ruta(),
		archivo.SubidoPor().String(),
		archivo.FechaSubida(),
	)

	return err
}

// GetByID busca un archivo por su ID
func (dao *ArchivoDAO) GetByID(id uuid.UUID) (*model.ArchivoMetadata, error) {
	query := `SELECT id, nombre_original, tamano_bytes, ruta, subido_por, fecha_subida 
              FROM archivos 
              WHERE id = ?`

	var idStr, subidoPorStr string
	var nombreOriginal, ruta string
	var tamanoBytes int64
	var fechaSubida time.Time

	err := dao.dbPool.DB().QueryRow(query, id.String()).Scan(
		&idStr,
		&nombreOriginal,
		&tamanoBytes,
		&ruta,
		&subidoPorStr,
		&fechaSubida,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	archivoID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	subidoPor, err := uuid.Parse(subidoPorStr)
	if err != nil {
		return nil, err
	}

	return model.NewArchivoMetadata(
		archivoID,
		nombreOriginal,
		tamanoBytes,
		ruta,
		subidoPor,
		fechaSubida,
	)
}

// Update actualiza un archivo existente en la base de datos
func (dao *ArchivoDAO) Update(archivo *model.ArchivoMetadata) error {
	query := `UPDATE archivos 
              SET nombre_original = ?, tamano_bytes = ?, ruta = ?, subido_por = ?, fecha_subida = ? 
              WHERE id = ?`

	_, err := dao.dbPool.DB().Exec(
		query,
		archivo.NombreOriginal(),
		archivo.TamanoBytes(),
		archivo.Ruta(),
		archivo.SubidoPor().String(),
		archivo.FechaSubida(),
		archivo.ID().String(),
	)

	return err
}

// Delete elimina un archivo por su ID
func (dao *ArchivoDAO) Delete(id uuid.UUID) error {
	query := `DELETE FROM archivos WHERE id = ?`

	_, err := dao.dbPool.DB().Exec(query, id.String())
	return err
}

// GetByUploader obtiene archivos subidos por un usuario específico
func (dao *ArchivoDAO) GetByUploader(uploaderID uuid.UUID) ([]*model.ArchivoMetadata, error) {
	query := `SELECT id, nombre_original, tamano_bytes, ruta, subido_por, fecha_subida 
              FROM archivos 
              WHERE subido_por = ?`

	rows, err := dao.dbPool.DB().Query(query, uploaderID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var archivos []*model.ArchivoMetadata

	for rows.Next() {
		var idStr, subidoPorStr string
		var nombreOriginal, ruta string
		var tamanoBytes int64
		var fechaSubida time.Time

		err := rows.Scan(
			&idStr,
			&nombreOriginal,
			&tamanoBytes,
			&ruta,
			&subidoPorStr,
			&fechaSubida,
		)

		if err != nil {
			return nil, err
		}

		archivoID, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		subidoPor, err := uuid.Parse(subidoPorStr)
		if err != nil {
			return nil, err
		}

		archivo, err := model.NewArchivoMetadata(
			archivoID,
			nombreOriginal,
			tamanoBytes,
			ruta,
			subidoPor,
			fechaSubida,
		)

		if err != nil {
			return nil, err
		}

		archivos = append(archivos, archivo)
	}

	return archivos, rows.Err()
}
