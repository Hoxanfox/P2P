package transport;

import java.io.*;
import java.net.Socket;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.LinkedBlockingQueue;

public class TcpConnectionWorker extends Thread {

    private final String host;
    private final int port;

    private Socket socket;
    private BufferedReader in;
    private PrintWriter out;
    private volatile boolean running = true;

    private final BlockingQueue<String> outgoingMessages;
    private final BlockingQueue<String> incomingResponses;

    public TcpConnectionWorker(String host, int port,
                               BlockingQueue<String> outgoingMessages,
                               BlockingQueue<String> incomingResponses) {
        this.host = host;
        this.port = port;
        this.outgoingMessages = outgoingMessages;
        this.incomingResponses = incomingResponses;
    }

    @Override
    public void run() {
        try {
            socket = new Socket(host, port);
            in = new BufferedReader(new InputStreamReader(socket.getInputStream()));
            out = new PrintWriter(socket.getOutputStream(), true);
            System.out.println("[INFO] Conexión establecida con " + host + ":" + port);

            while (running && !socket.isClosed()) {

                // Si hay mensajes para enviar
                String message = outgoingMessages.poll();
                if (message != null) {
                    out.println(message);
                    out.flush();
                    System.out.println("[DEBUG] Enviado: " + message);
                }

                // Si hay mensajes recibidos
                if (in.ready()) {
                    String response = in.readLine();
                    if (response != null) {
                        System.out.println("[DEBUG] Recibido: " + response);
                        incomingResponses.offer(response);
                    }
                }

                Thread.sleep(100); // evitar uso excesivo de CPU
            }

        } catch (Exception e) {
            System.err.println("[ERROR] Hilo de conexión TCP: " + e.getMessage());
        } finally {
            try {
                if (socket != null) socket.close();
                System.out.println("[INFO] Conexión cerrada.");
            } catch (IOException e) {
                System.err.println("[ERROR] Al cerrar socket: " + e.getMessage());
            }
        }
    }

    public void stopRunning() {
        this.running = false;
    }
}
