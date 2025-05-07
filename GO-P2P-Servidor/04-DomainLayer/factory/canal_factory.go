package factory

import (
    "github.com/google/uuid"
    "model"
)

// IChannelFactory define la creación de CanalServidor.
type IChannelFactory interface {
    // Create genera un nuevo CanalServidor con un ID aleatorio.
    Create(nombre, descripcion string, tipo model.CanalTipo) (*model.CanalServidor, error)
}

// CanalFactory es la implementación de IChannelFactory.
type CanalFactory struct{}

// NewCanalFactory retorna una instancia de IChannelFactory.
func NewCanalFactory() IChannelFactory {
    return &CanalFactory{}
}

// Create genera un UUID, valida los parámetros y construye el CanalServidor.
func (f *CanalFactory) Create(nombre, descripcion string, tipo model.CanalTipo) (*model.CanalServidor, error) {
    id := uuid.New()
    return model.NewCanalServidor(id, nombre, descripcion, tipo)
}
