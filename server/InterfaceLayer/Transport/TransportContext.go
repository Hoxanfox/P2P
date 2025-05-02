package transport

// TransportContext mantiene una referencia a una estrategia de transporte
type TransportContext struct {
    strategy ITransportStrategy
}

// NewTransportContext crea una nueva instancia de TransportContext
func NewTransportContext(strategy ITransportStrategy) *TransportContext {
    return &TransportContext{
        strategy: strategy,
    }
}

// SetStrategy cambia la estrategia de transporte
func (t *TransportContext) SetStrategy(strategy ITransportStrategy) {
    t.strategy = strategy
}

// ExecuteSend ejecuta el envío del JSON usando la estrategia actual
func (t *TransportContext) ExecuteSend(jsonToSend string) string {
    return t.strategy.SendJson(jsonToSend)
}
