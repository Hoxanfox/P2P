package factory_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"factory"
	"model"
)

func TestNewInvitacionCanalFactory(t *testing.T) {
	fact := factory.NewInvitacionCanalFactory()
	if fact == nil {
		t.Fatal("NewInvitacionCanalFactory devolvió nil")
	}
}

func TestInvitacionCanalFactory_Create_Success(t *testing.T) {
	fact := factory.NewInvitacionCanalFactory()
	canalID := uuid.New()
	destinatarioID := uuid.New()

	invitacion, err := fact.Create(canalID, destinatarioID)
	if err != nil {
		t.Fatalf("esperaba sin error, obtuvo %v", err)
	}
	
	if invitacion.ID() == uuid.Nil {
		t.Error("ID: esperaba no-nil UUID")
	}
	if invitacion.CanalID() != canalID {
		t.Errorf("CanalID: esperado %v, obtuvo %v", canalID, invitacion.CanalID())
	}
	if invitacion.DestinatarioID() != destinatarioID {
		t.Errorf("DestinatarioID: esperado %v, obtuvo %v", destinatarioID, invitacion.DestinatarioID())
	}
	if invitacion.Estado() != model.InvitacionPendiente {
		t.Errorf("Estado: esperado %q, obtuvo %q", model.InvitacionPendiente, invitacion.Estado())
	}
	if invitacion.FechaEnvio().IsZero() {
		t.Error("FechaEnvio: esperado timestamp no cero")
	}
	if diff := time.Since(invitacion.FechaEnvio()); diff < 0 || diff > time.Second {
		t.Errorf("FechaEnvio: marca de tiempo inesperada, dif=%v", diff)
	}
}

func TestInvitacionCanalFactory_CreateWithStatus_Success(t *testing.T) {
	fact := factory.NewInvitacionCanalFactory()
	canalID := uuid.New()
	destinatarioID := uuid.New()
	estado := model.InvitacionAceptada

	invitacion, err := fact.CreateWithStatus(canalID, destinatarioID, estado)
	if err != nil {
		t.Fatalf("esperaba sin error, obtuvo %v", err)
	}
	
	if invitacion.ID() == uuid.Nil {
		t.Error("ID: esperaba no-nil UUID")
	}
	if invitacion.CanalID() != canalID {
		t.Errorf("CanalID: esperado %v, obtuvo %v", canalID, invitacion.CanalID())
	}
	if invitacion.DestinatarioID() != destinatarioID {
		t.Errorf("DestinatarioID: esperado %v, obtuvo %v", destinatarioID, invitacion.DestinatarioID())
	}
	if invitacion.Estado() != estado {
		t.Errorf("Estado: esperado %q, obtuvo %q", estado, invitacion.Estado())
	}
	if invitacion.FechaEnvio().IsZero() {
		t.Error("FechaEnvio: esperado timestamp no cero")
	}
	if diff := time.Since(invitacion.FechaEnvio()); diff < 0 || diff > time.Second {
		t.Errorf("FechaEnvio: marca de tiempo inesperada, dif=%v", diff)
	}
}

func TestInvitacionCanalFactory_Create_Errors(t *testing.T) {
	fact := factory.NewInvitacionCanalFactory()
	validCanal := uuid.New()
	validDestinatario := uuid.New()

	cases := []struct {
		name          string
		canalID       uuid.UUID
		destinatarioID uuid.UUID
		wantErr       error
	}{
		{"CanalID nil", uuid.Nil, validDestinatario, model.ErrInvitacionCanalIDNil},
		{"DestinatarioID nil", validCanal, uuid.Nil, model.ErrInvitacionDestinatarioIDNil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := fact.Create(c.canalID, c.destinatarioID)
			if err != c.wantErr {
				t.Errorf("esperado error %v, obtuvo %v", c.wantErr, err)
			}
		})
	}
}

func TestInvitacionCanalFactory_CreateWithStatus_Errors(t *testing.T) {
	fact := factory.NewInvitacionCanalFactory()
	validCanal := uuid.New()
	validDestinatario := uuid.New()
	validEstado := model.InvitacionPendiente

	cases := []struct {
		name          string
		canalID       uuid.UUID
		destinatarioID uuid.UUID
		estado        model.EstadoInvitacion
		wantErr       error
	}{
		{"CanalID nil", uuid.Nil, validDestinatario, validEstado, model.ErrInvitacionCanalIDNil},
		{"DestinatarioID nil", validCanal, uuid.Nil, validEstado, model.ErrInvitacionDestinatarioIDNil},
		{"Estado inválido", validCanal, validDestinatario, model.EstadoInvitacion("estado_inválido"), model.ErrInvitacionEstadoInvalido},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := fact.CreateWithStatus(c.canalID, c.destinatarioID, c.estado)
			if err != c.wantErr {
				t.Errorf("esperado error %v, obtuvo %v", c.wantErr, err)
			}
		})
	}
}
