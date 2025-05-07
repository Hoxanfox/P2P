package observer

import (
	"model"
)

// IReplicaObserver define los callbacks para recibir notificaciones cuando el ReplicaManager
// envía con éxito datos a otro nodo: usuarios, canales, mensajes o archivos.
type IReplicaObserver interface {
	// OnUserReplicated se invoca cuando se replica con éxito un usuario a otro nodo
	OnUserReplicated(user *model.UsuarioServidor, toPeer *model.Peer)

	// OnChannelReplicated se invoca cuando se replica con éxito un canal a otro nodo
	OnChannelReplicated(channel *model.CanalServidor, toPeer *model.Peer)

	// OnMessageReplicated se invoca cuando se replica con éxito un mensaje a otro nodo
	OnMessageReplicated(msg *model.MensajeServidor, toPeer *model.Peer)

	// OnFileReplicated se invoca cuando se replica con éxito un archivo a otro nodo
	OnFileReplicated(file *model.ArchivoMetadata, toPeer *model.Peer)
}
