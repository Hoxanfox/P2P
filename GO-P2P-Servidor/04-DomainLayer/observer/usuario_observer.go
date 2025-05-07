package observer

import (
	"model"
)

// IUserObserver define los callbacks para todos los eventos relacionados con usuarios
// como registro, login/logout, actualización de perfil e invitaciones a canales.
type IUserObserver interface {
	// OnUserRegistered se invoca justo después de crear un nuevo usuario
	OnUserRegistered(user *model.UsuarioServidor)

	// OnUserLoggedIn se invoca tras cada login de usuario
	OnUserLoggedIn(user *model.UsuarioServidor)

	// OnUserLoggedOut se invoca tras cada logout de usuario
	OnUserLoggedOut(user *model.UsuarioServidor)

	// OnUserUpdated se invoca al modificar datos del perfil
	OnUserUpdated(user *model.UsuarioServidor, changedFields []string)

	// OnInvitationSent se invoca cuando se envía una invitación a un canal
	OnInvitationSent(canal *model.CanalServidor, invitedUser *model.UsuarioServidor, 
		byUser *model.UsuarioServidor)

	// OnInvitationResponded se invoca cuando un usuario acepta o rechaza una invitación
	OnInvitationResponded(canal *model.CanalServidor, user *model.UsuarioServidor, 
		accepted bool)
}
