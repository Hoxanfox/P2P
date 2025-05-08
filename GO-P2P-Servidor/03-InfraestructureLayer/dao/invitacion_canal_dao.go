package dao

import (
	"database/sql"
	"pool"
	"time"

	"github.com/google/uuid"
	"model"
)

// InvitacionCanalDAO maneja las operaciones de base de datos para entidades InvitacionCanal
type InvitacionCanalDAO struct {
	dbPool *pool.DBConnectionPool
}

// NuevoInvitacionCanalDAO crea una nueva instancia de InvitacionCanalDAO
func NuevoInvitacionCanalDAO(dbPool *pool.DBConnectionPool) *InvitacionCanalDAO {
	return &InvitacionCanalDAO{dbPool: dbPool}
}

// Guardar persiste una invitación de canal en la base de datos
func (dao *InvitacionCanalDAO) Guardar(invitacion *model.InvitacionCanal) error {
	query := `INSERT INTO invitaciones_canal (id, canal_id, destinatario_id, estado, fecha_envio)
              VALUES (?, ?, ?, ?, ?)`

	_, err := dao.dbPool.DB().Exec(
		query,
		invitacion.ID().String(),
		invitacion.CanalID().String(),
		invitacion.DestinatarioID().String(),
		string(invitacion.Estado()), // EstadoInvitacion es de tipo string
		invitacion.FechaEnvio(),
	)

	return err
}

// BuscarPorID recupera una invitación de canal por su ID
func (dao *InvitacionCanalDAO) BuscarPorID(id uuid.UUID) (*model.InvitacionCanal, error) {
	query := `SELECT id, canal_id, destinatario_id, estado, fecha_envio
              FROM invitaciones_canal WHERE id = ?`

	row := dao.dbPool.DB().QueryRow(query, id.String())
	return dao.escanearInvitacionCanal(row)
}

// BuscarPorCanalID recupera todas las invitaciones para un canal específico
func (dao *InvitacionCanalDAO) BuscarPorCanalID(canalID uuid.UUID) ([]*model.InvitacionCanal, error) {
	query := `SELECT id, canal_id, destinatario_id, estado, fecha_envio
              FROM invitaciones_canal WHERE canal_id = ?`

	rows, err := dao.dbPool.DB().Query(query, canalID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesInvitacionesCanal(rows)
}

// BuscarPorDestinatarioID recupera todas las invitaciones para un destinatario específico
func (dao *InvitacionCanalDAO) BuscarPorDestinatarioID(destinatarioID uuid.UUID) ([]*model.InvitacionCanal, error) {
	query := `SELECT id, canal_id, destinatario_id, estado, fecha_envio
              FROM invitaciones_canal WHERE destinatario_id = ?`

	rows, err := dao.dbPool.DB().Query(query, destinatarioID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesInvitacionesCanal(rows)
}

// BuscarPorEstado recupera todas las invitaciones con un estado específico
func (dao *InvitacionCanalDAO) BuscarPorEstado(estado model.EstadoInvitacion) ([]*model.InvitacionCanal, error) {
	query := `SELECT id, canal_id, destinatario_id, estado, fecha_envio
              FROM invitaciones_canal WHERE estado = ?`

	rows, err := dao.dbPool.DB().Query(query, string(estado))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesInvitacionesCanal(rows)
}

// Actualizar actualiza el estado de una invitación de canal
func (dao *InvitacionCanalDAO) Actualizar(invitacion *model.InvitacionCanal) error {
	query := `UPDATE invitaciones_canal 
              SET estado = ? 
              WHERE id = ?`

	_, err := dao.dbPool.DB().Exec(
		query,
		string(invitacion.Estado()),
		invitacion.ID().String(),
	)

	return err
}

// Eliminar elimina una invitación de canal
func (dao *InvitacionCanalDAO) Eliminar(id uuid.UUID) error {
	query := `DELETE FROM invitaciones_canal WHERE id = ?`
	_, err := dao.dbPool.DB().Exec(query, id.String())
	return err
}

// escanearInvitacionCanal escanea una fila en un objeto InvitacionCanal
func (dao *InvitacionCanalDAO) escanearInvitacionCanal(row *sql.Row) (*model.InvitacionCanal, error) {
	var (
		idStr            string
		canalIDStr       string
		destinatarioIDStr string
		estadoStr        string
		fechaEnvio       time.Time
	)

	if err := row.Scan(&idStr, &canalIDStr, &destinatarioIDStr, &estadoStr, &fechaEnvio); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Invitación no encontrada
		}
		return nil, err
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	canalID, err := uuid.Parse(canalIDStr)
	if err != nil {
		return nil, err
	}

	destinatarioID, err := uuid.Parse(destinatarioIDStr)
	if err != nil {
		return nil, err
	}

	// Convertir el string a EstadoInvitacion
	estado := model.EstadoInvitacion(estadoStr)

	return model.NewInvitacionCanal(id, canalID, destinatarioID, estado, fechaEnvio)
}

// escanearMultiplesInvitacionesCanal escanea múltiples filas en objetos InvitacionCanal
func (dao *InvitacionCanalDAO) escanearMultiplesInvitacionesCanal(rows *sql.Rows) ([]*model.InvitacionCanal, error) {
	var invitaciones []*model.InvitacionCanal

	for rows.Next() {
		var (
			idStr            string
			canalIDStr       string
			destinatarioIDStr string
			estadoStr        string
			fechaEnvio       time.Time
		)

		if err := rows.Scan(&idStr, &canalIDStr, &destinatarioIDStr, &estadoStr, &fechaEnvio); err != nil {
			return nil, err
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		canalID, err := uuid.Parse(canalIDStr)
		if err != nil {
			return nil, err
		}

		destinatarioID, err := uuid.Parse(destinatarioIDStr)
		if err != nil {
			return nil, err
		}

		// Convertir el string a EstadoInvitacion
		estado := model.EstadoInvitacion(estadoStr)

		invitacion, err := model.NewInvitacionCanal(id, canalID, destinatarioID, estado, fechaEnvio)
		if err != nil {
			return nil, err
		}

		invitaciones = append(invitaciones, invitacion)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return invitaciones, nil
}
