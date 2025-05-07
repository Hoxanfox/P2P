package model_test

import (
    "testing"
    "time"

    "github.com/google/uuid"
    "model"
)

func TestNewReplicaEvent_Success(t *testing.T) {
    id := uuid.New()
    entidadID := uuid.New()
    origenID := uuid.New()
    now := time.Now().UTC().Truncate(time.Second)

    re, err := model.NewReplicaEvent(id, "Usuario", entidadID, now, origenID)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if re.ID() != id {
        t.Errorf("ID: esperado %v, obtuvo %v", id, re.ID())
    }
    if re.EntidadTipo() != "Usuario" {
        t.Errorf("EntidadTipo: esperado %q, obtuvo %q", "Usuario", re.EntidadTipo())
    }
    if re.EntidadID() != entidadID {
        t.Errorf("EntidadID: esperado %v, obtuvo %v", entidadID, re.EntidadID())
    }
    if !re.EventoAt().Equal(now) {
        t.Errorf("EventoAt: esperado %v, obtuvo %v", now, re.EventoAt())
    }
    if re.OrigenNodoID() != origenID {
        t.Errorf("OrigenNodoID: esperado %v, obtuvo %v", origenID, re.OrigenNodoID())
    }
}

func TestNewReplicaEvent_Errors(t *testing.T) {
    baseID := uuid.New()
    entidadID := uuid.New()
    origenID := uuid.New()
    now := time.Now().UTC()

    cases := []struct {
        name      string
        fn        func() (*model.ReplicaEvent, error)
        wantErr   error
    }{
        {"ID inválido", func() (*model.ReplicaEvent, error) {
            return model.NewReplicaEvent(uuid.Nil, "X", entidadID, now, origenID)
        }, model.ErrReplicaEventIDNil},
        {"EntidadTipo vacío", func() (*model.ReplicaEvent, error) {
            return model.NewReplicaEvent(baseID, "", entidadID, now, origenID)
        }, model.ErrEntidadTipoVacio},
        {"EntidadID inválido", func() (*model.ReplicaEvent, error) {
            return model.NewReplicaEvent(baseID, "X", uuid.Nil, now, origenID)
        }, model.ErrEntidadIDNil},
        {"EventoAt cero", func() (*model.ReplicaEvent, error) {
            return model.NewReplicaEvent(baseID, "X", entidadID, time.Time{}, origenID)
        }, model.ErrReplicaEventAtZero},
        {"OrigenNodoID inválido", func() (*model.ReplicaEvent, error) {
            return model.NewReplicaEvent(baseID, "X", entidadID, now, uuid.Nil)
        }, model.ErrOrigenNodoIDNil},
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
