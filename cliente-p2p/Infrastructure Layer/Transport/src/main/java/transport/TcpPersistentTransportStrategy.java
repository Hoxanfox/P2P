package transport;

import transport.interfaces.ITransportStrategy;

import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;

public class TcpPersistentTransportStrategy implements ITransportStrategy {

    private final BlockingQueue<String> outgoingQueue = new LinkedBlockingQueue<>();
    private final BlockingQueue<String> incomingQueue = new LinkedBlockingQueue<>();

    private TcpConnectionWorker worker;

    public TcpPersistentTransportStrategy(String host, int port) {
        this.worker = new TcpConnectionWorker(host, port, outgoingQueue, incomingQueue);
        this.worker.start();
    }

    @Override
    public String sendJson(String jsonToSend) {
        try {
            outgoingQueue.offer(jsonToSend);
            // Esperamos respuesta (bloqueante hasta que llegue o timeout)
            String response = incomingQueue.take();
            return response;
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            return "[ERROR] Interrumpido mientras esperaba respuesta.";
        }
    }

    // Método para recibir mensajes
    public String receiveJson() {
        try {
            // Esperamos un mensaje entrante de la cola de mensajes
            return incomingQueue.take();  // Bloqueante hasta que un mensaje esté disponible
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            return "[ERROR] Interrumpido mientras esperaba mensaje entrante.";
        }
    }

    public void close() {
        if (worker != null) {
            worker.stopRunning();
        }
    }
}
