package factory_test

import (
    "testing"
    "time"

    "github.com/google/uuid"
    
    "factory"
)

func TestReplicaFactory_Create_Success(t *testing.T) {
    f := factory.NewReplicaFactory()
    tipo := "Usuario"
    entidadID := uuid.New()
    origenID := uuid.New()

    ev, err := f.Create(tipo, entidadID, origenID)
    if err != nil {
        t.Fatalf("se esperaba nil error, se obtuvo: %v", err)
    }

    if ev.EntidadTipo() != tipo {
        t.Errorf("EntidadTipo incorrecto. Esperado %q, obtenido %q", tipo, ev.EntidadTipo())
    }
    if ev.EntidadID() != entidadID {
        t.Errorf("EntidadID incorrecto. Esperado %v, obtenido %v", entidadID, ev.EntidadID())
    }
    if ev.OrigenNodoID() != origenID {
        t.Errorf("OrigenNodoID incorrecto. Esperado %v, obtenido %v", origenID, ev.OrigenNodoID())
    }
    // Comprobamos que EventoAt esté muy próximo al momento actual
    if time.Since(ev.EventoAt()) > time.Second {
        t.Errorf("EventoAt demasiado antiguo: %v", ev.EventoAt())
    }
    // ID no debe ser nil
    if ev.ID() == uuid.Nil {
        t.Error("ID no debería ser Nil")
    }
}

func TestReplicaFactory_Create_Errors(t *testing.T) {
    f := factory.NewReplicaFactory()
    validUUID := uuid.New()

    if _, err := f.Create("", validUUID, validUUID); err != factory.ErrEntidadTipoVacio {
        t.Errorf("se esperaba ErrEntidadTipoVacio, se obtuvo: %v", err)
    }
    if _, err := f.Create("Entidad", uuid.Nil, validUUID); err != factory.ErrEntidadIDNil {
        t.Errorf("se esperaba ErrEntidadIDNil, se obtuvo: %v", err)
    }
    if _, err := f.Create("Entidad", validUUID, uuid.Nil); err != factory.ErrOrigenNodoIDNil {
        t.Errorf("se esperaba ErrOrigenNodoIDNil, se obtuvo: %v", err)
    }
}
