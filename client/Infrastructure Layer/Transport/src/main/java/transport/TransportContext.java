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
}
