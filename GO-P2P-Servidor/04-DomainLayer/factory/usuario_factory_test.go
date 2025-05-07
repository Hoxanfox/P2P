package factory

import (
    "testing"
    "time"

    "github.com/google/uuid"
    "model"
)

func TestNewUsuarioFactory(t *testing.T) {
    fact := NewUsuarioFactory()
    if fact == nil {
        t.Fatal("NewUsuarioFactory devolvió nil")
    }
}

func TestUsuarioFactory_Create_Success(t *testing.T) {
    fact := NewUsuarioFactory()
    nombre := "juanperez"
    email := "juan.perez@example.com"
    hash := "hashedPassword123"
    foto := "https://example.com/avatar.png"
    ip := "192.168.1.42"

    usr, err := fact.Create(nombre, email, hash, foto, ip)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if usr.ID() == uuid.Nil {
        t.Errorf("ID: esperaba no-nil, obtuvo %v", usr.ID())
    }
    if usr.NombreUsuario() != nombre {
        t.Errorf("NombreUsuario: esperado %q, obtuvo %q", nombre, usr.NombreUsuario())
    }
    if usr.Email() != email {
        t.Errorf("Email: esperado %q, obtuvo %q", email, usr.Email())
    }
    if usr.ContrasenaHasheada() != hash {
        t.Errorf("ContrasenaHasheada: esperado %q, obtuvo %q", hash, usr.ContrasenaHasheada())
    }
    if usr.FotoURL() != foto {
        t.Errorf("FotoURL: esperado %q, obtuvo %q", foto, usr.FotoURL())
    }
    if usr.IPRegistrada() != ip {
        t.Errorf("IPRegistrada: esperado %q, obtuvo %q", ip, usr.IPRegistrada())
    }
    if usr.FechaRegistro().IsZero() {
        t.Error("FechaRegistro: esperado timestamp no cero")
    }
    
    // Verificar que por defecto no está conectado
    if usr.IsConnected() {
        t.Error("IsConnected: esperaba false por defecto, obtuvo true")
    }
    
    // Comprobación opcional de que FechaRegistro esté cerca de ahora
    if diff := time.Since(usr.FechaRegistro()); diff < 0 || diff > time.Second {
        t.Errorf("FechaRegistro: marca de tiempo inesperada, dif=%v", diff)
    }
}

func TestUsuarioFactory_CreateConnected_Success(t *testing.T) {
    fact := NewUsuarioFactory()
    nombre := "juanperez"
    email := "juan.perez@example.com"
    hash := "hashedPassword123"
    foto := "https://example.com/avatar.png"
    ip := "192.168.1.42"

    usr, err := fact.CreateConnected(nombre, email, hash, foto, ip)
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if usr.ID() == uuid.Nil {
        t.Errorf("ID: esperaba no-nil, obtuvo %v", usr.ID())
    }
    if usr.NombreUsuario() != nombre {
        t.Errorf("NombreUsuario: esperado %q, obtuvo %q", nombre, usr.NombreUsuario())
    }
    if usr.Email() != email {
        t.Errorf("Email: esperado %q, obtuvo %q", email, usr.Email())
    }
    if usr.ContrasenaHasheada() != hash {
        t.Errorf("ContrasenaHasheada: esperado %q, obtuvo %q", hash, usr.ContrasenaHasheada())
    }
    if usr.FotoURL() != foto {
        t.Errorf("FotoURL: esperado %q, obtuvo %q", foto, usr.FotoURL())
    }
    if usr.IPRegistrada() != ip {
        t.Errorf("IPRegistrada: esperado %q, obtuvo %q", ip, usr.IPRegistrada())
    }
    if usr.FechaRegistro().IsZero() {
        t.Error("FechaRegistro: esperado timestamp no cero")
    }
    
    // Verificar que está conectado
    if !usr.IsConnected() {
        t.Error("IsConnected: esperaba true, obtuvo false")
    }
    
    // Comprobación opcional de que FechaRegistro esté cerca de ahora
    if diff := time.Since(usr.FechaRegistro()); diff < 0 || diff > time.Second {
        t.Errorf("FechaRegistro: marca de tiempo inesperada, dif=%v", diff)
    }
}

func TestUsuarioFactory_Create_Errors(t *testing.T) {
    fact := NewUsuarioFactory()
    cases := []struct {
        name                string
        nombre, email, hash string
        wantErr             error
    }{
        {"nombre vacío", "", "u@e.com", "h", model.ErrNombreVacio},
        {"email inválido", "u", "invalid-email", "h", model.ErrEmailInvalido},
        {"hash vacío", "u", "u@e.com", "", model.ErrHashVacio},
    }
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            _, err := fact.Create(c.nombre, c.email, c.hash, "f", "ip")
            if err != c.wantErr {
                t.Errorf("%s: esperado error %v, obtuvo %v", c.name, c.wantErr, err)
            }
        })
    }
}
