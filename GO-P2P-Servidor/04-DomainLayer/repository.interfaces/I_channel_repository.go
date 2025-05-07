package repository

import (
	"context"
	
	"github.com/google/uuid"
	"model"
)

// IChannelRepository define las operaciones para el repositorio de canales
type IChannelRepository interface {
    // CRUD
    Save(ctx context.Context, c *model.CanalServidor) error
    Update(ctx context.Context, c *model.CanalServidor) error
    Delete(ctx context.Context, id uuid.UUID) error
    FindByID(ctx context.Context, id uuid.UUID) (*model.CanalServidor, error)
    FindAll(ctx context.Context) ([]*model.CanalServidor, error)

    // Miembros e invitaciones
    AddMember(ctx context.Context, channelID, userID uuid.UUID, rol string) error
    RemoveMember(ctx context.Context, channelID, userID uuid.UUID) error
    ListMembers(ctx context.Context, channelID uuid.UUID) ([]uuid.UUID, error)

    SaveInvitation(ctx context.Context, inv *model.InvitacionCanal) error
    UpdateInvitation(ctx context.Context, inv *model.InvitacionCanal) error
    ListInvitations(ctx context.Context, channelID uuid.UUID) ([]*model.InvitacionCanal, error)
}
