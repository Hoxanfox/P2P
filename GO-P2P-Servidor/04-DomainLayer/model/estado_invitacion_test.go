package model_test

import (
	"testing"

	"model"
)

func TestEstadoInvitacion_Valid(t *testing.T) {
	cases := []struct {
		name  string
		estado model.EstadoInvitacion
		want  bool
	}{
		{"Estado pendiente", model.InvitacionPendiente, true},
		{"Estado aceptada", model.InvitacionAceptada, true},
		{"Estado rechazada", model.InvitacionRechazada, true},
		{"Estado inválido", model.EstadoInvitacion("INVALIDO"), false},
		{"Estado vacío", model.EstadoInvitacion(""), false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := c.estado.Valid(); got != c.want {
				t.Errorf("Valid() = %v, quiere %v", got, c.want)
			}
		})
	}
}
