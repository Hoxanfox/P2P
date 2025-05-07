package factory

import (
    "time"

    "github.com/google/uuid"
    "model"
)

// UsuarioFactory define la abstracción para crear usuarios con toda la lógica de generación de ID y timestamp.
type UsuarioFactory interface {
    // Create construye un UsuarioServidor nuevo generando el ID y la fecha de registro automáticamente.
    Create(
        nombreUsuario, email, contrasenaHasheada, fotoURL, ipRegistrada string,
    ) (*model.UsuarioServidor, error)
    
    // CreateConnected construye un UsuarioServidor conectado.
    CreateConnected(
        nombreUsuario, email, contrasenaHasheada, fotoURL, ipRegistrada string,
    ) (*model.UsuarioServidor, error)
}

type usuarioFactory struct{}

// NewUsuarioFactory devuelve la implementación por defecto de UsuarioFactory.
func NewUsuarioFactory() UsuarioFactory {
    return &usuarioFactory{}
}

// Create implementa UsuarioFactory: genera el UUID y la marca de tiempo, y delega en el constructor de dominio.
func (f *usuarioFactory) Create(
    nombreUsuario, email, contrasenaHasheada, fotoURL, ipRegistrada string,
) (*model.UsuarioServidor, error) {
    id := uuid.New()
    ahora := time.Now().UTC()
    return model.NewUsuarioServidor(
        id,
        nombreUsuario,
        email,
        contrasenaHasheada,
        fotoURL,
        ipRegistrada,
        ahora,
    )
}

// CreateConnected implementa UsuarioFactory: crea un usuario y lo marca como conectado.
func (f *usuarioFactory) CreateConnected(
    nombreUsuario, email, contrasenaHasheada, fotoURL, ipRegistrada string,
) (*model.UsuarioServidor, error) {
    usuario, err := f.Create(nombreUsuario, email, contrasenaHasheada, fotoURL, ipRegistrada)
    if err != nil {
        return nil, err
    }
    
    // Marcar el usuario como conectado
    usuario.SetConnected(true)
    return usuario, nil
}
