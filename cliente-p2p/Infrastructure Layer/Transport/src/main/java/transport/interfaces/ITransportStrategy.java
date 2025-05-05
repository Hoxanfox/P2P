package transport.interfaces;

public interface ITransportStrategy {
    String sendJson(String jsonToSend);

    // Método para recibir mensajes
    String receiveJson();
}
