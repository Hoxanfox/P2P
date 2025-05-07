package model

import (
    "testing"
    "time"

    "github.com/google/uuid"
)

func TestNewArchivoMetadata_Success(t *testing.T) {
    id := uuid.New()
    uploader := uuid.New()
    now := time.Now().UTC().Truncate(time.Second)

    a, err := NewArchivoMetadata(
        id,
        "documento.pdf",
        2048,
        "/var/files/documento.pdf",
        uploader,
        now,
    )
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if a.ID() != id {
        t.Errorf("ID: esperado %v, obtuvo %v", id, a.ID())
    }
    if a.NombreOriginal() != "documento.pdf" {
        t.Errorf("NombreOriginal: esperado %q, obtuvo %q", "documento.pdf", a.NombreOriginal())
    }
    if a.TamanoBytes() != 2048 {
        t.Errorf("TamanoBytes: esperado %d, obtuvo %d", 2048, a.TamanoBytes())
    }
    if a.Ruta() != "/var/files/documento.pdf" {
        t.Errorf("Ruta: esperado %q, obtuvo %q", "/var/files/documento.pdf", a.Ruta())
    }
    if a.SubidoPor() != uploader {
        t.Errorf("SubidoPor: esperado %v, obtuvo %v", uploader, a.SubidoPor())
    }
    if !a.FechaSubida().Equal(now) {
        t.Errorf("FechaSubida: esperado %v, obtuvo %v", now, a.FechaSubida())
    }
}

func TestNewArchivoMetadata_Errors(t *testing.T) {
    validID := uuid.New()
    validUploader := uuid.New()
    now := time.Now().UTC()

    cases := []struct {
        name           string
        id             uuid.UUID
        nombreOriginal string
        tamanoBytes    int64
        ruta           string
        subidoPor      uuid.UUID
        fechaSubida    time.Time
        wantErr        error
    }{
        {"ID inválido", uuid.Nil, "f", 1, "/r", validUploader, now, ErrArchivoIDNil},
        {"Nombre vacío", validID, "", 1, "/r", validUploader, now, ErrNombreOriginalVacio},
        {"Tamaño negativo", validID, "f", -5, "/r", validUploader, now, ErrTamanoNegativo},
        {"Ruta vacía", validID, "f", 1, "", validUploader, now, ErrRutaVacia},
        {"Uploader inválido", validID, "f", 1, "/r", uuid.Nil, now, ErrSubidoPorNil},
        {"Fecha cero", validID, "f", 1, "/r", validUploader, time.Time{}, ErrFechaSubidaZero},
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            _, err := NewArchivoMetadata(
                c.id,
                c.nombreOriginal,
                c.tamanoBytes,
                c.ruta,
                c.subidoPor,
                c.fechaSubida,
            )
            if err != c.wantErr {
                t.Errorf("%s: esperado error %v, obtuvo %v", c.name, c.wantErr, err)
            }
        })
    }
}
