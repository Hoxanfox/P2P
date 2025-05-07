package factory

import (
    "time"

    "github.com/google/uuid"
    "model"
)

// IMessageFactory define los métodos para crear MensajeServidor.
type IMessageFactory interface {
    // CreateDirect genera un mensaje directo 1:1.
    CreateDirect(remitenteID, destinoUsuarioID uuid.UUID, contenido string, archivoID uuid.UUID) (*model.MensajeServidor, error)
    // CreateChannel genera un mensaje de canal.
    CreateChannel(remitenteID, canalID uuid.UUID, contenido string, archivoID uuid.UUID) (*model.MensajeServidor, error)
    // CreatePrivateChat genera un mensaje para un chat privado.
    CreatePrivateChat(remitenteID, chatPrivadoID uuid.UUID, contenido string, archivoID uuid.UUID) (*model.MensajeServidor, error)
}

// mensajeFactory es la implementación de IMessageFactory.
type mensajeFactory struct{}

// NewMensajeFactory devuelve una instancia de IMessageFactory.
func NewMensajeFactory() IMessageFactory {
    return &mensajeFactory{}
}

// CreateDirect crea un MensajeServidor directo, generando ID y timestamp.
func (f *mensajeFactory) CreateDirect(
    remitenteID, destinoUsuarioID uuid.UUID,
    contenido string,
    archivoID uuid.UUID,
) (*model.MensajeServidor, error) {
    id := uuid.New()
    ahora := time.Now().UTC()
    return model.NewMensajeDirecto(
        id,
        remitenteID,
        destinoUsuarioID,
        contenido,
        ahora,
        archivoID,
    )
}

// CreateChannel crea un MensajeServidor de canal, generando ID y timestamp.
func (f *mensajeFactory) CreateChannel(
    remitenteID, canalID uuid.UUID,
    contenido string,
    archivoID uuid.UUID,
) (*model.MensajeServidor, error) {
    id := uuid.New()
    ahora := time.Now().UTC()
    return model.NewMensajeCanal(
        id,
        remitenteID,
        canalID,
        contenido,
        ahora,
        archivoID,
    )
}

// CreatePrivateChat crea un MensajeServidor para un chat privado, generando ID y timestamp.
func (f *mensajeFactory) CreatePrivateChat(
    remitenteID, chatPrivadoID uuid.UUID,
    contenido string,
    archivoID uuid.UUID,
) (*model.MensajeServidor, error) {
    id := uuid.New()
    ahora := time.Now().UTC()
    return model.NewMensajeChatPrivado(
        id,
        remitenteID,
        chatPrivadoID,
        contenido,
        ahora,
        archivoID,
    )
}
