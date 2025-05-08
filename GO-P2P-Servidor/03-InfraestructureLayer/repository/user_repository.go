package repository

import (
	"context"
	"dao"
	"github.com/google/uuid"
	"model"
)

// UserRepository implementa la interfaz IUserRepository del dominio
// utilizando el UsuarioDAO para las operaciones de base de datos
type UserRepository struct {
	usuarioDAO *dao.UsuarioDAO
}

// NewUserRepository crea una nueva instancia de UserRepository
func NewUserRepository(usuarioDAO *dao.UsuarioDAO) *UserRepository {
	return &UserRepository{
		usuarioDAO: usuarioDAO,
	}
}

// Save persiste un usuario en la base de datos
func (r *UserRepository) Save(ctx context.Context, u *model.UsuarioServidor) error {
	// Usamos el método Crear del DAO
	return r.usuarioDAO.Crear(u)
}

// Update actualiza un usuario existente en la base de datos
func (r *UserRepository) Update(ctx context.Context, u *model.UsuarioServidor) error {
	// Usamos el método Actualizar del DAO
	return r.usuarioDAO.Actualizar(u)
}

// Delete elimina un usuario de la base de datos
func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	// Usamos el método Eliminar del DAO
	return r.usuarioDAO.Eliminar(id)
}

// FindByID busca un usuario por su ID
func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.UsuarioServidor, error) {
	// Usamos el método BuscarPorID del DAO
	return r.usuarioDAO.BuscarPorID(id)
}

// FindAll recupera todos los usuarios
func (r *UserRepository) FindAll(ctx context.Context) ([]*model.UsuarioServidor, error) {
	// Usamos el método BuscarTodos del DAO
	return r.usuarioDAO.BuscarTodos()
}

// FindConnected recupera todos los usuarios conectados
func (r *UserRepository) FindConnected(ctx context.Context) ([]*model.UsuarioServidor, error) {
	// Obtenemos todos los usuarios primero
	usuarios, err := r.usuarioDAO.BuscarTodos()
	if err != nil {
		return nil, err
	}

	// Filtramos solo los usuarios conectados
	var usuariosConectados []*model.UsuarioServidor
	for _, usuario := range usuarios {
		if usuario.IsConnected() {
			usuariosConectados = append(usuariosConectados, usuario)
		}
	}

	return usuariosConectados, nil
}

// FindByEmail busca un usuario por su email
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.UsuarioServidor, error) {
	// Usamos el método BuscarPorEmail del DAO
	return r.usuarioDAO.BuscarPorEmail(email)
}
