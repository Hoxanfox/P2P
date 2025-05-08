
package repository

import (
	"context"

	"dao"
	"github.com/google/uuid"
	"model"
)

// ChannelRepository implementa la interfaz IChannelRepository del dominio
// utilizando los DAOs correspondientes para las operaciones de base de datos
type ChannelRepository struct {
	canalDAO           *dao.CanalDAO
	invitacionCanalDAO *dao.InvitacionCanalDAO
	canalMiembroDAO    *dao.CanalMiembroDAO
}

// NewChannelRepository crea una nueva instancia de ChannelRepository
func NewChannelRepository(
	canalDAO *dao.CanalDAO,
	invitacionCanalDAO *dao.InvitacionCanalDAO,
	canalMiembroDAO *dao.CanalMiembroDAO,
) *ChannelRepository {
	return &ChannelRepository{
		canalDAO:           canalDAO,
		invitacionCanalDAO: invitacionCanalDAO,
		canalMiembroDAO:    canalMiembroDAO,
	}
}

// Save persiste un canal en la base de datos
func (r *ChannelRepository) Save(ctx context.Context, c *model.CanalServidor) error {
	return r.canalDAO.Crear(c)
}

// Update actualiza un canal existente en la base de datos
func (r *ChannelRepository) Update(ctx context.Context, c *model.CanalServidor) error {
	return r.canalDAO.Actualizar(c)
}

// Delete elimina un canal de la base de datos
func (r *ChannelRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.canalDAO.Eliminar(id)
}

// FindByID busca un canal por su ID
func (r *ChannelRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.CanalServidor, error) {
	return r.canalDAO.BuscarPorID(id)
}

// FindAll recupera todos los canales
func (r *ChannelRepository) FindAll(ctx context.Context) ([]*model.CanalServidor, error) {
	return r.canalDAO.BuscarTodos()
}

// AddMember añade un miembro a un canal
func (r *ChannelRepository) AddMember(ctx context.Context, channelID, userID uuid.UUID, rol string) error {
	miembro, err := model.NewCanalMiembro(channelID, userID, rol)
	if err != nil {
		return err
	}
	return r.canalMiembroDAO.Guardar(miembro)
}

// RemoveMember elimina un miembro de un canal
func (r *ChannelRepository) RemoveMember(ctx context.Context, channelID, userID uuid.UUID) error {
	return r.canalMiembroDAO.Eliminar(channelID, userID)
}

// ListMembers lista los miembros de un canal
func (r *ChannelRepository) ListMembers(ctx context.Context, channelID uuid.UUID) ([]uuid.UUID, error) {
	miembros, err := r.canalMiembroDAO.BuscarPorCanalID(channelID)
	if err != nil {
		return nil, err
	}

	var userIDs []uuid.UUID
	for _, miembro := range miembros {
		userIDs = append(userIDs, miembro.UsuarioID())
	}

	return userIDs, nil
}

// SaveInvitation guarda una invitación a un canal
func (r *ChannelRepository) SaveInvitation(ctx context.Context, inv *model.InvitacionCanal) error {
	return r.invitacionCanalDAO.Guardar(inv)
}

// UpdateInvitation actualiza una invitación a un canal
func (r *ChannelRepository) UpdateInvitation(ctx context.Context, inv *model.InvitacionCanal) error {
	return r.invitacionCanalDAO.Actualizar(inv)
}

// ListInvitations lista las invitaciones de un canal
func (r *ChannelRepository) ListInvitations(ctx context.Context, channelID uuid.UUID) ([]*model.InvitacionCanal, error) {
	return r.invitacionCanalDAO.BuscarPorCanalID(channelID)
}
