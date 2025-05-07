package model_test

import (
	"testing"

	"github.com/google/uuid"
	"model"
)

func TestNewChatPrivadoUsuario_Success(t *testing.T) {
	chatID := uuid.New()
	usuarioID := uuid.New()

	relacion, err := model.NewChatPrivadoUsuario(chatID, usuarioID)
	if err != nil {
		t.Fatalf("esperaba sin error, obtuvo %v", err)
	}

	if relacion.ChatPrivadoID() != chatID {
		t.Errorf("ChatPrivadoID: esperado %v, obtuvo %v", chatID, relacion.ChatPrivadoID())
	}

	if relacion.UsuarioID() != usuarioID {
		t.Errorf("UsuarioID: esperado %v, obtuvo %v", usuarioID, relacion.UsuarioID())
	}
}

func TestNewChatPrivadoUsuario_Errors(t *testing.T) {
	validChatID := uuid.New()
	validUsuarioID := uuid.New()

	cases := []struct {
		name      string
		chatID    uuid.UUID
		usuarioID uuid.UUID
		wantErr   error
	}{
		{"ChatID nil", uuid.Nil, validUsuarioID, model.ErrChatPrivadoUsuarioChatIDNil},
		{"UsuarioID nil", validChatID, uuid.Nil, model.ErrChatPrivadoUsuarioUsuarioIDNil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := model.NewChatPrivadoUsuario(c.chatID, c.usuarioID)
			if err != c.wantErr {
				t.Errorf("esperado error %v, obtuvo %v", c.wantErr, err)
			}
		})
	}
}
