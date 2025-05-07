package model

import (
    "errors"

    "github.com/google/uuid"
)

// Errores de validación para CanalServidor
var (
    ErrCanalIDNil         = errors.New("id de canal inválido")
    ErrCanalNombreVacio   = errors.New("nombre de canal vacío")
    ErrCanalTipoInvalido  = errors.New("tipo de canal inválido")
)

// CanalServidor representa un canal de comunicación en el servidor.
type CanalServidor struct {
    id          uuid.UUID
    nombre      string
    descripcion string
    tipo        CanalTipo
}

// NewCanalServidor crea un CanalServidor validando sus invariantes.
func NewCanalServidor(
    id uuid.UUID,
    nombre, descripcion string,
    tipo CanalTipo,
) (*CanalServidor, error) {
    if id == uuid.Nil {
        return nil, ErrCanalIDNil
    }
    if nombre == "" {
        return nil, ErrCanalNombreVacio
    }
    if !tipo.Valid() {
        return nil, ErrCanalTipoInvalido
    }
    return &CanalServidor{
        id:          id,
        nombre:      nombre,
        descripcion: descripcion,
        tipo:        tipo,
    }, nil
}

// Getters
func (c *CanalServidor) ID() uuid.UUID  { return c.id }
func (c *CanalServidor) Nombre() string { return c.nombre }
func (c *CanalServidor) Descripcion() string { return c.descripcion }
func (c *CanalServidor) Tipo() CanalTipo { return c.tipo }
