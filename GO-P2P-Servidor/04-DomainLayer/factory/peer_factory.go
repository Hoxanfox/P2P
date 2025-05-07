package factory

import (
    "github.com/google/uuid"
    "model"
)

// IPeerFactory define la creación de Peers P2P.
type IPeerFactory interface {
    // Create genera un Peer con estado inicial CONECTADO.
    Create(direccion string) (*model.Peer, error)
    // CreateWithState genera un Peer con el estado dado.
    CreateWithState(direccion string, estado model.NodoEstado) (*model.Peer, error)
}

type peerFactory struct{}

// NewPeerFactory devuelve la implementación por defecto de IPeerFactory.
func NewPeerFactory() IPeerFactory {
    return &peerFactory{}
}

// Create genera un Peer en estado model.NodoEstadoCONECTADO.
func (f *peerFactory) Create(direccion string) (*model.Peer, error) {
    return f.CreateWithState(direccion, model.NodoConectado)
}

// CreateWithState genera un Peer con el estado especificado.
func (f *peerFactory) CreateWithState(direccion string, estado model.NodoEstado) (*model.Peer, error) {
    id := uuid.New()
    return model.NewPeer(id, direccion, estado)
}
