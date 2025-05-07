package service

import (
	"github.com/google/uuid"
	"model"
)

// ArchivoChunk representa un fragmento de archivo para transferencia progresiva
type ArchivoChunk struct {
	FileID     uuid.UUID // ID del archivo al que pertenece
	ChunkIndex int       // Índice del fragmento (0-based)
	Data       []byte    // Datos del fragmento
	Size       int       // Tamaño en bytes
	IsLast     bool      // Indica si es el último fragmento
}

// FileService define las operaciones para transferencia y descarga de archivos en chats
type FileService interface {
	// SendToUser envía un archivo a un usuario específico
	SendToUser(
		remitenteID, destinatarioID uuid.UUID,
		nombre string, datos []byte,
	) (*model.ArchivoMetadata, error)

	// SendToChannel envía un archivo a un canal
	SendToChannel(
		remitenteID, channelID uuid.UUID,
		nombre string, datos []byte,
	) (*model.ArchivoMetadata, error)

	// DownloadChunk descarga un fragmento específico de un archivo
	DownloadChunk(
		fileID uuid.UUID,
		chunkIndex int,
	) (*ArchivoChunk, error)

	// FinalizeDownload finaliza la descarga de un archivo y confirma su integridad
	FinalizeDownload(
		fileID uuid.UUID,
	) (*model.ArchivoMetadata, error)
}
