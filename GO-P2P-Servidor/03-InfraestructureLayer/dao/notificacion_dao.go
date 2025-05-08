package dao

import (
	"database/sql"
	"pool"
	"time"

	"github.com/google/uuid"
	"model"
)

// NotificacionDAO maneja las operaciones de base de datos para entidades Notificacion
type NotificacionDAO struct {
	dbPool *pool.DBConnectionPool
}

// NuevoNotificacionDAO crea una nueva instancia de NotificacionDAO
func NuevoNotificacionDAO(dbPool *pool.DBConnectionPool) *NotificacionDAO {
	return &NotificacionDAO{dbPool: dbPool}
}

// Guardar persiste una notificación en la base de datos
func (dao *NotificacionDAO) Guardar(notificacion *model.Notificacion) error {
	query := `INSERT INTO notificaciones (id, usuario_id, contenido, fecha, leido, invitacion_id)
              VALUES (?, ?, ?, ?, ?, ?)`

	// Si invitacionID es uuid.Nil, se guardará como NULL en la base de datos
	var invitacionIDStr *string
	if notificacion.InvitacionID() != uuid.Nil {
		id := notificacion.InvitacionID().String()
		invitacionIDStr = &id
	}

	_, err := dao.dbPool.DB().Exec(
		query,
		notificacion.ID().String(),
		notificacion.UsuarioID().String(),
		notificacion.Contenido(),
		notificacion.Fecha(),
		notificacion.Leido(),
		invitacionIDStr,
	)

	return err
}

// BuscarPorID recupera una notificación por su ID
func (dao *NotificacionDAO) BuscarPorID(id uuid.UUID) (*model.Notificacion, error) {
	query := `SELECT id, usuario_id, contenido, fecha, leido, invitacion_id
              FROM notificaciones WHERE id = ?`

	row := dao.dbPool.DB().QueryRow(query, id.String())
	return dao.escanearNotificacion(row)
}

// BuscarPorUsuarioID recupera todas las notificaciones para un usuario específico
func (dao *NotificacionDAO) BuscarPorUsuarioID(usuarioID uuid.UUID) ([]*model.Notificacion, error) {
	query := `SELECT id, usuario_id, contenido, fecha, leido, invitacion_id
              FROM notificaciones WHERE usuario_id = ? ORDER BY fecha DESC`

	rows, err := dao.dbPool.DB().Query(query, usuarioID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesNotificaciones(rows)
}

// BuscarNoLeidas recupera todas las notificaciones no leídas para un usuario
func (dao *NotificacionDAO) BuscarNoLeidas(usuarioID uuid.UUID) ([]*model.Notificacion, error) {
	query := `SELECT id, usuario_id, contenido, fecha, leido, invitacion_id
              FROM notificaciones WHERE usuario_id = ? AND leido = false ORDER BY fecha DESC`

	rows, err := dao.dbPool.DB().Query(query, usuarioID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesNotificaciones(rows)
}

// BuscarPorInvitacionID recupera todas las notificaciones relacionadas con una invitación
func (dao *NotificacionDAO) BuscarPorInvitacionID(invitacionID uuid.UUID) ([]*model.Notificacion, error) {
	query := `SELECT id, usuario_id, contenido, fecha, leido, invitacion_id
              FROM notificaciones WHERE invitacion_id = ? ORDER BY fecha DESC`

	rows, err := dao.dbPool.DB().Query(query, invitacionID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesNotificaciones(rows)
}

// ActualizarEstadoLeido actualiza el estado de leído de una notificación
func (dao *NotificacionDAO) ActualizarEstadoLeido(id uuid.UUID, leido bool) error {
	query := `UPDATE notificaciones SET leido = ? WHERE id = ?`
	_, err := dao.dbPool.DB().Exec(query, leido, id.String())
	return err
}

// MarcarTodasComoLeidas marca todas las notificaciones de un usuario como leídas
func (dao *NotificacionDAO) MarcarTodasComoLeidas(usuarioID uuid.UUID) error {
	query := `UPDATE notificaciones SET leido = true WHERE usuario_id = ? AND leido = false`
	_, err := dao.dbPool.DB().Exec(query, usuarioID.String())
	return err
}

// Eliminar elimina una notificación
func (dao *NotificacionDAO) Eliminar(id uuid.UUID) error {
	query := `DELETE FROM notificaciones WHERE id = ?`
	_, err := dao.dbPool.DB().Exec(query, id.String())
	return err
}

// escanearNotificacion escanea una fila en un objeto Notificacion
func (dao *NotificacionDAO) escanearNotificacion(row *sql.Row) (*model.Notificacion, error) {
	var (
		idStr           string
		usuarioIDStr    string
		contenido       string
		fecha           time.Time
		leido           bool
		invitacionIDStr sql.NullString
	)

	if err := row.Scan(&idStr, &usuarioIDStr, &contenido, &fecha, &leido, &invitacionIDStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Notificación no encontrada
		}
		return nil, err
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	usuarioID, err := uuid.Parse(usuarioIDStr)
	if err != nil {
		return nil, err
	}

	// Procesar invitacionID (puede ser NULL)
	var invitacionID uuid.UUID
	if invitacionIDStr.Valid {
		invitacionID, err = uuid.Parse(invitacionIDStr.String)
		if err != nil {
			return nil, err
		}
	} else {
		invitacionID = uuid.Nil
	}

	notificacion, err := model.NewNotificacion(id, usuarioID, contenido, fecha, invitacionID)
	if err != nil {
		return nil, err
	}

	// Establecer el estado de leído según lo guardado en la BD
	if leido {
		notificacion.MarcarComoLeida()
	} else {
		notificacion.MarcarComoNoLeida()
	}

	return notificacion, nil
}

// escanearMultiplesNotificaciones escanea múltiples filas en objetos Notificacion
func (dao *NotificacionDAO) escanearMultiplesNotificaciones(rows *sql.Rows) ([]*model.Notificacion, error) {
	var notificaciones []*model.Notificacion

	for rows.Next() {
		var (
			idStr           string
			usuarioIDStr    string
			contenido       string
			fecha           time.Time
			leido           bool
			invitacionIDStr sql.NullString
		)

		if err := rows.Scan(&idStr, &usuarioIDStr, &contenido, &fecha, &leido, &invitacionIDStr); err != nil {
			return nil, err
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		usuarioID, err := uuid.Parse(usuarioIDStr)
		if err != nil {
			return nil, err
		}

		// Procesar invitacionID (puede ser NULL)
		var invitacionID uuid.UUID
		if invitacionIDStr.Valid {
			invitacionID, err = uuid.Parse(invitacionIDStr.String)
			if err != nil {
				return nil, err
			}
		} else {
			invitacionID = uuid.Nil
		}

		notificacion, err := model.NewNotificacion(id, usuarioID, contenido, fecha, invitacionID)
		if err != nil {
			return nil, err
		}

		// Establecer el estado de leído según lo guardado en la BD
		if leido {
			notificacion.MarcarComoLeida()
		} else {
			notificacion.MarcarComoNoLeida()
		}

		notificaciones = append(notificaciones, notificacion)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notificaciones, nil
}
