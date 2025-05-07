package model

import (
    "testing"
    "time"

    "github.com/google/uuid"
)

func TestNewLogEntry_Success_WithUsuario(t *testing.T) {
    id := uuid.New()
    userID := uuid.New()
    now := time.Now().UTC().Truncate(time.Second)

    le, err := NewLogEntry(id, EventoLogin, "Usuario inició sesión", now, userID)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if le.ID() != id {
        t.Errorf("ID: esperado %v, obtuvo %v", id, le.ID())
    }
    if le.TipoEvento() != EventoLogin {
        t.Errorf("TipoEvento: esperado %v, obtuvo %v", EventoLogin, le.TipoEvento())
    }
    if le.Detalle() != "Usuario inició sesión" {
        t.Errorf("Detalle: esperado %q, obtuvo %q", "Usuario inició sesión", le.Detalle())
    }
    if !le.Timestamp().Equal(now) {
        t.Errorf("Timestamp: esperado %v, obtuvo %v", now, le.Timestamp())
    }
    if le.UsuarioID() != userID {
        t.Errorf("UsuarioID: esperado %v, obtuvo %v", userID, le.UsuarioID())
    }
}

func TestNewLogEntry_Success_WithoutUsuario(t *testing.T) {
    id := uuid.New()
    now := time.Now().UTC().Truncate(time.Second)

    le, err := NewLogEntry(id, EventoCanal, "Canal creado", now, uuid.Nil)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if le.UsuarioID() != uuid.Nil {
        t.Errorf("UsuarioID: esperado Nil, obtuvo %v", le.UsuarioID())
    }
}

func TestNewLogEntry_Errors(t *testing.T) {
    baseID := uuid.New()
    now := time.Now().UTC()

    cases := []struct {
        name      string
        fn        func() (*LogEntry, error)
        wantErr   error
    }{
        {"ID inválido", func() (*LogEntry, error) {
            return NewLogEntry(uuid.Nil, EventoLogin, "x", now, uuid.Nil)
        }, ErrLogEntryIDNil},

        {"Evento inválido", func() (*LogEntry, error) {
            return NewLogEntry(baseID, EventoTipo("XYZ"), "x", now, uuid.Nil)
        }, ErrEventoTipoInvalido},

        {"Detalle vacío", func() (*LogEntry, error) {
            return NewLogEntry(baseID, EventoMensaje, "", now, uuid.Nil)
        }, ErrDetalleVacio},

        {"Timestamp cero", func() (*LogEntry, error) {
            return NewLogEntry(baseID, EventoArchivo, "x", time.Time{}, uuid.Nil)
        }, ErrTimestampLogZero},
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            _, err := c.fn()
            if err != c.wantErr {
                t.Errorf("%s: esperado %v, obtuvo %v", c.name, c.wantErr, err)
            }
        })
    }
}
