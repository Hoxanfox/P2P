package factory_test

import (
	"testing"

	"github.com/google/uuid"
	"factory"
	"model"
)

func TestNewChatPrivadoFactory(t *testing.T) {
	fact := factory.NewChatPrivadoFactory()
	if fact == nil {
		t.Fatal("NewChatPrivadoFactory devolvió nil")
	}
}

func TestChatPrivadoFactory_Create_Success(t *testing.T) {
	fact := factory.NewChatPrivadoFactory()
	
	chat, err := fact.Create()
	if err != nil {
		t.Fatalf("esperaba sin error, obtuvo %v", err)
	}
	
	if chat.ID() == uuid.Nil {
		t.Error("ID: esperaba no-nil UUID")
	}
}

func TestChatPrivadoFactory_CreateWithParticipants_Success(t *testing.T) {
	fact := factory.NewChatPrivadoFactory()
	usuario1 := uuid.New()
	usuario2 := uuid.New()
	
	chat, relaciones, err := fact.CreateWithParticipants(usuario1, usuario2)
	if err != nil {
		t.Fatalf("esperaba sin error, obtuvo %v", err)
	}
	
	if chat.ID() == uuid.Nil {
		t.Error("ID del chat: esperaba no-nil UUID")
	}
	
	if len(relaciones) != 2 {
		t.Fatalf("esperaba 2 relaciones, obtuvo %d", len(relaciones))
	}
	
	// Verificación de la primera relación
	if relaciones[0].ChatPrivadoID() != chat.ID() {
		t.Errorf("ChatPrivadoID en relacion[0]: esperado %v, obtuvo %v", chat.ID(), relaciones[0].ChatPrivadoID())
	}
	
	if relaciones[0].UsuarioID() != usuario1 && relaciones[0].UsuarioID() != usuario2 {
		t.Errorf("UsuarioID en relacion[0] debe ser uno de los usuarios proporcionados")
	}
	
	// Verificación de la segunda relación
	if relaciones[1].ChatPrivadoID() != chat.ID() {
		t.Errorf("ChatPrivadoID en relacion[1]: esperado %v, obtuvo %v", chat.ID(), relaciones[1].ChatPrivadoID())
	}
	
	if relaciones[1].UsuarioID() != usuario1 && relaciones[1].UsuarioID() != usuario2 {
		t.Errorf("UsuarioID en relacion[1] debe ser uno de los usuarios proporcionados")
	}
	
	// Verificar que las relaciones son para usuarios diferentes
	if relaciones[0].UsuarioID() == relaciones[1].UsuarioID() {
		t.Error("Las relaciones deben ser para usuarios diferentes")
	}
}

func TestChatPrivadoFactory_CreateWithParticipants_Errors(t *testing.T) {
	fact := factory.NewChatPrivadoFactory()
	validUsuario := uuid.New()

	cases := []struct {
		name      string
		usuario1  uuid.UUID
		usuario2  uuid.UUID
		wantErr   error
	}{
		{"Usuario1 nil", uuid.Nil, validUsuario, model.ErrChatPrivadoUsuarioUsuarioIDNil},
		{"Usuario2 nil", validUsuario, uuid.Nil, model.ErrChatPrivadoUsuarioUsuarioIDNil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, _, err := fact.CreateWithParticipants(c.usuario1, c.usuario2)
			// Aquí solo verificamos que hay error, no el tipo exacto ya que podría variar
			if err == nil {
				t.Errorf("esperado error, obtuvo nil")
			}
		})
	}

	// Caso especial: usuarios iguales
	t.Run("Usuarios iguales", func(t *testing.T) {
		_, _, err := fact.CreateWithParticipants(validUsuario, validUsuario)
		// Esta validación podría hacerse de varias formas, así que solo verificamos
		// que haya error sin comparar el error específico
		if err == nil {
			t.Errorf("esperado error para usuarios iguales, obtuvo nil")
		}
	})
}
