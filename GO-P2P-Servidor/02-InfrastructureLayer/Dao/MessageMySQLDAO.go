package Dao

import (
	"database/sql"
	"time"

	"github.com/P2P/GO-P2P-SERVIDOR/04-DomainLayer/model"
	"github.com/google/uuid"
)

// MessageMySQLDAO maneja las operaciones de base de datos para entidades MensajeServidor
type MessageMySQLDAO struct {
	db *sql.DB
}

// NewMessageMySQLDAO crea una nueva instancia de MessageMySQLDAO
func NewMessageMySQLDAO(db *sql.DB) *MessageMySQLDAO {
	return &MessageMySQLDAO{db: db}
}

// Create persiste un nuevo mensaje en la base de datos
func (dao *MessageMySQLDAO) Create(mensaje *model.MensajeServidor) error {
	query := `INSERT INTO mensajes (id, remitente_id, destino_usuario_id, canal_id, 
              chat_privado_id, contenido, timestamp, archivo_id)
              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := dao.db.Exec(
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

// FindByID recupera un mensaje por su ID
func (dao *MessageMySQLDAO) FindByID(id uuid.UUID) (*model.MensajeServidor, error) {
	query := `SELECT id, remitente_id, destino_usuario_id, canal_id, 
              chat_privado_id, contenido, timestamp, archivo_id 
              FROM mensajes WHERE id = ?`

	row := dao.db.QueryRow(query, id.String())
	return dao.scanMessage(row)
}

// FindByChannelID recupera mensajes de un canal específico
func (dao *MessageMySQLDAO) FindByChannelID(canalID uuid.UUID, limit int) ([]*model.MensajeServidor, error) {
	query := `SELECT id, remitente_id, destino_usuario_id, canal_id, 
              chat_privado_id, contenido, timestamp, archivo_id 
              FROM mensajes WHERE canal_id = ? 
              ORDER BY timestamp DESC LIMIT ?`

	rows, err := dao.db.Query(query, canalID.String(), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMessageRows(rows)
}

// FindByChatPrivadoID recupera mensajes de un chat privado
func (dao *MessageMySQLDAO) FindByChatPrivadoID(chatPrivadoID uuid.UUID, limit int) ([]*model.MensajeServidor, error) {
	query := `SELECT id, remitente_id, destino_usuario_id, canal_id, 
              chat_privado_id, contenido, timestamp, archivo_id 
              FROM mensajes WHERE chat_privado_id = ? 
              ORDER BY timestamp DESC LIMIT ?`

	rows, err := dao.db.Query(query, chatPrivadoID.String(), limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMessageRows(rows)
}

// FindDirectMessages recupera mensajes directos entre dos usuarios
func (dao *MessageMySQLDAO) FindDirectMessages(senderID, recipientID uuid.UUID, limit int) ([]*model.MensajeServidor, error) {
	query := `SELECT id, remitente_id, destino_usuario_id, canal_id, 
              chat_privado_id, contenido, timestamp, archivo_id 
              FROM mensajes 
              WHERE (remitente_id = ? AND destino_usuario_id = ?) 
              OR (remitente_id = ? AND destino_usuario_id = ?) 
              ORDER BY timestamp DESC LIMIT ?`

	rows, err := dao.db.Query(query,
		senderID.String(), recipientID.String(),
		recipientID.String(), senderID.String(),
		limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.scanMessageRows(rows)
}

// Delete elimina un mensaje de la base de datos
func (dao *MessageMySQLDAO) Delete(id uuid.UUID) error {
	query := `DELETE FROM mensajes WHERE id = ?`
	_, err := dao.db.Exec(query, id.String())
	return err
}

// Helper para escanear una fila en un mensaje
func (dao *MessageMySQLDAO) scanMessage(row *sql.Row) (*model.MensajeServidor, error) {
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

	return dao.buildMessageFromScannedData(
		idStr, remitenteIDStr, destinoUsuarioIDStr,
		canalIDStr, chatPrivadoIDStr, contenido,
		timestamp, archivoIDStr)
}

// Helper para escanear filas en mensajes
func (dao *MessageMySQLDAO) scanMessageRow(rows *sql.Rows) (*model.MensajeServidor, error) {
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

	return dao.buildMessageFromScannedData(
		idStr, remitenteIDStr, destinoUsuarioIDStr,
		canalIDStr, chatPrivadoIDStr, contenido,
		timestamp, archivoIDStr)
}

// Helper para escanear múltiples filas
func (dao *MessageMySQLDAO) scanMessageRows(rows *sql.Rows) ([]*model.MensajeServidor, error) {
	var mensajes []*model.MensajeServidor
	for rows.Next() {
		mensaje, err := dao.scanMessageRow(rows)
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

// Helper para convertir valores escaneados en un objeto MensajeServidor
func (dao *MessageMySQLDAO) buildMessageFromScannedData(
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
