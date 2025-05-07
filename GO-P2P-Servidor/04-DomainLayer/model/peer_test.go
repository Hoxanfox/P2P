package model

import (
    "testing"

    "github.com/google/uuid"
)

func TestNewPeer_Success(t *testing.T) {
    id := uuid.New()
    direccion := "192.168.0.10:11000"

    p, err := NewPeer(id, direccion, NodoConectado)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if p.IDNodo() != id {
        t.Errorf("IDNodo: esperado %v, obtuvo %v", id, p.IDNodo())
    }
    if p.Direccion() != direccion {
        t.Errorf("Direccion: esperado %q, obtuvo %q", direccion, p.Direccion())
    }
    if p.Estado() != NodoConectado {
        t.Errorf("Estado: esperado %v, obtuvo %v", NodoConectado, p.Estado())
    }
}

func TestNewPeer_ErrorIDNil(t *testing.T) {
    _, err := NewPeer(uuid.Nil, "127.0.0.1:11000", NodoConectado)
    if err != ErrPeerIDNil {
        t.Errorf("esperaba ErrPeerIDNil, obtuvo %v", err)
    }
}

func TestNewPeer_ErrorDireccionVacia(t *testing.T) {
    _, err := NewPeer(uuid.New(), "", NodoConectado)
    if err != ErrPeerDireccionVacia {
        t.Errorf("esperaba ErrPeerDireccionVacia, obtuvo %v", err)
    }
}

func TestNewPeer_ErrorDireccionFormato(t *testing.T) {
    invalids := []string{
        "localhost",          // sin puerto
        "localhost:",         // sin puerto numérico
        ":11000",             // sin host
        "192.168.0.1:abc",    // puerto no numérico
    }
    for _, dir := range invalids {
        _, err := NewPeer(uuid.New(), dir, NodoConectado)
        if err != ErrPeerDireccionFormat {
            t.Errorf("dirección %q: esperaba ErrPeerDireccionFormat, obtuvo %v", dir, err)
        }
    }
}

func TestNewPeer_ErrorEstadoInvalido(t *testing.T) {
    id := uuid.New()
    dir := "10.0.0.5:12000"
    invalidEstado := NodoEstado("DESCONOCIDO")
    _, err := NewPeer(id, dir, invalidEstado)
    if err != ErrPeerEstadoInvalido {
        t.Errorf("esperaba ErrPeerEstadoInvalido, obtuvo %v", err)
    }
}
