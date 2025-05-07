package model

import (
    "testing"
    "time"

    "github.com/google/uuid"
)

func TestNewHeartbeatLog_Success(t *testing.T) {
    id := uuid.New()
    nodo := uuid.New()
    now := time.Now().UTC().Truncate(time.Second)
    later := now.Add(500 * time.Millisecond)

    hb, err := NewHeartbeatLog(id, nodo, now, later)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if hb.ID() != id {
        t.Errorf("ID: esperado %v, obtuvo %v", id, hb.ID())
    }
    if hb.NodoID() != nodo {
        t.Errorf("NodoID: esperado %v, obtuvo %v", nodo, hb.NodoID())
    }
    if !hb.EnviadoAt().Equal(now) {
        t.Errorf("EnviadoAt: esperado %v, obtuvo %v", now, hb.EnviadoAt())
    }
    if !hb.RecibidoAt().Equal(later) {
        t.Errorf("RecibidoAt: esperado %v, obtuvo %v", later, hb.RecibidoAt())
    }
}

func TestNewHeartbeatLog_Errors(t *testing.T) {
    baseID := uuid.New()
    nodo := uuid.New()
    now := time.Now().UTC()

    cases := []struct {
        name    string
        fn      func() (*HeartbeatLog, error)
        wantErr error
    }{
        {"ID inválido", func() (*HeartbeatLog, error) {
            return NewHeartbeatLog(uuid.Nil, nodo, now, now.Add(time.Second))
        }, ErrHeartbeatLogIDNil},
        {"NodoID inválido", func() (*HeartbeatLog, error) {
            return NewHeartbeatLog(baseID, uuid.Nil, now, now.Add(time.Second))
        }, ErrHeartbeatNodoIDNil},
        {"EnviadoAt cero", func() (*HeartbeatLog, error) {
            return NewHeartbeatLog(baseID, nodo, time.Time{}, now)
        }, ErrHeartbeatEnviadoZero},
        {"RecibidoAt cero", func() (*HeartbeatLog, error) {
            return NewHeartbeatLog(baseID, nodo, now, time.Time{})
        }, ErrHeartbeatRecibidoZero},
        {"Recibido antes de enviado", func() (*HeartbeatLog, error) {
            return NewHeartbeatLog(baseID, nodo, now, now.Add(-time.Second))
        }, ErrHeartbeatRecibidoAntes},
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
