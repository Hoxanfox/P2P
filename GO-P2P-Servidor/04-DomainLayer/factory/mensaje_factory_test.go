package factory_test

import (
    "testing"
    "time"

    "github.com/google/uuid"
    "factory"
    "model"
)

func TestNewMensajeFactory_NotNil(t *testing.T) {
    mf := factory.NewMensajeFactory()
    if mf == nil {
        t.Fatal("NewMensajeFactory devolvió nil")
    }
}

func TestCreateDirect_Success(t *testing.T) {
    mf := factory.NewMensajeFactory()
    remitente := uuid.New()
    receptor := uuid.New()
    contenido := "Hola, ¿cómo estás?"
    archivo := uuid.Nil // sin adjunto

    msg, err := mf.CreateDirect(remitente, receptor, contenido, archivo)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if msg.ID() == uuid.Nil {
        t.Error("esperaba ID no-nil")
    }
    if msg.RemitenteID() != remitente {
        t.Errorf("RemitenteID esperado %v, obtuvo %v", remitente, msg.RemitenteID())
    }
    if msg.DestinoUsuarioID() != receptor {
        t.Errorf("DestinoUsuarioID esperado %v, obtuvo %v", receptor, msg.DestinoUsuarioID())
    }
    if msg.CanalID() != uuid.Nil {
        t.Errorf("CanalID esperado Nil, obtuvo %v", msg.CanalID())
    }
    if msg.Contenido() != contenido {
        t.Errorf("Contenido esperado %q, obtuvo %q", contenido, msg.Contenido())
    }
    if msg.ArchivoID() != archivo {
        t.Errorf("ArchivoID esperado %v, obtuvo %v", archivo, msg.ArchivoID())
    }
    if msg.Timestamp().IsZero() {
        t.Error("Timestamp no debe ser cero")
    }
    if d := time.Since(msg.Timestamp()); d < 0 || d > time.Second {
        t.Errorf("Timestamp inesperado, dif=%v", d)
    }
}

func TestCreateChannel_Success(t *testing.T) {
    mf := factory.NewMensajeFactory()
    remitente := uuid.New()
    canal := uuid.New()
    contenido := "Mensaje de canal"
    archivo := uuid.Nil

    msg, err := mf.CreateChannel(remitente, canal, contenido, archivo)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if msg.ID() == uuid.Nil {
        t.Error("esperaba ID no-nil")
    }
    if msg.RemitenteID() != remitente {
        t.Errorf("RemitenteID esperado %v, obtuvo %v", remitente, msg.RemitenteID())
    }
    if msg.CanalID() != canal {
        t.Errorf("CanalID esperado %v, obtuvo %v", canal, msg.CanalID())
    }
    if msg.DestinoUsuarioID() != uuid.Nil {
        t.Errorf("DestinoUsuarioID esperado Nil, obtuvo %v", msg.DestinoUsuarioID())
    }
    if msg.Contenido() != contenido {
        t.Errorf("Contenido esperado %q, obtuvo %q", contenido, msg.Contenido())
    }
    if msg.Timestamp().IsZero() {
        t.Error("Timestamp no debe ser cero")
    }
}

func TestCreatePrivateChat_Success(t *testing.T) {
    mf := factory.NewMensajeFactory()
    remitente := uuid.New()
    chatPrivado := uuid.New()
    contenido := "Mensaje en chat privado"
    archivo := uuid.Nil // sin adjunto

    msg, err := mf.CreatePrivateChat(remitente, chatPrivado, contenido, archivo)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if msg.ID() == uuid.Nil {
        t.Error("esperaba ID no-nil")
    }
    if msg.RemitenteID() != remitente {
        t.Errorf("RemitenteID esperado %v, obtuvo %v", remitente, msg.RemitenteID())
    }
    if msg.DestinoUsuarioID() != uuid.Nil {
        t.Errorf("DestinoUsuarioID esperado Nil, obtuvo %v", msg.DestinoUsuarioID())
    }
    if msg.CanalID() != uuid.Nil {
        t.Errorf("CanalID esperado Nil, obtuvo %v", msg.CanalID())
    }
    if msg.ChatPrivadoID() != chatPrivado {
        t.Errorf("ChatPrivadoID esperado %v, obtuvo %v", chatPrivado, msg.ChatPrivadoID())
    }
    if msg.Contenido() != contenido {
        t.Errorf("Contenido esperado %q, obtuvo %q", contenido, msg.Contenido())
    }
    if msg.ArchivoID() != archivo {
        t.Errorf("ArchivoID esperado %v, obtuvo %v", archivo, msg.ArchivoID())
    }
    if msg.Timestamp().IsZero() {
        t.Error("Timestamp no debe ser cero")
    }
    if d := time.Since(msg.Timestamp()); d < 0 || d > time.Second {
        t.Errorf("Timestamp fuera de rango, diff=%v", d)
    }
}

func TestCreateDirect_Errors(t *testing.T) {
    mf := factory.NewMensajeFactory()
    valid := uuid.New()

    cases := []struct {
        name       string
        remitente  uuid.UUID
        destino    uuid.UUID
        canal      uuid.UUID
        contenido  string
        wantErr    error
    }{
        {"sin remitente", uuid.Nil, uuid.New(), uuid.Nil, "x", model.ErrRemitenteIDNil},
        {"sin destino ni canal", valid, uuid.Nil, uuid.Nil, "x", model.ErrSinDestinoNiCanal},
        // El caso de destino y canal no se puede probar directamente en estos tests porque
        // los métodos CreateDirect y CreateChannel ya están separados y cada uno solo acepta un destino
        {"contenido vacío", valid, uuid.New(), uuid.Nil, "", model.ErrContenidoVacio},
    }
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            _, err := mf.CreateDirect(c.remitente, c.destino, c.contenido, uuid.Nil)
            if err != c.wantErr {
                t.Errorf("%s: esperado %v, obtuvo %v", c.name, c.wantErr, err)
            }
        })
    }
}

func TestCreateChannel_Errors(t *testing.T) {
    mf := factory.NewMensajeFactory()
    valid := uuid.New()

    cases := []struct {
        name      string
        remitente uuid.UUID
        canal     uuid.UUID
        contenido string
        wantErr   error
    }{
        {"sin remitente", uuid.Nil, uuid.New(), "x", model.ErrRemitenteIDNil},
        {"sin destino ni canal", valid, uuid.Nil, "x", model.ErrSinDestinoNiCanal},
        // El caso de destino y canal no se puede probar directamente en estos tests porque
        // los métodos CreateDirect y CreateChannel ya están separados y cada uno solo acepta un destino
        {"contenido vacío", valid, uuid.New(), "", model.ErrContenidoVacio},
    }
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            _, err := mf.CreateChannel(c.remitente, c.canal, c.contenido, uuid.Nil)
            if err != c.wantErr {
                t.Errorf("%s: esperado %v, obtuvo %v", c.name, c.wantErr, err)
            }
        })
    }
}
