package factory

import (
	"github.com/google/uuid"
	"model"
)

// IChatPrivadoFactory define la creación de ChatPrivado.
type IChatPrivadoFactory interface {
	// Create genera un ChatPrivado con un ID aleatorio.
	Create() (*model.ChatPrivado, error)
	
	// CreateWithParticipants genera un ChatPrivado y lo asocia con dos usuarios.
	CreateWithParticipants(usuario1ID, usuario2ID uuid.UUID) (*model.ChatPrivado, []*model.ChatPrivadoUsuario, error)
}

// chatPrivadoFactory es la implementación de IChatPrivadoFactory.
type chatPrivadoFactory struct{}

// NewChatPrivadoFactory devuelve una instancia de IChatPrivadoFactory.
func NewChatPrivadoFactory() IChatPrivadoFactory {
	return &chatPrivadoFactory{}
}

// Create implementa IChatPrivadoFactory: crea un ChatPrivado con ID aleatorio.
func (f *chatPrivadoFactory) Create() (*model.ChatPrivado, error) {
	id := uuid.New()
	return model.NewChatPrivado(id)
}

// CreateWithParticipants implementa IChatPrivadoFactory: crea un ChatPrivado y lo asocia con dos usuarios.
func (f *chatPrivadoFactory) CreateWithParticipants(
	usuario1ID, usuario2ID uuid.UUID,
) (*model.ChatPrivado, []*model.ChatPrivadoUsuario, error) {
	// Validar que los usuarios no sean iguales
	if usuario1ID == usuario2ID {
		return nil, nil, model.ErrChatPrivadoUsuariosIguales
	}

	chat, err := f.Create()
	if err != nil {
		return nil, nil, err
	}
	
	// Crear las relaciones con los usuarios
	relacion1, err := model.NewChatPrivadoUsuario(chat.ID(), usuario1ID)
	if err != nil {
		return nil, nil, err
	}
	
	relacion2, err := model.NewChatPrivadoUsuario(chat.ID(), usuario2ID)
	if err != nil {
		return nil, nil, err
	}
	
	relaciones := []*model.ChatPrivadoUsuario{relacion1, relacion2}
	return chat, relaciones, nil
}
