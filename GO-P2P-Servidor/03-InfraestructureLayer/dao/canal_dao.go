package dao

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"model"
	"pool"
)

// CanalDAO maneja las operaciones de base de datos para entidades CanalServidor
type CanalDAO struct {
	dbPool *pool.DBConnectionPool
}

// NuevoCanalDAO crea una nueva instancia de CanalDAO
func NuevoCanalDAO(dbPool *pool.DBConnectionPool) *CanalDAO {
	return &CanalDAO{dbPool: dbPool}
}

// Crear guarda un nuevo canal en la base de datos
func (dao *CanalDAO) Crear(canal *model.CanalServidor) error {
	query := `INSERT INTO canales (id, nombre, descripcion, tipo)
              VALUES (?, ?, ?, ?)`

	_, err := dao.dbPool.DB().Exec(
		query,
		canal.ID().String(),
		canal.Nombre(),
		canal.Descripcion(),
		string(canal.Tipo()), // Convertimos CanalTipo a string
	)

	return err
}

// BuscarPorID recupera un canal por su ID
func (dao *CanalDAO) BuscarPorID(id uuid.UUID) (*model.CanalServidor, error) {
	query := `SELECT id, nombre, descripcion, tipo
              FROM canales WHERE id = ?`

	row := dao.dbPool.DB().QueryRow(query, id.String())
	return dao.escanearCanal(row)
}

// Actualizar actualiza un canal existente en la base de datos
func (dao *CanalDAO) Actualizar(canal *model.CanalServidor) error {
	query := `UPDATE canales 
              SET nombre = ?, descripcion = ?, tipo = ?
              WHERE id = ?`

	_, err := dao.dbPool.DB().Exec(
		query,
		canal.Nombre(),
		canal.Descripcion(),
		string(canal.Tipo()), // Convertimos CanalTipo a string
		canal.ID().String(),
	)

	return err
}

// Eliminar elimina un canal de la base de datos
func (dao *CanalDAO) Eliminar(id uuid.UUID) error {
	query := `DELETE FROM canales WHERE id = ?`
	_, err := dao.dbPool.DB().Exec(query, id.String())
	return err
}

// escanearCanal escanea una fila de la consulta en un objeto CanalServidor
func (dao *CanalDAO) escanearCanal(row *sql.Row) (*model.CanalServidor, error) {
	var (
		idStr, nombre, descripcion, tipoStr string
	)

	if err := row.Scan(&idStr, &nombre, &descripcion, &tipoStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Canal no encontrado
		}
		return nil, err
	}

	// Convertir string ID a UUID
	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	// Convertir string a CanalTipo
	var tipo model.CanalTipo
	switch tipoStr {
	case "PUBLICO":
		tipo = model.CanalPublico
	case "PRIVADO":
		tipo = model.CanalPrivado
	default:
		return nil, fmt.Errorf("tipo de canal desconocido: %s", tipoStr)
	}

	return model.NewCanalServidor(parsedID, nombre, descripcion, tipo)
}

// BuscarTodos recupera todos los canales de la base de datos
func (dao *CanalDAO) BuscarTodos() ([]*model.CanalServidor, error) {
	query := `SELECT id, nombre, descripcion, tipo 
              FROM canales`

	rows, err := dao.dbPool.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var canales []*model.CanalServidor
	for rows.Next() {
		canal, err := dao.escanearFilaCanal(rows)
		if err != nil {
			return nil, err
		}
		canales = append(canales, canal)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return canales, nil
}

// escanearFilaCanal escanea una fila de sql.Rows en un objeto CanalServidor
func (dao *CanalDAO) escanearFilaCanal(rows *sql.Rows) (*model.CanalServidor, error) {
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

	// Convertir string a CanalTipo
	var tipo model.CanalTipo
	switch tipoStr {
	case "PUBLICO":
		tipo = model.CanalPublico
	case "PRIVADO":
		tipo = model.CanalPrivado
	default:
		return nil, fmt.Errorf("tipo de canal desconocido: %s", tipoStr)
	}

	return model.NewCanalServidor(parsedID, nombre, descripcion, tipo)
}
