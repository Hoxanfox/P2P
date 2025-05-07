package service

import (
	"github.com/google/uuid"
)

// FileService define las operaciones para transferencia y descarga de archivos en chats
type FileService interface {
	// SendToUser envía un archivo a un usuario específico
	SendToUser(
		remitenteID, destinatarioID uuid.UUID,
		nombre string, datos []byte,
	) (*FileDTO, error)

	// SendToChannel envía un archivo a un canal
	SendToChannel(
		remitenteID, channelID uuid.UUID,
		nombre string, datos []byte,
	) (*FileDTO, error)

	// DownloadChunk descarga un fragmento específico de un archivo
	DownloadChunk(
		fileID uuid.UUID,
		chunkIndex int,
	) (*ChunkDTO, error)

	// FinalizeDownload finaliza la descarga de un archivo y confirma su integridad
	FinalizeDownload(
		fileID uuid.UUID,
	) (*FileDTO, error)
}
