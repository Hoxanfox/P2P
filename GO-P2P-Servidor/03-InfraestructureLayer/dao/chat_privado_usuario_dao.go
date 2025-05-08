package dao

import (
	"database/sql"
	"pool"

	"github.com/google/uuid"
	"model"
)

// ChatPrivadoUsuarioDAO maneja las operaciones de base de datos para entidades ChatPrivadoUsuario
type ChatPrivadoUsuarioDAO struct {
	dbPool *pool.DBConnectionPool
}

// NuevoChatPrivadoUsuarioDAO crea una nueva instancia de ChatPrivadoUsuarioDAO
func NuevoChatPrivadoUsuarioDAO(dbPool *pool.DBConnectionPool) *ChatPrivadoUsuarioDAO {
	return &ChatPrivadoUsuarioDAO{dbPool: dbPool}
}

// Guardar persiste una relación entre chat privado y usuario en la base de datos
func (dao *ChatPrivadoUsuarioDAO) Guardar(chatUsuario *model.ChatPrivadoUsuario) error {
	query := `INSERT INTO chat_privado_usuarios (chat_privado_id, usuario_id)
              VALUES (?, ?)`

	_, err := dao.dbPool.DB().Exec(
		query,
		chatUsuario.ChatPrivadoID().String(),
		chatUsuario.UsuarioID().String(),
	)

	return err
}

// BuscarPorIDs recupera una relación entre chat privado y usuario por sus IDs
func (dao *ChatPrivadoUsuarioDAO) BuscarPorIDs(chatPrivadoID, usuarioID uuid.UUID) (*model.ChatPrivadoUsuario, error) {
	query := `SELECT chat_privado_id, usuario_id
              FROM chat_privado_usuarios 
              WHERE chat_privado_id = ? AND usuario_id = ?`

	row := dao.dbPool.DB().QueryRow(query, chatPrivadoID.String(), usuarioID.String())
	return dao.escanearChatPrivadoUsuario(row)
}

// BuscarPorChatPrivadoID recupera todos los usuarios de un chat privado específico
func (dao *ChatPrivadoUsuarioDAO) BuscarPorChatPrivadoID(chatPrivadoID uuid.UUID) ([]*model.ChatPrivadoUsuario, error) {
	query := `SELECT chat_privado_id, usuario_id
              FROM chat_privado_usuarios 
              WHERE chat_privado_id = ?`

	rows, err := dao.dbPool.DB().Query(query, chatPrivadoID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesChatPrivadoUsuarios(rows)
}

// BuscarPorUsuarioID recupera todos los chats privados de un usuario específico
func (dao *ChatPrivadoUsuarioDAO) BuscarPorUsuarioID(usuarioID uuid.UUID) ([]*model.ChatPrivadoUsuario, error) {
	query := `SELECT chat_privado_id, usuario_id
              FROM chat_privado_usuarios 
              WHERE usuario_id = ?`

	rows, err := dao.dbPool.DB().Query(query, usuarioID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesChatPrivadoUsuarios(rows)
}

// BuscarChatEntreUsuarios busca un chat privado entre dos usuarios específicos
func (dao *ChatPrivadoUsuarioDAO) BuscarChatEntreUsuarios(usuarioID1, usuarioID2 uuid.UUID) (*model.ChatPrivado, error) {
	query := `
		SELECT c1.chat_privado_id
		FROM chat_privado_usuarios c1
		JOIN chat_privado_usuarios c2 ON c1.chat_privado_id = c2.chat_privado_id
		WHERE c1.usuario_id = ? AND c2.usuario_id = ?
	`

	row := dao.dbPool.DB().QueryRow(query, usuarioID1.String(), usuarioID2.String())
	
	var chatPrivadoIDStr string
	if err := row.Scan(&chatPrivadoIDStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No se encontró un chat entre estos usuarios
		}
		return nil, err
	}

	chatPrivadoID, err := uuid.Parse(chatPrivadoIDStr)
	if err != nil {
		return nil, err
	}

	return model.NewChatPrivado(chatPrivadoID)
}

// Eliminar elimina una relación entre chat privado y usuario
func (dao *ChatPrivadoUsuarioDAO) Eliminar(chatPrivadoID, usuarioID uuid.UUID) error {
	query := `DELETE FROM chat_privado_usuarios 
              WHERE chat_privado_id = ? AND usuario_id = ?`

	_, err := dao.dbPool.DB().Exec(query, chatPrivadoID.String(), usuarioID.String())
	return err
}

// EliminarPorChatPrivadoID elimina todas las relaciones de un chat privado específico
func (dao *ChatPrivadoUsuarioDAO) EliminarPorChatPrivadoID(chatPrivadoID uuid.UUID) error {
	query := `DELETE FROM chat_privado_usuarios WHERE chat_privado_id = ?`
	_, err := dao.dbPool.DB().Exec(query, chatPrivadoID.String())
	return err
}

// escanearChatPrivadoUsuario escanea una fila en un objeto ChatPrivadoUsuario
func (dao *ChatPrivadoUsuarioDAO) escanearChatPrivadoUsuario(row *sql.Row) (*model.ChatPrivadoUsuario, error) {
	var (
		chatPrivadoIDStr string
		usuarioIDStr     string
	)

	if err := row.Scan(&chatPrivadoIDStr, &usuarioIDStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Relación no encontrada
		}
		return nil, err
	}

	chatPrivadoID, err := uuid.Parse(chatPrivadoIDStr)
	if err != nil {
		return nil, err
	}

	usuarioID, err := uuid.Parse(usuarioIDStr)
	if err != nil {
		return nil, err
	}

	return model.NewChatPrivadoUsuario(chatPrivadoID, usuarioID)
}

// escanearMultiplesChatPrivadoUsuarios escanea múltiples filas en objetos ChatPrivadoUsuario
func (dao *ChatPrivadoUsuarioDAO) escanearMultiplesChatPrivadoUsuarios(rows *sql.Rows) ([]*model.ChatPrivadoUsuario, error) {
	var relaciones []*model.ChatPrivadoUsuario

	for rows.Next() {
		var (
			chatPrivadoIDStr string
			usuarioIDStr     string
		)

		if err := rows.Scan(&chatPrivadoIDStr, &usuarioIDStr); err != nil {
			return nil, err
		}

		chatPrivadoID, err := uuid.Parse(chatPrivadoIDStr)
		if err != nil {
			return nil, err
		}

		usuarioID, err := uuid.Parse(usuarioIDStr)
		if err != nil {
			return nil, err
		}

		relacion, err := model.NewChatPrivadoUsuario(chatPrivadoID, usuarioID)
		if err != nil {
			return nil, err
		}

		relaciones = append(relaciones, relacion)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return relaciones, nil
}
