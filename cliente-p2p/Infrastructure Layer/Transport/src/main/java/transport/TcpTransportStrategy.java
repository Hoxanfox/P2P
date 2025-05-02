package transport;

import java.io.*;
import java.net.Socket;
import transport.interfaces.ITransportStrategy;

    public class TcpTransportStrategy implements ITransportStrategy {

    private final String host;
    private final int port;

    public TcpTransportStrategy(String host, int port) {
        this.host = host;
        this.port = port;
    }

    @Override
    public String sendJson(String jsonToSend) {
        System.out.println("[DEBUG] Intentando conectar a " + host + ":" + port);

        try (Socket socket = new Socket(host, port)) {
            System.out.println("[DEBUG] Conexión establecida.");

            // Enviar el JSON
            OutputStream output = socket.getOutputStream();
            PrintWriter writer = new PrintWriter(new OutputStreamWriter(output), true);
            System.out.println("[DEBUG] Enviando JSON: " + jsonToSend);
            writer.println(jsonToSend);

            // Leer la respuesta
            InputStream input = socket.getInputStream();
            BufferedReader reader = new BufferedReader(new InputStreamReader(input));
            String response = reader.readLine();

            System.out.println("[DEBUG] Respuesta recibida: " + response);

            return response;

        } catch (IOException e) {
            System.err.println("[ERROR] Error durante la comunicación: " + e.getMessage());
            e.printStackTrace();
            return null;
        }
    }
}
