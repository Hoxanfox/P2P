package interfaces

// ITransportStrategy define el contrato que debe cumplir cualquier estrategia de transporte.
type ITransportStrategy interface {
    SendJson(jsonToSend string) string
}
