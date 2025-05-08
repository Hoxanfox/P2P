package dao

import (
	"database/sql"
	"github.com/google/uuid"
	"model"
	"pool"
)

// NodoDAO maneja las operaciones de base de datos para entidades Peer (Nodo)
type NodoDAO struct {
	dbPool *pool.DBConnectionPool
}

// NuevoNodoDAO crea una nueva instancia de NodoDAO
func NuevoNodoDAO(dbPool *pool.DBConnectionPool) *NodoDAO {
	return &NodoDAO{dbPool: dbPool}
}

// Guardar persiste un nodo en la base de datos
func (dao *NodoDAO) Guardar(nodo *model.Peer) error {
	query := `INSERT INTO peers (id_nodo, direccion, estado)
              VALUES (?, ?, ?)`

	_, err := dao.dbPool.DB().Exec(
		query,
		nodo.IDNodo().String(),
		nodo.Direccion(),
		string(nodo.Estado()), // NodoEstado es un tipo string
	)

	return err
}

// BuscarPorID recupera un nodo por su ID
func (dao *NodoDAO) BuscarPorID(id uuid.UUID) (*model.Peer, error) {
	query := `SELECT id_nodo, direccion, estado
              FROM peers WHERE id_nodo = ?`

	row := dao.dbPool.DB().QueryRow(query, id.String())
	return dao.escanearNodo(row)
}

// BuscarTodos recupera todos los nodos de la base de datos
func (dao *NodoDAO) BuscarTodos() ([]*model.Peer, error) {
	query := `SELECT id_nodo, direccion, estado
              FROM peers`

	rows, err := dao.dbPool.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesNodos(rows)
}

// Actualizar actualiza un nodo existente
func (dao *NodoDAO) Actualizar(nodo *model.Peer) error {
	query := `UPDATE peers SET direccion = ?, estado = ?
              WHERE id_nodo = ?`

	_, err := dao.dbPool.DB().Exec(
		query,
		nodo.Direccion(),
		string(nodo.Estado()), // NodoEstado es un tipo string
		nodo.IDNodo().String(),
	)

	return err
}

// Eliminar elimina un nodo de la base de datos
func (dao *NodoDAO) Eliminar(id uuid.UUID) error {
	query := `DELETE FROM peers WHERE id_nodo = ?`
	_, err := dao.dbPool.DB().Exec(query, id.String())
	return err
}

// BuscarPorEstado recupera todos los nodos con un estado específico
func (dao *NodoDAO) BuscarPorEstado(estado model.NodoEstado) ([]*model.Peer, error) {
	query := `SELECT id_nodo, direccion, estado
              FROM peers WHERE estado = ?`

	rows, err := dao.dbPool.DB().Query(query, string(estado))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return dao.escanearMultiplesNodos(rows)
}

// escanearNodo escanea una fila en un objeto Peer (Nodo)
func (dao *NodoDAO) escanearNodo(row *sql.Row) (*model.Peer, error) {
	var (
		idStr     string
		direccion string
		estadoStr string
	)

	if err := row.Scan(&idStr, &direccion, &estadoStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Nodo no encontrado
		}
		return nil, err
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	// Convertir el string a NodoEstado
	estado := model.NodoEstado(estadoStr)

	return model.NewPeer(id, direccion, estado)
}

// escanearMultiplesNodos escanea múltiples filas en objetos Peer (Nodo)
func (dao *NodoDAO) escanearMultiplesNodos(rows *sql.Rows) ([]*model.Peer, error) {
	var nodos []*model.Peer

	for rows.Next() {
		var (
			idStr     string
			direccion string
			estadoStr string
		)

		if err := rows.Scan(&idStr, &direccion, &estadoStr); err != nil {
			return nil, err
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		// Convertir el string a NodoEstado
		estado := model.NodoEstado(estadoStr)

		nodo, err := model.NewPeer(id, direccion, estado)
		if err != nil {
			return nil, err
		}

		nodos = append(nodos, nodo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nodos, nil
}
