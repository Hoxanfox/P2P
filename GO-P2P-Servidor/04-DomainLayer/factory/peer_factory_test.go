package factory_test

import (
    "testing"

    "github.com/google/uuid"
    "factory"
    "model"
)

func TestNewPeerFactory_NotNil(t *testing.T) {
    pf := factory.NewPeerFactory()
    if pf == nil {
        t.Fatal("NewPeerFactory devolvió nil")
    }
}

func TestCreate_DefaultStateConnected(t *testing.T) {
    pf := factory.NewPeerFactory()
    dir := "127.0.0.1:8000"

    peer, err := pf.Create(dir)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if peer.IDNodo() == uuid.Nil {
        t.Error("esperaba IDNodo no-nil")
    }
    if peer.Direccion() != dir {
        t.Errorf("Direccion esperada %q, obtuvo %q", dir, peer.Direccion())
    }
    if peer.Estado() != model.NodoConectado {
        t.Errorf("Estado esperado NodoConectado, obtuvo %v", peer.Estado())
    }
}

func TestCreateWithState_CustomState(t *testing.T) {
    pf := factory.NewPeerFactory()
    dir := "10.0.0.5:11000"
    wantState := model.NodoDesconectado

    peer, err := pf.CreateWithState(dir, wantState)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if peer.Direccion() != dir {
        t.Errorf("Direccion esperada %q, obtuvo %q", dir, peer.Direccion())
    }
    if peer.Estado() != wantState {
        t.Errorf("Estado esperado %v, obtuvo %v", wantState, peer.Estado())
    }
}

func TestCreate_Errors(t *testing.T) {
    pf := factory.NewPeerFactory()
    cases := []struct {
        name    string
        dir     string
        state   model.NodoEstado
        wantErr error
    }{
        {"dirección vacía", "", model.NodoConectado, model.ErrPeerDireccionVacia},
        {"formato inválido", "127.0.0.1", model.NodoConectado, model.ErrPeerDireccionFormat},
        {"estado inválido", "127.0.0.1:9000", model.NodoEstado("INVALIDO"), model.ErrPeerEstadoInvalido},
    }
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            _, err := pf.CreateWithState(c.dir, c.state)
            if err != c.wantErr {
                t.Errorf("%s: esperado %v, obtuvo %v", c.name, c.wantErr, err)
            }
        })
    }
}
