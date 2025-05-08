package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"model"
	"pool"
)

// MensajeDAO maneja las operaciones de base de datos para entidades MensajeServidor
type MensajeDAO struct {
	dbPool *pool.DBConnectionPool
}

// NuevoMensajeDAO crea una nueva instancia de MensajeDAO
func NuevoMensajeDAO(dbPool *pool.DBConnectionPool) *MensajeDAO {
	return &MensajeDAO{dbPool: dbPool}
}

// Crear persiste un nuevo mensaje en la base de datos
func (dao *MensajeDAO) Crear(mensaje *model.MensajeServidor) error {
	query := `INSERT INTO mensajes (id, remitente_id, destino_usuario_id, canal_id, 
              chat_privado_id, contenido, timestamp, archivo_id)
              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := dao.dbPool.DB().Exec(
		query,
		mensaje.ID().String(),
		mensaje.RemitenteID().String(),
		nullableUUID(mensaje.DestinoUsuarioID()),
		nullableUUID(mensaje.CanalID()),
		nullableUUID(mensaje.ChatPrivadoID()),
		mensaje.Contenido(),
		mensaje.Timestamp(),
		nullableUUID(mensaje.ArchivoID()),
	)

	return err
}

// BuscarPorID recupera un mensaje por su ID
func (dao *MensajeDAO) BuscarPorID(id uuid.UUID) (*model.MensajeServidor, error) {
	query := `SELECT id, remitente_id, destino_usuario_id, canal_id, 
              chat_privado_id, contenido, timestamp, archivo_id 
              FROM mensajes WHERE id = ?`

	row := dao.dbPool.DB().QueryRow(query, id.String())
	return dao.escanearMensaje(row)
}

// BuscarPorCanalID recupera mensajes de un canal específico
func (dao *MensajeDAO) BuscarPorCanalID(canalID uuid.UUID, limite int) ([]*model.MensajeServidor, error) {
	query := `SELECT id, remitente_id, destino_usuario_id, canal_id, 
              chat_privado_id, contenido, timestamp, archivo_id 
              FROM mensajes WHERE canal_id = ? 
              ORDER BY timestamp DESC LIMIT ?`

	rows, err := dao.dbPool.DB().Query(query, canalID.String(), limite)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearFilasMensajes(rows)
}

// BuscarPorChatPrivadoID recupera mensajes de un chat privado
func (dao *MensajeDAO) BuscarPorChatPrivadoID(chatPrivadoID uuid.UUID, limite int) ([]*model.MensajeServidor, error) {
	query := `SELECT id, remitente_id, destino_usuario_id, canal_id, 
              chat_privado_id, contenido, timestamp, archivo_id 
              FROM mensajes WHERE chat_privado_id = ? 
              ORDER BY timestamp DESC LIMIT ?`

	rows, err := dao.dbPool.DB().Query(query, chatPrivadoID.String(), limite)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearFilasMensajes(rows)
}

// BuscarMensajesDirectos recupera mensajes directos entre dos usuarios
func (dao *MensajeDAO) BuscarMensajesDirectos(remitenteID, destinatarioID uuid.UUID, limite int) ([]*model.MensajeServidor, error) {
	query := `SELECT id, remitente_id, destino_usuario_id, canal_id, 
              chat_privado_id, contenido, timestamp, archivo_id 
              FROM mensajes 
              WHERE (remitente_id = ? AND destino_usuario_id = ?) 
              OR (remitente_id = ? AND destino_usuario_id = ?) 
              ORDER BY timestamp DESC LIMIT ?`

	rows, err := dao.dbPool.DB().Query(query,
		remitenteID.String(), destinatarioID.String(),
		destinatarioID.String(), remitenteID.String(),
		limite)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearFilasMensajes(rows)
}

// Eliminar elimina un mensaje de la base de datos
func (dao *MensajeDAO) Eliminar(id uuid.UUID) error {
	query := `DELETE FROM mensajes WHERE id = ?`
	_, err := dao.dbPool.DB().Exec(query, id.String())
	return err
}

// escanearMensaje es un método auxiliar para escanear una fila en un mensaje
func (dao *MensajeDAO) escanearMensaje(row *sql.Row) (*model.MensajeServidor, error) {
	var (
		idStr, remitenteIDStr                                           string
		destinoUsuarioIDStr, canalIDStr, chatPrivadoIDStr, archivoIDStr sql.NullString
		contenido                                                       string
		timestamp                                                       time.Time
	)

	if err := row.Scan(
		&idStr,
		&remitenteIDStr,
		&destinoUsuarioIDStr,
		&canalIDStr,
		&chatPrivadoIDStr,
		&contenido,
		&timestamp,
		&archivoIDStr,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Mensaje no encontrado
		}
		return nil, err
	}

	return dao.construirMensajeDesdeValoresEscaneados(
		idStr, remitenteIDStr, destinoUsuarioIDStr,
		canalIDStr, chatPrivadoIDStr, contenido,
		timestamp, archivoIDStr)
}

// escanearFilaMensaje es un método auxiliar para escanear una fila de mensajes
func (dao *MensajeDAO) escanearFilaMensaje(rows *sql.Rows) (*model.MensajeServidor, error) {
	var (
		idStr, remitenteIDStr                                           string
		destinoUsuarioIDStr, canalIDStr, chatPrivadoIDStr, archivoIDStr sql.NullString
		contenido                                                       string
		timestamp                                                       time.Time
	)

	if err := rows.Scan(
		&idStr,
		&remitenteIDStr,
		&destinoUsuarioIDStr,
		&canalIDStr,
		&chatPrivadoIDStr,
		&contenido,
		&timestamp,
		&archivoIDStr,
	); err != nil {
		return nil, err
	}

	return dao.construirMensajeDesdeValoresEscaneados(
		idStr, remitenteIDStr, destinoUsuarioIDStr,
		canalIDStr, chatPrivadoIDStr, contenido,
		timestamp, archivoIDStr)
}

// escanearFilasMensajes es un método auxiliar para escanear múltiples filas
func (dao *MensajeDAO) escanearFilasMensajes(rows *sql.Rows) ([]*model.MensajeServidor, error) {
	var mensajes []*model.MensajeServidor
	for rows.Next() {
		mensaje, err := dao.escanearFilaMensaje(rows)
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

// construirMensajeDesdeValoresEscaneados es un método auxiliar para convertir valores escaneados en un objeto MensajeServidor
func (dao *MensajeDAO) construirMensajeDesdeValoresEscaneados(
	idStr, remitenteIDStr string,
	destinoUsuarioIDStr, canalIDStr, chatPrivadoIDStr sql.NullString,
	contenido string,
	timestamp time.Time,
	archivoIDStr sql.NullString,
) (*model.MensajeServidor, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	remitenteID, err := uuid.Parse(remitenteIDStr)
	if err != nil {
		return nil, err
	}

	var destinoUsuarioID, canalID, chatPrivadoID, archivoID uuid.UUID

	if destinoUsuarioIDStr.Valid {
		destinoUsuarioID, err = uuid.Parse(destinoUsuarioIDStr.String)
		if err != nil {
			return nil, err
		}
		return model.NewMensajeDirecto(id, remitenteID, destinoUsuarioID, contenido, timestamp, archivoID)
	} else if canalIDStr.Valid {
		canalID, err = uuid.Parse(canalIDStr.String)
		if err != nil {
			return nil, err
		}
		return model.NewMensajeCanal(id, remitenteID, canalID, contenido, timestamp, archivoID)
	} else if chatPrivadoIDStr.Valid {
		chatPrivadoID, err = uuid.Parse(chatPrivadoIDStr.String)
		if err != nil {
			return nil, err
		}
		return model.NewMensajeChatPrivado(id, remitenteID, chatPrivadoID, contenido, timestamp, archivoID)
	}

	// No debería llegar a este punto si los datos son válidos
	return nil, sql.ErrNoRows
}

// Helper para manejar UUID que pueden ser nulas en la base de datos
func nullableUUID(id uuid.UUID) interface{} {
	if id == uuid.Nil {
		return nil
	}
	return id.String()
}
