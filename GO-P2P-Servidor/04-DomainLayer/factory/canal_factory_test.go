package factory_test

import (
    "testing"

    "github.com/google/uuid"
    "model"
    "factory"
)

func TestCreateValidCanal(t *testing.T) {
    f := factory.NewCanalFactory()
    canal, err := f.Create("soporte", "Canal de soporte", model.CanalPublico)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if canal == nil {
        t.Fatal("esperaba un CanalServidor no nulo")
    }
    if canal.ID() == uuid.Nil {
        t.Error("esperaba un ID válido, obtuvo Nil")
    }
    if canal.Nombre() != "soporte" {
        t.Errorf("esperaba Nombre 'soporte', obtuvo '%s'", canal.Nombre())
    }
    if canal.Descripcion() != "Canal de soporte" {
        t.Errorf("esperaba Descripcion 'Canal de soporte', obtuvo '%s'", canal.Descripcion())
    }
    if canal.Tipo() != model.CanalPublico {
        t.Errorf("esperaba Tipo PUBLICO, obtuvo %v", canal.Tipo())
    }
}

func TestCreateEmptyNombre(t *testing.T) {
    f := factory.NewCanalFactory()
    _, err := f.Create("", "desc", model.CanalPrivado)
    if err != model.ErrCanalNombreVacio {
        t.Errorf("esperaba ErrCanalNombreVacio, obtuvo %v", err)
    }
}

func TestCreateInvalidTipo(t *testing.T) {
    f := factory.NewCanalFactory()
    // "DESCONOCIDO" no es ni PUBLICO ni PRIVADO
    _, err := f.Create("test", "desc", model.CanalTipo("DESCONOCIDO"))
    if err != model.ErrCanalTipoInvalido {
        t.Errorf("esperaba ErrCanalTipoInvalido, obtuvo %v", err)
    }
}
