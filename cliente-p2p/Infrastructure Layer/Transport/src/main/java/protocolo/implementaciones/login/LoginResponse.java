package protocolo.implementaciones.login;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import protocolo.interfaces.ResponseRoute;
import dto.implementacion.login.LoginResponseDto;

import java.util.UUID;

public class LoginResponse implements ResponseRoute {

    private LoginResponseDto dto;

    @Override
    public void fromJson(String json) {
        if (json == null || json.isBlank()) {
            System.out.println("[ERROR] JSON nulo o vacío en LoginResponse");
            return;
        }

        System.out.println("[INFO] Iniciando parseo de JSON con Jackson");

        ObjectMapper mapper = new ObjectMapper();

        try {
            JsonNode root = mapper.readTree(json);
            String status = root.path("status").asText(null);
            String message = root.path("message").asText(null);

            if ("success".equalsIgnoreCase(status)) {
                JsonNode data = root.path("data");

                if (!data.isMissingNode()) {
                    dto = new LoginResponseDto(
                            status,
                            message,
                            data.has("id") ? UUID.fromString(data.get("id").asText()).hashCode() : null,
                            data.path("nombre").asText(null),
                            data.path("email").asText(null),
                            data.path("photo").asText(null),
                            data.path("ip").asText(null),
                            data.path("created_at").asText(null),
                            data.path("is_connected").asBoolean(false)
                    );

                    System.out.println("[INFO] Datos cargados correctamente con Jackson");
                } else {
                    System.out.println("[WARN] Nodo 'data' no encontrado");
                }
            } else {
                dto = new LoginResponseDto(status, message, null, null, null, null, null, null, false);
                System.out.println("[INFO] Solo se procesaron status y message");
            }

        } catch (Exception e) {
            System.out.println("[ERROR] Error al parsear JSON con Jackson: " + e.getMessage());
        }
    }

    public LoginResponseDto toDto() {
        return dto;
    }

    public boolean isConnected() {
        return dto != null && dto.isConnected();
    }
}
