package model_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"model"
)

func TestNewInvitacionCanal_Success(t *testing.T) {
	id := uuid.New()
	canalID := uuid.New()
	destinatarioID := uuid.New()
	estado := model.InvitacionPendiente
	fechaEnvio := time.Now().UTC().Truncate(time.Second)

	invitacion, err := model.NewInvitacionCanal(id, canalID, destinatarioID, estado, fechaEnvio)
	if err != nil {
		t.Fatalf("esperaba sin error, obtuvo %v", err)
	}

	if invitacion.ID() != id {
		t.Errorf("ID: esperado %v, obtuvo %v", id, invitacion.ID())
	}
	if invitacion.CanalID() != canalID {
		t.Errorf("CanalID: esperado %v, obtuvo %v", canalID, invitacion.CanalID())
	}
	if invitacion.DestinatarioID() != destinatarioID {
		t.Errorf("DestinatarioID: esperado %v, obtuvo %v", destinatarioID, invitacion.DestinatarioID())
	}
	if invitacion.Estado() != estado {
		t.Errorf("Estado: esperado %v, obtuvo %v", estado, invitacion.Estado())
	}
	if !invitacion.FechaEnvio().Equal(fechaEnvio) {
		t.Errorf("FechaEnvio: esperado %v, obtuvo %v", fechaEnvio, invitacion.FechaEnvio())
	}
}

func TestNewInvitacionCanal_Errors(t *testing.T) {
	validID := uuid.New()
	validCanalID := uuid.New()
	validDestinatarioID := uuid.New()
	validEstado := model.InvitacionPendiente
	validFechaEnvio := time.Now().UTC()

	cases := []struct {
		name           string
		id             uuid.UUID
		canalID        uuid.UUID
		destinatarioID uuid.UUID
		estado         model.EstadoInvitacion
		fechaEnvio     time.Time
		wantErr        error
	}{
		{"ID nil", uuid.Nil, validCanalID, validDestinatarioID, validEstado, validFechaEnvio, model.ErrInvitacionIDNil},
		{"CanalID nil", validID, uuid.Nil, validDestinatarioID, validEstado, validFechaEnvio, model.ErrInvitacionCanalIDNil},
		{"DestinatarioID nil", validID, validCanalID, uuid.Nil, validEstado, validFechaEnvio, model.ErrInvitacionDestinatarioIDNil},
		{"Estado inválido", validID, validCanalID, validDestinatarioID, model.EstadoInvitacion("INVALIDO"), validFechaEnvio, model.ErrInvitacionEstadoInvalido},
		{"FechaEnvio cero", validID, validCanalID, validDestinatarioID, validEstado, time.Time{}, model.ErrInvitacionFechaEnvioZero},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := model.NewInvitacionCanal(c.id, c.canalID, c.destinatarioID, c.estado, c.fechaEnvio)
			if err != c.wantErr {
				t.Errorf("esperado error %v, obtuvo %v", c.wantErr, err)
			}
		})
	}
}

func TestInvitacionCanal_CambiarEstado(t *testing.T) {
	id := uuid.New()
	canalID := uuid.New()
	destinatarioID := uuid.New()
	estadoInicial := model.InvitacionPendiente
	fechaEnvio := time.Now().UTC()

	invitacion, _ := model.NewInvitacionCanal(id, canalID, destinatarioID, estadoInicial, fechaEnvio)

	// Caso de éxito: cambiar a aceptada
	err := invitacion.CambiarEstado(model.InvitacionAceptada)
	if err != nil {
		t.Errorf("CambiarEstado a Aceptada: esperaba sin error, obtuvo %v", err)
	}
	if invitacion.Estado() != model.InvitacionAceptada {
		t.Errorf("Estado después de cambio: esperado %v, obtuvo %v", model.InvitacionAceptada, invitacion.Estado())
	}

	// Caso de éxito: cambiar a rechazada
	err = invitacion.CambiarEstado(model.InvitacionRechazada)
	if err != nil {
		t.Errorf("CambiarEstado a Rechazada: esperaba sin error, obtuvo %v", err)
	}
	if invitacion.Estado() != model.InvitacionRechazada {
		t.Errorf("Estado después de cambio: esperado %v, obtuvo %v", model.InvitacionRechazada, invitacion.Estado())
	}

	// Caso de error: estado inválido
	err = invitacion.CambiarEstado(model.EstadoInvitacion("INVALIDO"))
	if err != model.ErrInvitacionEstadoInvalido {
		t.Errorf("CambiarEstado a inválido: esperado error %v, obtuvo %v", model.ErrInvitacionEstadoInvalido, err)
	}
}
