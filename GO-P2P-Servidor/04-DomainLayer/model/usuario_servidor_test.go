package model

import (
    "testing"
    "time"

    "github.com/google/uuid"
)

func TestNewUsuarioServidor_Success(t *testing.T) {
    id := uuid.New()
    now := time.Now().UTC().Truncate(time.Second)

    u, err := NewUsuarioServidor(
        id,
        "alice",
        "alice@example.com",
        "hashed_pw",
        "http://example.com/avatar.png",
        "192.168.0.10",
        now,
    )
    if err != nil {
        t.Fatalf("esperaba sin error, obtuve %v", err)
    }
    if u.ID() != id {
        t.Errorf("ID: esperado %v, obtuvo %v", id, u.ID())
    }
    if u.NombreUsuario() != "alice" {
        t.Errorf("NombreUsuario: esperado %q, obtuvo %q", "alice", u.NombreUsuario())
    }
    if u.Email() != "alice@example.com" {
        t.Errorf("Email: esperado %q, obtuvo %q", "alice@example.com", u.Email())
    }
    if u.ContrasenaHasheada() != "hashed_pw" {
        t.Errorf("ContrasenaHasheada: esperado %q, obtuvo %q", "hashed_pw", u.ContrasenaHasheada())
    }
    if u.FotoURL() != "http://example.com/avatar.png" {
        t.Errorf("FotoURL: esperado %q, obtuvo %q", "http://example.com/avatar.png", u.FotoURL())
    }
    if u.IPRegistrada() != "192.168.0.10" {
        t.Errorf("IPRegistrada: esperado %q, obtuvo %q", "192.168.0.10", u.IPRegistrada())
    }
    if !u.FechaRegistro().Equal(now) {
        t.Errorf("FechaRegistro: esperado %v, obtuvo %v", now, u.FechaRegistro())
    }
}

func TestNewUsuarioServidor_Errors(t *testing.T) {
    now := time.Now().UTC().Truncate(time.Second)
    cases := []struct {
        name          string
        id            uuid.UUID
        nombreUsuario string
        email         string
        hash          string
        fechaRegistro time.Time
        wantErr       error
    }{
        {"ID vacío", uuid.Nil, "alice", "a@b.com", "h", now, ErrIDNil},
        {"Nombre vacío", uuid.New(), "", "a@b.com", "h", now, ErrNombreVacio},
        {"Email inválido", uuid.New(), "alice", "no-email", "h", now, ErrEmailInvalido},
        {"Hash vacío", uuid.New(), "alice", "a@b.com", "", now, ErrHashVacio},
        {"Fecha cero", uuid.New(), "alice", "a@b.com", "h", time.Time{}, ErrFechaRegistroZero},
    }
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            _, err := NewUsuarioServidor(
                c.id,
                c.nombreUsuario,
                c.email,
                c.hash,
                "", "", 
                c.fechaRegistro,
            )
            if err != c.wantErr {
                t.Errorf("%s: esperado error %v, obtuvo %v", c.name, c.wantErr, err)
            }
        })
    }
}
