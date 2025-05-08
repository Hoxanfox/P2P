package dao

import (
	"database/sql"
	"pool"

	"github.com/google/uuid"
	"model"
)

// ChatPrivadoDAO maneja las operaciones de base de datos para entidades ChatPrivado
type ChatPrivadoDAO struct {
	dbPool *pool.DBConnectionPool
}

// NuevoChatPrivadoDAO crea una nueva instancia de ChatPrivadoDAO
func NuevoChatPrivadoDAO(dbPool *pool.DBConnectionPool) *ChatPrivadoDAO {
	return &ChatPrivadoDAO{dbPool: dbPool}
}

// Guardar persiste un chat privado en la base de datos
func (dao *ChatPrivadoDAO) Guardar(chat *model.ChatPrivado) error {
	query := `INSERT INTO chats_privados (id) VALUES (?)`

	_, err := dao.dbPool.DB().Exec(
		query,
		chat.ID().String(),
	)

	return err
}

// BuscarPorID recupera un chat privado por su ID
func (dao *ChatPrivadoDAO) BuscarPorID(id uuid.UUID) (*model.ChatPrivado, error) {
	query := `SELECT id FROM chats_privados WHERE id = ?`

	row := dao.dbPool.DB().QueryRow(query, id.String())
	return dao.escanearChatPrivado(row)
}

// BuscarTodos recupera todos los chats privados
func (dao *ChatPrivadoDAO) BuscarTodos() ([]*model.ChatPrivado, error) {
	query := `SELECT id FROM chats_privados`

	rows, err := dao.dbPool.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesChatsPrivados(rows)
}

// Eliminar elimina un chat privado
func (dao *ChatPrivadoDAO) Eliminar(id uuid.UUID) error {
	query := `DELETE FROM chats_privados WHERE id = ?`
	_, err := dao.dbPool.DB().Exec(query, id.String())
	return err
}

// escanearChatPrivado escanea una fila en un objeto ChatPrivado
func (dao *ChatPrivadoDAO) escanearChatPrivado(row *sql.Row) (*model.ChatPrivado, error) {
	var idStr string

	if err := row.Scan(&idStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Chat privado no encontrado
		}
		return nil, err
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	return model.NewChatPrivado(id)
}

// escanearMultiplesChatsPrivados escanea múltiples filas en objetos ChatPrivado
func (dao *ChatPrivadoDAO) escanearMultiplesChatsPrivados(rows *sql.Rows) ([]*model.ChatPrivado, error) {
	var chats []*model.ChatPrivado

	for rows.Next() {
		var idStr string

		if err := rows.Scan(&idStr); err != nil {
			return nil, err
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		chat, err := model.NewChatPrivado(id)
		if err != nil {
			return nil, err
		}

		chats = append(chats, chat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chats, nil
}
