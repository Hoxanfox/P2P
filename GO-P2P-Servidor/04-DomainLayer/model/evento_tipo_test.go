package model

import "testing"

func TestEventoTipo_Valid(t *testing.T) {
    validos := []EventoTipo{
        EventoLogin,
        EventoMensaje,
        EventoArchivo,
        EventoCanal,
    }
    for _, e := range validos {
        if !e.Valid() {
            t.Errorf("EventoTipo.Valid(): se esperaba %q válido", e)
        }
    }
}

func TestEventoTipo_Invalid(t *testing.T) {
    invalidos := []EventoTipo{
        "",              // vacío
        "LOGIN ",        // con espacio
        "logout",        // distinto caso
        EventoTipo("XYZ"),
    }
    for _, e := range invalidos {
        if e.Valid() {
            t.Errorf("EventoTipo.Valid(): se esperaba %q inválido", e)
        }
    }
}
