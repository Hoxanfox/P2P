package model

import (
    "errors"
    "time"

    "github.com/google/uuid"
)

// Errores de validación para ArchivoMetadata
var (
    ErrArchivoIDNil          = errors.New("id de archivo inválido")
    ErrNombreOriginalVacio   = errors.New("nombre original vacío")
    ErrTamanoNegativo        = errors.New("tamaño de archivo negativo")
    ErrRutaVacia             = errors.New("ruta de almacenamiento vacía")
    ErrSubidoPorNil          = errors.New("uploader inválido")
    ErrFechaSubidaZero       = errors.New("fecha de subida no puede ser cero")
)

// ArchivoMetadata representa los metadatos de un archivo compartido.
type ArchivoMetadata struct {
    id             uuid.UUID
    nombreOriginal string
    tamanoBytes    int64
    ruta           string
    subidoPor      uuid.UUID
    fechaSubida    time.Time
}

// NewArchivoMetadata crea un ArchivoMetadata validando sus invariantes.
func NewArchivoMetadata(
    id uuid.UUID,
    nombreOriginal string,
    tamanoBytes int64,
    ruta string,
    subidoPor uuid.UUID,
    fechaSubida time.Time,
) (*ArchivoMetadata, error) {
    if id == uuid.Nil {
        return nil, ErrArchivoIDNil
    }
    if nombreOriginal == "" {
        return nil, ErrNombreOriginalVacio
    }
    if tamanoBytes < 0 {
        return nil, ErrTamanoNegativo
    }
    if ruta == "" {
        return nil, ErrRutaVacia
    }
    if subidoPor == uuid.Nil {
        return nil, ErrSubidoPorNil
    }
    if fechaSubida.IsZero() {
        return nil, ErrFechaSubidaZero
    }
    return &ArchivoMetadata{
        id:             id,
        nombreOriginal: nombreOriginal,
        tamanoBytes:    tamanoBytes,
        ruta:           ruta,
        subidoPor:      subidoPor,
        fechaSubida:    fechaSubida,
    }, nil
}

// Getters
func (a *ArchivoMetadata) ID() uuid.UUID     { return a.id }
func (a *ArchivoMetadata) NombreOriginal() string { return a.nombreOriginal }
func (a *ArchivoMetadata) TamanoBytes() int64 { return a.tamanoBytes }
func (a *ArchivoMetadata) Ruta() string       { return a.ruta }
func (a *ArchivoMetadata) SubidoPor() uuid.UUID { return a.subidoPor }
func (a *ArchivoMetadata) FechaSubida() time.Time { return a.fechaSubida }
