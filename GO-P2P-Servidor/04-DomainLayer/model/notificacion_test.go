package model_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"model"
)

func TestNewNotificacion_Success(t *testing.T) {
	id := uuid.New()
	usuarioID := uuid.New()
	contenido := "Nueva notificación de prueba"
	fecha := time.Now().UTC().Truncate(time.Second)
	invitacionID := uuid.New() // Opcional, puede ser uuid.Nil

	notificacion, err := model.NewNotificacion(id, usuarioID, contenido, fecha, invitacionID)
	if err != nil {
		t.Fatalf("esperaba sin error, obtuvo %v", err)
	}

	if notificacion.ID() != id {
		t.Errorf("ID: esperado %v, obtuvo %v", id, notificacion.ID())
	}
	if notificacion.UsuarioID() != usuarioID {
		t.Errorf("UsuarioID: esperado %v, obtuvo %v", usuarioID, notificacion.UsuarioID())
	}
	if notificacion.Contenido() != contenido {
		t.Errorf("Contenido: esperado %v, obtuvo %v", contenido, notificacion.Contenido())
	}
	if !notificacion.Fecha().Equal(fecha) {
		t.Errorf("Fecha: esperado %v, obtuvo %v", fecha, notificacion.Fecha())
	}
	if notificacion.Leido() != false {
		t.Errorf("Leido: esperado %v, obtuvo %v", false, notificacion.Leido())
	}
	if notificacion.InvitacionID() != invitacionID {
		t.Errorf("InvitacionID: esperado %v, obtuvo %v", invitacionID, notificacion.InvitacionID())
	}

	// Probar con invitacionID nil
	notificacionSinInvitacion, err := model.NewNotificacion(id, usuarioID, contenido, fecha, uuid.Nil)
	if err != nil {
		t.Fatalf("esperaba sin error para notificación sin invitación, obtuvo %v", err)
	}
	if notificacionSinInvitacion.InvitacionID() != uuid.Nil {
		t.Errorf("InvitacionID: esperado Nil, obtuvo %v", notificacionSinInvitacion.InvitacionID())
	}
}

func TestNewNotificacion_Errors(t *testing.T) {
	validID := uuid.New()
	validUsuarioID := uuid.New()
	validContenido := "Contenido válido"
	validFecha := time.Now().UTC()
	validInvitacionID := uuid.Nil // Opcional

	cases := []struct {
		name         string
		id           uuid.UUID
		usuarioID    uuid.UUID
		contenido    string
		fecha        time.Time
		invitacionID uuid.UUID
		wantErr      error
	}{
		{"ID nil", uuid.Nil, validUsuarioID, validContenido, validFecha, validInvitacionID, model.ErrNotificacionIDNil},
		{"UsuarioID nil", validID, uuid.Nil, validContenido, validFecha, validInvitacionID, model.ErrNotificacionUsuarioIDNil},
		{"Contenido vacío", validID, validUsuarioID, "", validFecha, validInvitacionID, model.ErrNotificacionContenidoVacio},
		{"Fecha cero", validID, validUsuarioID, validContenido, time.Time{}, validInvitacionID, model.ErrNotificacionFechaZero},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := model.NewNotificacion(c.id, c.usuarioID, c.contenido, c.fecha, c.invitacionID)
			if err != c.wantErr {
				t.Errorf("esperado error %v, obtuvo %v", c.wantErr, err)
			}
		})
	}
}

func TestNotificacion_MarcarComoLeida(t *testing.T) {
	id := uuid.New()
	usuarioID := uuid.New()
	contenido := "Notificación para marcar como leída"
	fecha := time.Now().UTC()

	notificacion, _ := model.NewNotificacion(id, usuarioID, contenido, fecha, uuid.Nil)

	// Verificar estado inicial
	if notificacion.Leido() != false {
		t.Errorf("Estado inicial Leido: esperado %v, obtuvo %v", false, notificacion.Leido())
	}

	// Marcar como leída
	notificacion.MarcarComoLeida()
	if notificacion.Leido() != true {
		t.Errorf("Después de MarcarComoLeida: esperado %v, obtuvo %v", true, notificacion.Leido())
	}

	// Marcar como no leída
	notificacion.MarcarComoNoLeida()
	if notificacion.Leido() != false {
		t.Errorf("Después de MarcarComoNoLeida: esperado %v, obtuvo %v", false, notificacion.Leido())
	}
}
