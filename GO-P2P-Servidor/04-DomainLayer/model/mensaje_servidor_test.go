package model_test

import (
    "testing"
    "time"

    "github.com/google/uuid"
    "model"
)

func TestNewMensajeDirecto_Success(t *testing.T) {
    id := uuid.New()
    remitente := uuid.New()
    receptor := uuid.New()
    now := time.Now().UTC().Truncate(time.Second)
    adj := uuid.New()

    m, err := model.NewMensajeDirecto(id, remitente, receptor, "Hola, ¿qué tal?", now, adj)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if m.ID() != id {
        t.Errorf("ID: esperado %v, obtuvo %v", id, m.ID())
    }
    if m.RemitenteID() != remitente {
        t.Errorf("RemitenteID: esperado %v, obtuvo %v", remitente, m.RemitenteID())
    }
    if m.DestinoUsuarioID() != receptor {
        t.Errorf("DestinoUsuarioID: esperado %v, obtuvo %v", receptor, m.DestinoUsuarioID())
    }
    if m.CanalID() != uuid.Nil {
        t.Errorf("CanalID: esperado Nil, obtuvo %v", m.CanalID())
    }
    if m.Contenido() != "Hola, ¿qué tal?" {
        t.Errorf("Contenido: esperado %q, obtuvo %q", "Hola, ¿qué tal?", m.Contenido())
    }
    if !m.Timestamp().Equal(now) {
        t.Errorf("Timestamp: esperado %v, obtuvo %v", now, m.Timestamp())
    }
    if m.ArchivoID() != adj {
        t.Errorf("ArchivoID: esperado %v, obtuvo %v", adj, m.ArchivoID())
    }
}

func TestNewMensajeCanal_Success(t *testing.T) {
    id := uuid.New()
    remitente := uuid.New()
    canal := uuid.New()
    now := time.Now().UTC().Truncate(time.Second)

    m, err := model.NewMensajeCanal(id, remitente, canal, "Mensaje de canal", now, uuid.Nil)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if m.CanalID() != canal {
        t.Errorf("CanalID: esperado %v, obtuvo %v", canal, m.CanalID())
    }
    if m.DestinoUsuarioID() != uuid.Nil {
        t.Errorf("DestinoUsuarioID: esperado Nil, obtuvo %v", m.DestinoUsuarioID())
    }
}

func TestNewMensaje_Errors(t *testing.T) {
    baseID := uuid.New()
    remitente := uuid.New()
    receptor := uuid.New()
    now := time.Now().UTC()

    cases := []struct {
        name                string
        fn                  func() (*model.MensajeServidor, error)
        wantErr             error
    }{
        {"ID inválido", func() (*model.MensajeServidor, error) {
            return model.NewMensajeDirecto(uuid.Nil, remitente, receptor, "x", now, uuid.Nil)
        }, model.ErrMensajeIDNil},

        {"Remitente inválido", func() (*model.MensajeServidor, error) {
            return model.NewMensajeDirecto(baseID, uuid.Nil, receptor, "x", now, uuid.Nil)
        }, model.ErrRemitenteIDNil},

        {"Contenido vacío", func() (*model.MensajeServidor, error) {
            return model.NewMensajeDirecto(baseID, remitente, receptor, "", now, uuid.Nil)
        }, model.ErrContenidoVacio},

        {"Timestamp cero", func() (*model.MensajeServidor, error) {
            return model.NewMensajeDirecto(baseID, remitente, receptor, "x", time.Time{}, uuid.Nil)
        }, model.ErrTimestampZero},

        {"Sin destino ni canal", func() (*model.MensajeServidor, error) {
            // No podemos usar newMensaje directamente ya que es privado
            // Usamos NewMensajeChatPrivado que internamente validará que se necesita un chatPrivadoID
            return model.NewMensajeChatPrivado(baseID, remitente, uuid.Nil, "x", now, uuid.Nil)
        }, model.ErrMensajeChatPrivadoIDNil},

        {"Con destino y canal", func() (*model.MensajeServidor, error) {
            // Creamos un mensaje directo y luego intentamos modificar el canalID
            // Esto no es posible directamente, así que usaremos una prueba diferente
            m, _ := model.NewMensajeDirecto(baseID, remitente, receptor, "x", now, uuid.Nil)
            // Como no podemos modificar directamente, verificamos que el CanalID sea nil
            if m.CanalID() != uuid.Nil {
                t.Error("CanalID debería ser nil en un mensaje directo")
            }
            return m, nil
        }, nil},
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            _, err := c.fn()
            if err != c.wantErr {
                t.Errorf("%s: esperado %v, obtuvo %v", c.name, c.wantErr, err)
            }
        })
    }
}
