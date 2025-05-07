package service

import (
	"time"

	"github.com/google/uuid"
	"model"
)

// FileDTO representa un archivo simplificado para transferencia de datos
type FileDTO struct {
	ID             uuid.UUID `json:"id"`
	Nombre         string    `json:"nombre"`
	Tamaño         int64     `json:"tamaño"`
	RutaAlmacenada string    `json:"rutaAlmacenada"`
	UploaderID     uuid.UUID `json:"uploaderID"`
	FechaSubida    time.Time `json:"fechaSubida"`
	NumChunks      int       `json:"numChunks"`
	Completo       bool      `json:"completo"`
}

// ChunkDTO representa un fragmento de archivo para transferencia progresiva
type ChunkDTO struct {
	FileID     uuid.UUID `json:"fileID"`
	ChunkIndex int       `json:"chunkIndex"`
	Data       []byte    `json:"data"`
	Size       int       `json:"size"`
	IsLast     bool      `json:"isLast"`
}

// MapArchivoToDTO convierte un modelo ArchivoMetadata a un DTO
// Nota: Asumimos que estos métodos existen o serán implementados en ArchivoMetadata
func MapArchivoToDTO(a *model.ArchivoMetadata, numChunks int, completo bool) *FileDTO {
	// Los nombres exactos de los métodos deben coincidir con los de ArchivoMetadata
	return &FileDTO{
		ID:             a.ID(), // Asumimos que existe ID()
		Nombre:         a.Nombre(), // Asumimos que existe Nombre()
		Tamaño:         a.Tamaño(), // Asumimos que existe Tamaño()
		RutaAlmacenada: a.RutaAlmacenada(), // Asumimos que existe RutaAlmacenada()
		UploaderID:     a.UploaderID(), // Asumimos que existe UploaderID()
		FechaSubida:    a.FechaSubida(), // Asumimos que existe FechaSubida()
		NumChunks:      numChunks,
		Completo:       completo,
	}
}
