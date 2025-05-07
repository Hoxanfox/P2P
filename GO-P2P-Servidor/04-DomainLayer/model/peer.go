package model

import (
    "errors"
    "net"

    "github.com/google/uuid"
)

// Errores de validación para Peer
var (
    ErrPeerIDNil           = errors.New("idNodo inválido")
    ErrPeerDireccionVacia  = errors.New("dirección de peer vacía")
    ErrPeerDireccionFormat = errors.New("dirección debe tener formato host:puerto")
    ErrPeerEstadoInvalido  = errors.New("estado de peer inválido")
)

// Peer representa un nodo en la red P2P.
type Peer struct {
    idNodo    uuid.UUID
    direccion string
    estado    NodoEstado
}

// NewPeer crea un nuevo Peer validando sus invariantes:
// - idNodo no puede ser Nil
// - dirección no puede estar vacía y debe cumplir host:port
// - estado debe ser una constante válida de NodoEstado
func NewPeer(
    idNodo uuid.UUID,
    direccion string,
    estado NodoEstado,
) (*Peer, error) {
    if idNodo == uuid.Nil {
        return nil, ErrPeerIDNil
    }
    if direccion == "" {
        return nil, ErrPeerDireccionVacia
    }
    host, port, err := net.SplitHostPort(direccion)
    if err != nil || host == "" || port == "" {
        return nil, ErrPeerDireccionFormat
    }
    
    // Verificar que el puerto sea numérico
    _, err = net.LookupPort("tcp", port)
    if err != nil {
        return nil, ErrPeerDireccionFormat
    }
    if !estado.Valid() {
        return nil, ErrPeerEstadoInvalido
    }
    return &Peer{
        idNodo:    idNodo,
        direccion: direccion,
        estado:    estado,
    }, nil
}

// IDNodo retorna el UUID del peer.
func (p *Peer) IDNodo() uuid.UUID {
    return p.idNodo
}

// Direccion retorna la dirección (host:puerto) del peer.
func (p *Peer) Direccion() string {
    return p.direccion
}

// Estado retorna el estado de conexión del peer.
func (p *Peer) Estado() NodoEstado {
    return p.estado
}
