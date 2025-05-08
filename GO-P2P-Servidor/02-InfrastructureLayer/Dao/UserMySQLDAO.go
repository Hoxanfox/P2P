package Dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/P2P/GO-P2P-SERVIDOR/04-DomainLayer/model"
)

// UserMySQLDAO handles database operations for UsuarioServidor entities
type UserMySQLDAO struct {
	db *sql.DB
}

// NewUserMySQLDAO creates a new UserMySQLDAO instance
func NewUserMySQLDAO(db *sql.DB) *UserMySQLDAO {
	return &UserMySQLDAO{db: db}
}

// Create persists a new user to the database
func (dao *UserMySQLDAO) Create(user *model.UsuarioServidor) error {
	query := `INSERT INTO usuarios (id, nombre_usuario, email, contrasena_hasheada, 
              foto_url, ip_registrada, fecha_registro, is_connected)
              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := dao.db.Exec(
		query,
		user.ID().String(),
		user.NombreUsuario(),
		user.Email(),
		user.ContrasenaHasheada(),
		user.FotoURL(),
		user.IPRegistrada(),
		user.FechaRegistro(),
		user.IsConnected(),
	)

	return err
}

// FindByID retrieves a user by their ID
func (dao *UserMySQLDAO) FindByID(id uuid.UUID) (*model.UsuarioServidor, error) {
	query := `SELECT id, nombre_usuario, email, contrasena_hasheada, 
              foto_url, ip_registrada, fecha_registro, is_connected 
              FROM usuarios WHERE id = ?`

	row := dao.db.QueryRow(query, id.String())
	return dao.scanUser(row)
}

// FindByEmail retrieves a user by their email
func (dao *UserMySQLDAO) FindByEmail(email string) (*model.UsuarioServidor, error) {
	query := `SELECT id, nombre_usuario, email, contrasena_hasheada, 
              foto_url, ip_registrada, fecha_registro, is_connected 
              FROM usuarios WHERE email = ?`

	row := dao.db.QueryRow(query, email)
	return dao.scanUser(row)
}

// Update updates an existing user in the database
func (dao *UserMySQLDAO) Update(user *model.UsuarioServidor) error {
	query := `UPDATE usuarios 
              SET nombre_usuario = ?, email = ?, contrasena_hasheada = ?,
              foto_url = ?, ip_registrada = ?, is_connected = ?
              WHERE id = ?`

	_, err := dao.db.Exec(
		query,
		user.NombreUsuario(),
		user.Email(),
		user.ContrasenaHasheada(),
		user.FotoURL(),
		user.IPRegistrada(),
		user.IsConnected(),
		user.ID().String(),
	)

	return err
}

// Delete removes a user from the database
func (dao *UserMySQLDAO) Delete(id uuid.UUID) error {
	query := `DELETE FROM usuarios WHERE id = ?`
	_, err := dao.db.Exec(query, id.String())
	return err
}

// FindAll retrieves all users from the database
func (dao *UserMySQLDAO) FindAll() ([]*model.UsuarioServidor, error) {
	query := `SELECT id, nombre_usuario, email, contrasena_hasheada, 
              foto_url, ip_registrada, fecha_registro, is_connected 
              FROM usuarios`

	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.UsuarioServidor
	for rows.Next() {
		user, err := dao.scanUserRow(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Helper method to scan a row into a user
func (dao *UserMySQLDAO) scanUser(row *sql.Row) (*model.UsuarioServidor, error) {
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
			return nil, nil // User not found
		}
		return nil, err
	}

	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	user, err := model.NewUsuarioServidor(
		parsedID,
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
		user.SetConnected(true)
	}

	return user, nil
}

// Helper method to scan a row from rows.Next()
func (dao *UserMySQLDAO) scanUserRow(rows *sql.Rows) (*model.UsuarioServidor, error) {
	var (
		idStr, nombreUsuario, email, contrasenaHasheada string
		fotoURL, ipRegistrada                           string
		fechaRegistro