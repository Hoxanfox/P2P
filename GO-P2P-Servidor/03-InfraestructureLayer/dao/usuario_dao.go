package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"model"
	"pool"
)

// UsuarioDAO gestiona las operaciones de base de datos para entidades UsuarioServidor
type UsuarioDAO struct {
	dbPool *pool.DBConnectionPool
}

// NuevoUsuarioDAO crea una nueva instancia de UsuarioDAO
func NuevoUsuarioDAO(dbPool *pool.DBConnectionPool) *UsuarioDAO {
	return &UsuarioDAO{dbPool: dbPool}
}

// Crear persiste un nuevo usuario en la base de datos
func (dao *UsuarioDAO) Crear(usuario *model.UsuarioServidor) error {
	query := `INSERT INTO usuarios (id, nombre_usuario, email, contrasena_hasheada, 
              foto_url, ip_registrada, fecha_registro, is_connected)
              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := dao.dbPool.DB().Exec(
		query,
		usuario.ID().String(),
		usuario.NombreUsuario(),
		usuario.Email(),
		usuario.ContrasenaHasheada(),
		usuario.FotoURL(),
		usuario.IPRegistrada(),
		usuario.FechaRegistro(),
		usuario.IsConnected(),
	)

	return err
}

// BuscarPorID recupera un usuario por su ID
func (dao *UsuarioDAO) BuscarPorID(id uuid.UUID) (*model.UsuarioServidor, error) {
	query := `SELECT id, nombre_usuario, email, contrasena_hasheada, 
              foto_url, ip_registrada, fecha_registro, is_connected 
              FROM usuarios WHERE id = ?`

	row := dao.dbPool.DB().QueryRow(query, id.String())
	return dao.escanearUsuario(row)
}

// BuscarPorEmail recupera un usuario por su correo electrónico
func (dao *UsuarioDAO) BuscarPorEmail(email string) (*model.UsuarioServidor, error) {
	query := `SELECT id, nombre_usuario, email, contrasena_hasheada, 
              foto_url, ip_registrada, fecha_registro, is_connected 
              FROM usuarios WHERE email = ?`

	row := dao.dbPool.DB().QueryRow(query, email)
	return dao.escanearUsuario(row)
}

// Actualizar actualiza un usuario existente en la base de datos
func (dao *UsuarioDAO) Actualizar(usuario *model.UsuarioServidor) error {
	query := `UPDATE usuarios 
              SET nombre_usuario = ?, email = ?, contrasena_hasheada = ?,
              foto_url = ?, ip_registrada = ?, is_connected = ?
              WHERE id = ?`

	_, err := dao.dbPool.DB().Exec(
		query,
		usuario.NombreUsuario(),
		usuario.Email(),
		usuario.ContrasenaHasheada(),
		usuario.FotoURL(),
		usuario.IPRegistrada(),
		usuario.IsConnected(),
		usuario.ID().String(),
	)

	return err
}

// Eliminar borra un usuario de la base de datos
func (dao *UsuarioDAO) Eliminar(id uuid.UUID) error {
	query := `DELETE FROM usuarios WHERE id = ?`
	_, err := dao.dbPool.DB().Exec(query, id.String())
	return err
}

// BuscarTodos recupera todos los usuarios de la base de datos
func (dao *UsuarioDAO) BuscarTodos() ([]*model.UsuarioServidor, error) {
	query := `SELECT id, nombre_usuario, email, contrasena_hasheada, 
              foto_url, ip_registrada, fecha_registro, is_connected 
              FROM usuarios`

	rows, err := dao.dbPool.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []*model.UsuarioServidor
	for rows.Next() {
		usuario, err := dao.escanearUsuarioFila(rows)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usuarios, nil
}

// Método auxiliar para escanear una fila en un usuario
func (dao *UsuarioDAO) escanearUsuario(row *sql.Row) (*model.UsuarioServidor, error) {
	var (
		idStr, nombreUsuario, email, contrasenaHasheada string
		fotoURL, ipRegistrada                           string
		fechaRegistro                                   time.Time
		isConnected                                     bool
	)

	if err := row.Scan(
		&idStr,
		&nombreUsuario,
		&email,
		&contrasenaHasheada,
		&fotoURL,
		&ipRegistrada,
		&fechaRegistro,
		&isConnected,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuario no encontrado
		}
		return nil, err
	}

	idParseado, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	usuario, err := model.NewUsuarioServidor(
		idParseado,
		nombreUsuario,
		email,
		contrasenaHasheada,
		fotoURL,
		ipRegistrada,
		fechaRegistro,
	)

	if err != nil {
		return nil, err
	}

	if isConnected {
		usuario.SetConnected(true)
	}

	return usuario, nil
}

// Método auxiliar para escanear una fila de rows.Next()
func (dao *UsuarioDAO) escanearUsuarioFila(rows *sql.Rows) (*model.UsuarioServidor, error) {
	var (
		idStr, nombreUsuario, email, contrasenaHasheada string
		fotoURL, ipRegistrada                           string
		fechaRegistro                                   time.Time
		isConnected                                     bool
	)

	if err := rows.Scan(
		&idStr,
		&nombreUsuario,
		&email,
		&contrasenaHasheada,
		&fotoURL,
		&ipRegistrada,
		&fechaRegistro,
		&isConnected,
	); err != nil {
		return nil, err
	}

	idParseado, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	usuario, err := model.NewUsuarioServidor(
		idParseado,
		nombreUsuario,
		email,
		contrasenaHasheada,
		fotoURL,
		ipRegistrada,
		fechaRegistro,
	)

	if err != nil {
		return nil, err
	}

	if isConnected {
		usuario.SetConnected(true)
	}

	return usuario, nil}
