package model

import (
    "errors"
    "net/mail"
    "time"

    "github.com/google/uuid"
)

// Errores de validación para UsuarioServidor
var (
    ErrIDNil             = errors.New("id de usuario inválido")
    ErrNombreVacio       = errors.New("nombre de usuario vacío")
    ErrEmailInvalido     = errors.New("email con formato inválido")
    ErrHashVacio         = errors.New("contraseña hasheada vacía")
    ErrFechaRegistroZero = errors.New("fecha de registro no puede ser cero")
)

// UsuarioServidor representa la entidad Usuario del dominio del servidor.
type UsuarioServidor struct {
    id                 uuid.UUID
    nombreUsuario      string
    email              string
    contrasenaHasheada string
    fotoURL            string
    ipRegistrada       string
    fechaRegistro      time.Time
    isConnected        bool
}

// NewUsuarioServidor construye un nuevo UsuarioServidor garantizando invariantes.
func NewUsuarioServidor(
    id uuid.UUID,
    nombreUsuario, email, contrasenaHasheada, fotoURL, ipRegistrada string,
    fechaRegistro time.Time,
) (*UsuarioServidor, error) {
    if id == uuid.Nil {
        return nil, ErrIDNil
    }
    if nombreUsuario == "" {
        return nil, ErrNombreVacio
    }
    if _, err := mail.ParseAddress(email); err != nil {
        return nil, ErrEmailInvalido
    }
    if contrasenaHasheada == "" {
        return nil, ErrHashVacio
    }
    if fechaRegistro.IsZero() {
        return nil, ErrFechaRegistroZero
    }
    return &UsuarioServidor{
        id:                 id,
        nombreUsuario:      nombreUsuario,
        email:              email,
        contrasenaHasheada: contrasenaHasheada,
        fotoURL:            fotoURL,
        ipRegistrada:       ipRegistrada,
        fechaRegistro:      fechaRegistro,
        isConnected:        false, // Por defecto, al crear un usuario está desconectado
    }, nil
}

// Getters (no setters para preservar invariantes)
func (u *UsuarioServidor) ID() uuid.UUID              { return u.id }
func (u *UsuarioServidor) NombreUsuario() string      { return u.nombreUsuario }
func (u *UsuarioServidor) Email() string              { return u.email }
func (u *UsuarioServidor) ContrasenaHasheada() string { return u.contrasenaHasheada }
func (u *UsuarioServidor) FotoURL() string            { return u.fotoURL }
func (u *UsuarioServidor) IPRegistrada() string       { return u.ipRegistrada }
func (u *UsuarioServidor) FechaRegistro() time.Time   { return u.fechaRegistro }
func (u *UsuarioServidor) IsConnected() bool          { return u.isConnected }

// SetConnected establece el estado de conexión del usuario
func (u *UsuarioServidor) SetConnected(connected bool) {
    u.isConnected = connected
}
