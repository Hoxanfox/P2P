package transport;

import transport.interfaces.ITransportStrategy;

public class TransportContext {
    private ITransportStrategy strategy;

    public TransportContext(ITransportStrategy strategy) {
        this.strategy = strategy;
    }

    public void setStrategy(ITransportStrategy strategy) {
        this.strategy = strategy;
    }

    public String executeSend(String jsonToSend) {
        return strategy.sendJson(jsonToSend);
    }

    // Método para recibir mensajes
    public String executeReceive() {
        return strategy.receiveJson();  // Llamamos a un método que implementa la estrategia de recepción
    }

    public ITransportStrategy getStrategy() {
        return strategy;
    }
}
