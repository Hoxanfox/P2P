package factory_test

import (
	"testing"

	"github.com/google/uuid"
	"factory"
	"model"
)

func TestNewCanalMiembroFactory(t *testing.T) {
	fact := factory.NewCanalMiembroFactory()
	if fact == nil {
		t.Fatal("NewCanalMiembroFactory devolvió nil")
	}
}

func TestCanalMiembroFactory_Create_Success(t *testing.T) {
	fact := factory.NewCanalMiembroFactory()
	canalID := uuid.New()
	usuarioID := uuid.New()
	rol := "propietario"

	miembro, err := fact.Create(canalID, usuarioID, rol)
	if err != nil {
		t.Fatalf("esperaba sin error, obtuvo %v", err)
	}
	
	if miembro.CanalID() != canalID {
		t.Errorf("CanalID: esperado %v, obtuvo %v", canalID, miembro.CanalID())
	}
	if miembro.UsuarioID() != usuarioID {
		t.Errorf("UsuarioID: esperado %v, obtuvo %v", usuarioID, miembro.UsuarioID())
	}
	if miembro.Rol() != rol {
		t.Errorf("Rol: esperado %q, obtuvo %q", rol, miembro.Rol())
	}
}

func TestCanalMiembroFactory_CreateOwner_Success(t *testing.T) {
	fact := factory.NewCanalMiembroFactory()
	canalID := uuid.New()
	usuarioID := uuid.New()

	miembro, err := fact.CreateOwner(canalID, usuarioID)
	if err != nil {
		t.Fatalf("esperaba sin error, obtuvo %v", err)
	}
	
	if miembro.CanalID() != canalID {
		t.Errorf("CanalID: esperado %v, obtuvo %v", canalID, miembro.CanalID())
	}
	if miembro.UsuarioID() != usuarioID {
		t.Errorf("UsuarioID: esperado %v, obtuvo %v", usuarioID, miembro.UsuarioID())
	}
	if miembro.Rol() != "owner" {
		t.Errorf("Rol: esperado %q, obtuvo %q", "owner", miembro.Rol())
	}
}

func TestCanalMiembroFactory_CreateMember_Success(t *testing.T) {
	fact := factory.NewCanalMiembroFactory()
	canalID := uuid.New()
	usuarioID := uuid.New()

	miembro, err := fact.CreateMember(canalID, usuarioID)
	if err != nil {
		t.Fatalf("esperaba sin error, obtuvo %v", err)
	}
	
	if miembro.CanalID() != canalID {
		t.Errorf("CanalID: esperado %v, obtuvo %v", canalID, miembro.CanalID())
	}
	if miembro.UsuarioID() != usuarioID {
		t.Errorf("UsuarioID: esperado %v, obtuvo %v", usuarioID, miembro.UsuarioID())
	}
	if miembro.Rol() != "member" {
		t.Errorf("Rol: esperado %q, obtuvo %q", "member", miembro.Rol())
	}
}

func TestCanalMiembroFactory_Create_Errors(t *testing.T) {
	fact := factory.NewCanalMiembroFactory()
	validCanal := uuid.New()
	validUsuario := uuid.New()
	validRol := "member"

	cases := []struct {
		name      string
		canalID   uuid.UUID
		usuarioID uuid.UUID
		rol       string
		wantErr   error
	}{
		{"CanalID nil", uuid.Nil, validUsuario, validRol, model.ErrCanalMiembroCanalIDNil},
		{"UsuarioID nil", validCanal, uuid.Nil, validRol, model.ErrCanalMiembroUsuarioIDNil},
		{"Rol inválido", validCanal, validUsuario, "", model.ErrCanalMiembroRolVacio},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := fact.Create(c.canalID, c.usuarioID, c.rol)
			if err != c.wantErr {
				t.Errorf("esperado error %v, obtuvo %v", c.wantErr, err)
			}
		})
	}
}
