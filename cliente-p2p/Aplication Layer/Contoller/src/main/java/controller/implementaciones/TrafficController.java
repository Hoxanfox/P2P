package controller.implementaciones;

import dto.implementacion.Traffic.TrafficResponseDto;
import protocolo.implementaciones.traffic.TrafficResponse;
import transport.TransportContext;

public class TrafficController {

    private final TransportContext transportContext;

    public TrafficController(TransportContext transportContext) {
        this.transportContext = transportContext;
    }

    public void escuchar() {
        // Inicia un hilo separado para escuchar los mensajes entrantes
        new Thread(() -> {
            while (true) {
                try {
                    // Simulamos la recepción de un JSON desde el servidor
                    String jsonResponse = transportContext.executeReceive();  // método bloqueante

                    TrafficResponse response = new TrafficResponse();
                    response.fromJson(jsonResponse);

                    TrafficResponseDto responseDTO = response.getResponseDTO();
                    if (responseDTO == null) {
                        System.out.println("Respuesta malformada.");
                        continue;
                    }

                    // Verifica el comando recibido
                    manejarComando(responseDTO);

                } catch (Exception e) {
                    e.printStackTrace();
                    break;
                }
            }
        }).start();
    }

    private void manejarComando(TrafficResponseDto dto) {
        String command = dto.getCommand();
        System.out.println("Comando recibido: " + command);

        // Verifica si el comando es "refresh-users"
        if ("refresh-users".equals(command)) {
            // Llama al controlador de usuarios para refrescar la lista
            System.out.println("🟢 Refrescando lista de usuarios...");
            ListUsersController listUsersController = new ListUsersController(transportContext);
            listUsersController.obtenerUsuarios();
        } else {
            System.out.println("⚠️ Comando no reconocido: " + command);
        }
    }
}
