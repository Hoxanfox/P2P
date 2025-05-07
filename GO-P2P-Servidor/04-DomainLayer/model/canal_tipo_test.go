package model

import "testing"

func TestCanalTipo_Valid(t *testing.T) {
    validos := []CanalTipo{
        CanalPublico,
        CanalPrivado,
    }
    for _, c := range validos {
        if !c.Valid() {
            t.Errorf("CanalTipo.Valid(): se esperaba %q válido", c)
        }
    }
}

func TestCanalTipo_Invalid(t *testing.T) {
    invalidos := []CanalTipo{
        "",           // vacío
        "publico",    // minúsculas
        "PRIV",       // acortado
        "PUBLICO ",   // espacio
        CanalTipo("XYZ"),
    }
    for _, c := range invalidos {
        if c.Valid() {
            t.Errorf("CanalTipo.Valid(): se esperaba %q inválido", c)
        }
    }
}
