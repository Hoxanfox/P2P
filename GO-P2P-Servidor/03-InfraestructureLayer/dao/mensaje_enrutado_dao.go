package dao

import (
	"database/sql"
	"github.com/google/uuid"
	"model"
	"pool"
	"time"
)

// MensajeEnrutadoDAO maneja las operaciones de base de datos para entidades RoutedMessage (MensajeEnrutado)
type MensajeEnrutadoDAO struct {
	dbPool *pool.DBConnectionPool
}

// NuevoMensajeEnrutadoDAO crea una nueva instancia de MensajeEnrutadoDAO
func NuevoMensajeEnrutadoDAO(dbPool *pool.DBConnectionPool) *MensajeEnrutadoDAO {
	return &MensajeEnrutadoDAO{dbPool: dbPool}
}

// Crear persiste un nuevo mensaje enrutado en la base de datos
func (dao *MensajeEnrutadoDAO) Crear(mensaje *model.RoutedMessage) error {
	query := `INSERT INTO routed_messages (mensaje_id, nodo_destino_id, enruta_at)
              VALUES (?, ?, ?)`

	_, err := dao.dbPool.DB().Exec(
		query,
		mensaje.MensajeID().String(),
		mensaje.NodoDestinoID().String(),
		mensaje.EnrutaAt(),
	)

	return err
}

// BuscarPorMensajeID recupera un mensaje enrutado por su ID de mensaje
func (dao *MensajeEnrutadoDAO) BuscarPorMensajeID(mensajeID uuid.UUID) (*model.RoutedMessage, error) {
	query := `SELECT mensaje_id, nodo_destino_id, enruta_at
              FROM routed_messages WHERE mensaje_id = ?`

	row := dao.dbPool.DB().QueryRow(query, mensajeID.String())
	return dao.escanearMensajeEnrutado(row)
}

// BuscarTodos recupera todos los mensajes enrutados de la base de datos
func (dao *MensajeEnrutadoDAO) BuscarTodos() ([]*model.RoutedMessage, error) {
	query := `SELECT mensaje_id, nodo_destino_id, enruta_at
              FROM routed_messages`

	rows, err := dao.dbPool.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesMensajesEnrutados(rows)
}

// BuscarPorNodoDestinoID recupera todos los mensajes enrutados para un nodo destino específico
func (dao *MensajeEnrutadoDAO) BuscarPorNodoDestinoID(nodoDestinoID uuid.UUID) ([]*model.RoutedMessage, error) {
	query := `SELECT mensaje_id, nodo_destino_id, enruta_at
              FROM routed_messages WHERE nodo_destino_id = ?`

	rows, err := dao.dbPool.DB().Query(query, nodoDestinoID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesMensajesEnrutados(rows)
}

// BuscarPorRangoTiempo recupera mensajes enrutados dentro de un rango de tiempo específico
func (dao *MensajeEnrutadoDAO) BuscarPorRangoTiempo(inicio, fin time.Time) ([]*model.RoutedMessage, error) {
	query := `SELECT mensaje_id, nodo_destino_id, enruta_at
              FROM routed_messages
              WHERE enruta_at BETWEEN ? AND ?`

	rows, err := dao.dbPool.DB().Query(query, inicio, fin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesMensajesEnrutados(rows)
}

// Eliminar elimina un mensaje enrutado de la base de datos
func (dao *MensajeEnrutadoDAO) Eliminar(mensajeID uuid.UUID) error {
	query := `DELETE FROM routed_messages WHERE mensaje_id = ?`
	_, err := dao.dbPool.DB().Exec(query, mensajeID.String())
	return err
}

// escanearMensajeEnrutado escanea una fila en un objeto RoutedMessage (MensajeEnrutado)
func (dao *MensajeEnrutadoDAO) escanearMensajeEnrutado(row *sql.Row) (*model.RoutedMessage, error) {
	var (
		mensajeIDStr, nodoDestinoIDStr string
		enrutaAt                       time.Time
	)

	if err := row.Scan(&mensajeIDStr, &nodoDestinoIDStr, &enrutaAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Mensaje enrutado no encontrado
		}
		return nil, err
	}

	mensajeID, err := uuid.Parse(mensajeIDStr)
	if err != nil {
		return nil, err
	}

	nodoDestinoID, err := uuid.Parse(nodoDestinoIDStr)
	if err != nil {
		return nil, err
	}

	return model.NewRoutedMessage(mensajeID, nodoDestinoID, enrutaAt)
}

// escanearMultiplesMensajesEnrutados escanea múltiples filas en objetos RoutedMessage (MensajeEnrutado)
func (dao *MensajeEnrutadoDAO) escanearMultiplesMensajesEnrutados(rows *sql.Rows) ([]*model.RoutedMessage, error) {
	var mensajes []*model.RoutedMessage

	for rows.Next() {
		var (
			mensajeIDStr, nodoDestinoIDStr string
			enrutaAt                       time.Time
		)

		if err := rows.Scan(&mensajeIDStr, &nodoDestinoIDStr, &enrutaAt); err != nil {
			return nil, err
		}

		mensajeID, err := uuid.Parse(mensajeIDStr)
		if err != nil {
			return nil, err
		}

		nodoDestinoID, err := uuid.Parse(nodoDestinoIDStr)
		if err != nil {
			return nil, err
		}

		mensaje, err := model.NewRoutedMessage(mensajeID, nodoDestinoID, enrutaAt)
		if err != nil {
			return nil, err
		}

		mensajes = append(mensajes, mensaje)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return mensajes, nil
}
