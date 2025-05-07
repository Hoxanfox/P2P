package model

import "testing"

func TestNodoEstado_Valid(t *testing.T) {
    validos := []NodoEstado{
        NodoConectado,
        NodoDesconectado,
    }
    for _, n := range validos {
        if !n.Valid() {
            t.Errorf("NodoEstado.Valid(): se esperaba %q válido", n)
        }
    }
}

func TestNodoEstado_Invalid(t *testing.T) {
    invalidos := []NodoEstado{
        "",                  // vacío
        "conectado",         // minúsculas
        "CONECTA D O",       // espacios intermedios
        NodoEstado("UNKNOWN"), // valor fuera de enum
    }
    for _, n := range invalidos {
        if n.Valid() {
            t.Errorf("NodoEstado.Valid(): se esperaba %q inválido", n)
        }
    }
}
