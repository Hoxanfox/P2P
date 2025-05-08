package dao

import (
	"database/sql"
	"pool"

	"github.com/google/uuid"
	"model"
)

// CanalMiembroDAO maneja las operaciones de base de datos para entidades CanalMiembro
type CanalMiembroDAO struct {
	dbPool *pool.DBConnectionPool
}

// NuevoCanalMiembroDAO crea una nueva instancia de CanalMiembroDAO
func NuevoCanalMiembroDAO(dbPool *pool.DBConnectionPool) *CanalMiembroDAO {
	return &CanalMiembroDAO{dbPool: dbPool}
}

// Guardar persiste un miembro de canal en la base de datos
func (dao *CanalMiembroDAO) Guardar(miembro *model.CanalMiembro) error {
	query := `INSERT INTO canal_miembros (canal_id, usuario_id, rol)
              VALUES (?, ?, ?)`

	_, err := dao.dbPool.DB().Exec(
		query,
		miembro.CanalID().String(),
		miembro.UsuarioID().String(),
		miembro.Rol(),
	)

	return err
}

// BuscarPorIDs recupera un miembro de canal por su canal_id y usuario_id
func (dao *CanalMiembroDAO) BuscarPorIDs(canalID, usuarioID uuid.UUID) (*model.CanalMiembro, error) {
	query := `SELECT canal_id, usuario_id, rol
              FROM canal_miembros 
              WHERE canal_id = ? AND usuario_id = ?`

	row := dao.dbPool.DB().QueryRow(query, canalID.String(), usuarioID.String())
	return dao.escanearCanalMiembro(row)
}

// BuscarPorCanalID recupera todos los miembros de un canal específico
func (dao *CanalMiembroDAO) BuscarPorCanalID(canalID uuid.UUID) ([]*model.CanalMiembro, error) {
	query := `SELECT canal_id, usuario_id, rol
              FROM canal_miembros 
              WHERE canal_id = ?`

	rows, err := dao.dbPool.DB().Query(query, canalID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesCanalMiembros(rows)
}

// BuscarPorUsuarioID recupera todos los canales a los que pertenece un usuario
func (dao *CanalMiembroDAO) BuscarPorUsuarioID(usuarioID uuid.UUID) ([]*model.CanalMiembro, error) {
	query := `SELECT canal_id, usuario_id, rol
              FROM canal_miembros 
              WHERE usuario_id = ?`

	rows, err := dao.dbPool.DB().Query(query, usuarioID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesCanalMiembros(rows)
}

// Actualizar actualiza el rol de un miembro de canal
func (dao *CanalMiembroDAO) Actualizar(miembro *model.CanalMiembro) error {
	query := `UPDATE canal_miembros 
              SET rol = ?
              WHERE canal_id = ? AND usuario_id = ?`

	_, err := dao.dbPool.DB().Exec(
		query,
		miembro.Rol(),
		miembro.CanalID().String(),
		miembro.UsuarioID().String(),
	)

	return err
}

// Eliminar elimina un miembro de canal
func (dao *CanalMiembroDAO) Eliminar(canalID, usuarioID uuid.UUID) error {
	query := `DELETE FROM canal_miembros 
              WHERE canal_id = ? AND usuario_id = ?`

	_, err := dao.dbPool.DB().Exec(query, canalID.String(), usuarioID.String())
	return err
}

// escanearCanalMiembro escanea una fila en un objeto CanalMiembro
func (dao *CanalMiembroDAO) escanearCanalMiembro(row *sql.Row) (*model.CanalMiembro, error) {
	var (
		canalIDStr   string
		usuarioIDStr string
		rol          string
	)

	if err := row.Scan(&canalIDStr, &usuarioIDStr, &rol); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Miembro de canal no encontrado
		}
		return nil, err
	}

	canalID, err := uuid.Parse(canalIDStr)
	if err != nil {
		return nil, err
	}

	usuarioID, err := uuid.Parse(usuarioIDStr)
	if err != nil {
		return nil, err
	}

	return model.NewCanalMiembro(canalID, usuarioID, rol)
}

// escanearMultiplesCanalMiembros escanea múltiples filas en objetos CanalMiembro
func (dao *CanalMiembroDAO) escanearMultiplesCanalMiembros(rows *sql.Rows) ([]*model.CanalMiembro, error) {
	var miembros []*model.CanalMiembro

	for rows.Next() {
		var (
			canalIDStr   string
			usuarioIDStr string
			rol          string
		)

		if err := rows.Scan(&canalIDStr, &usuarioIDStr, &rol); err != nil {
			return nil, err
		}

		canalID, err := uuid.Parse(canalIDStr)
		if err != nil {
			return nil, err
		}

		usuarioID, err := uuid.Parse(usuarioIDStr)
		if err != nil {
			return nil, err
		}

		miembro, err := model.NewCanalMiembro(canalID, usuarioID, rol)
		if err != nil {
			return nil, err
		}

		miembros = append(miembros, miembro)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return miembros, nil
}
