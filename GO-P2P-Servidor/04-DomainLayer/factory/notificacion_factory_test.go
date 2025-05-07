package factory_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"factory"
	"model"
)

func TestNotificacionFactory_Create_Success(t *testing.T) {
	fact := factory.NewNotificacionFactory()
	usuarioID := uuid.New()
	contenido := "Nueva notificación de prueba"
	
	notif, err := fact.Create(usuarioID, contenido)
	if err != nil {
		t.Fatalf("esperaba sin error, obtuvo %v", err)
	}
	
	if notif.ID() == uuid.Nil {
		t.Error("ID: esperaba no-nil UUID")
	}
	if notif.UsuarioID() != usuarioID {
		t.Errorf("UsuarioID: esperado %v, obtuvo %v", usuarioID, notif.UsuarioID())
	}
	if notif.Contenido() != contenido {
		t.Errorf("Contenido: esperado %q, obtuvo %q", contenido, notif.Contenido())
	}
	if notif.Fecha().IsZero() {
		t.Error("Fecha: esperado timestamp no cero")
	}
	if diff := time.Since(notif.Fecha()); diff < 0 || diff > time.Second {
		t.Errorf("Fecha: marca de tiempo inesperada, dif=%v", diff)
	}
	if notif.InvitacionID() != uuid.Nil {
		t.Errorf("InvitacionID: esperado Nil, obtuvo %v", notif.InvitacionID())
	}
	if notif.Leido() {
		t.Error("Leido: esperado false por defecto, obtuvo true")
	}
}

func TestNotificacionFactory_Create_Errors(t *testing.T) {
	fact := factory.NewNotificacionFactory()
	validUsuario := uuid.New()
	validContenido := "Notificación de prueba"

	cases := []struct {
		name         string
		usuarioID    uuid.UUID
		contenido    string
		wantErr      error
	}{
		{"UsuarioID nil", uuid.Nil, validContenido, model.ErrNotificacionUsuarioIDNil},
		{"Contenido vacío", validUsuario, "", model.ErrNotificacionContenidoVacio},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := fact.Create(c.usuarioID, c.contenido)
			if err != c.wantErr {
				t.Errorf("esperado error %v, obtuvo %v", c.wantErr, err)
			}
		})
	}
}
