package model_test

import (
	"testing"

	"github.com/google/uuid"
	"model"
)

func TestNewChatPrivado_Success(t *testing.T) {
	id := uuid.New()

	chat, err := model.NewChatPrivado(id)
	if err != nil {
		t.Fatalf("esperaba sin error, obtuvo %v", err)
	}

	if chat.ID() != id {
		t.Errorf("ID: esperado %v, obtuvo %v", id, chat.ID())
	}
}

func TestNewChatPrivado_Errors(t *testing.T) {
	cases := []struct {
		name    string
		id      uuid.UUID
		wantErr error
	}{
		{"ID nil", uuid.Nil, model.ErrChatPrivadoIDNil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := model.NewChatPrivado(c.id)
			if err != c.wantErr {
				t.Errorf("esperado error %v, obtuvo %v", c.wantErr, err)
			}
		})
	}
}
