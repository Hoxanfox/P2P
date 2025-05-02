package app;

import protocolo.implementaciones.register.RegisterRequest;
import protocolo.implementaciones.register.RegisterResponse;
import transport.TcpTransportStrategy;
import transport.TransportContext;

public class Main {
    public static void main(String[] args) {
        // Datos para registrar el usuario
        String username = "juan123";
        String email = "juan@mail.com";
        String password = "secreta123";
        String photo = "foto_base64"; // Podría ser codificado en base64 real
        String ip = "192.168.1.23";   // IP manual (en futuro podrías obtenerlo dinámicamente)

        // 1. Crear el objeto de la solicitud (Request)
        RegisterRequest registerRequest = new RegisterRequest(username, email, password, photo, ip);

        // 2. Crear el contexto de transporte (TransportContext)
        // Cambia "localhost" y 12345 por tu servidor real si es necesario
        TransportContext context = new TransportContext(new TcpTransportStrategy("localhost", 9000));

        // 3. Enviar el mensaje y recibir la respuesta JSON
        System.out.println("[DEBUG] Enviando solicitud de registro...");
        String jsonResponse = context.executeSend(registerRequest.toJson());

        // 4. Procesar la respuesta
        if (jsonResponse != null) {
            RegisterResponse registerResponse = new RegisterResponse();
            registerResponse.fromJson(jsonResponse);1

            // 5. Mostrar los resultados
            System.out.println("[RESULTADO] Estado: " + registerResponse.getStatus());
            System.out.println("[RESULTADO] Mensaje: " + registerResponse.getMessage());
            System.out.println("[RESULTADO] Usuario: " + registerResponse.getUsername());
            System.out.println("[RESULTADO] ID Usuario: " + registerResponse.getUserId());
        } else {
            System.out.println("[ERROR] No se recibió respuesta del servidor.");
        }
    }
}
