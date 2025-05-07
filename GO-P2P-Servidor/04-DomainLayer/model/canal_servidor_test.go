package model

import (
    "testing"

    "github.com/google/uuid"
)

func TestNewCanalServidor_Success(t *testing.T) {
    id := uuid.New()
    c, err := NewCanalServidor(id, "General", "Canal público para todos", CanalPublico)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if c.ID() != id {
        t.Errorf("ID: esperado %v, obtuvo %v", id, c.ID())
    }
    if c.Nombre() != "General" {
        t.Errorf("Nombre: esperado %q, obtuvo %q", "General", c.Nombre())
    }
    if c.Descripcion() != "Canal público para todos" {
        t.Errorf("Descripcion: esperado %q, obtuvo %q", "Canal público para todos", c.Descripcion())
    }
    if c.Tipo() != CanalPublico {
        t.Errorf("Tipo: esperado %q, obtuvo %q", CanalPublico, c.Tipo())
    }
}

func TestNewCanalServidor_Errors(t *testing.T) {
    validID := uuid.New()
    cases := []struct {
        name        string
        id          uuid.UUID
        nombre      string
        tipo        CanalTipo
        wantErr     error
    }{
        {"ID vacío", uuid.Nil, "General", CanalPublico, ErrCanalIDNil},
        {"Nombre vacío", validID, "", CanalPublico, ErrCanalNombreVacio},
        {"Tipo inválido", validID, "General", CanalTipo("XYZ"), ErrCanalTipoInvalido},
    }
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            _, err := NewCanalServidor(c.id, c.nombre, "desc", c.tipo)
            if err != c.wantErr {
                t.Errorf("%s: esperado error %v, obtuvo %v", c.name, c.wantErr, err)
            }
        })
    }
}
