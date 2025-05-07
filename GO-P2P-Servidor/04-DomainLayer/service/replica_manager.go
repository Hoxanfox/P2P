package service

import (
	"github.com/google/uuid"
	"model"
)

// ReplicaManager define las operaciones para propagar cambios de entidades a peers
type ReplicaManager interface {
	// ReplicaUser propaga cambios de usuario a otros nodos
	ReplicaUser(usuario *model.UsuarioServidor) error
	
	// ReplicaMessage propaga cambios de mensaje a otros nodos
	ReplicaMessage(message *model.MensajeServidor) error
	
	// ReplicaChannel propaga cambios de canal a otros nodos
	ReplicaChannel(channel *model.CanalServidor) error
	
	// ReplicaFile propaga cambios de archivo a otros nodos
	ReplicaFile(file *model.ArchivoMetadata) error
	
	// ListPendingEvents lista los eventos de replicación pendientes para un nodo específico
	ListPendingEvents(peerID uuid.UUID) ([]*model.ReplicaEvent, error)
}
