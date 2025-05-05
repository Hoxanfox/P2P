package transport;

import java.io.*;
import java.net.*;
import java.util.concurrent.BlockingQueue;

public class TcpConnectionWorker extends Thread {

    private final String host;
    private final int port;
    private final BlockingQueue<String> outgoingQueue;
    private final BlockingQueue<String> incomingQueue;
    private volatile boolean running = true;
    private Socket socket;
    private PrintWriter writer;
    private BufferedReader reader;

    public TcpConnectionWorker(String host, int port, BlockingQueue<String> outgoingQueue, BlockingQueue<String> incomingQueue) {
        this.host = host;
        this.port = port;
        this.outgoingQueue = outgoingQueue;
        this.incomingQueue = incomingQueue;
    }

    @Override
    public void run() {
        try {
            socket = new Socket(host, port);
            writer = new PrintWriter(socket.getOutputStream(), true);
            reader = new BufferedReader(new InputStreamReader(socket.getInputStream()));

            // Hilo para escuchar mensajes entrantes
            new Thread(() -> {
                try {
                    String line;
                    while (running && (line = reader.readLine()) != null) {
                        incomingQueue.offer(line);  // Colocamos los mensajes entrantes en la cola
                    }
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }).start();

            // Hilo para enviar mensajes
            while (running) {
                String message = outgoingQueue.take();  // Tomamos los mensajes de la cola de salida
                writer.println(message);
            }

        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }
    }

    public void stopRunning() {
        running = false;
        try {
            socket.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
