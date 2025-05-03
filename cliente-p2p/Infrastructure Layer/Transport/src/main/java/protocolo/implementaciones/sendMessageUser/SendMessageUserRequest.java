package protocolo.implementaciones.sendMessageUser;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import dto.implementacion.SendMessageUser.Request.SendMessageUserRequestDto;
import protocolo.interfaces.RequestRoute;

public class SendMessageUserRequest implements RequestRoute {

    private String command = "send-message-user";  // Campo command
    private SendMessageUserRequestDto data;

    public SendMessageUserRequest(SendMessageUserRequestDto data) {
        System.out.println("[DEBUG] Clase: SendMessageUserRequest - Constructor ejecutado.");
        System.out.println("[DEBUG] DTO recibido: " + data);
        this.data = data;
    }

    @Override
    public String toJson() {
        System.out.println("[DEBUG] Clase: SendMessageUserRequest - Método: toJson()");
        ObjectMapper mapper = new ObjectMapper();
        try {
            // Crear el wrapper para incluir tanto el command como los datos
            RequestWrapper wrapper = new RequestWrapper();
            wrapper.setCommand(this.command);
            wrapper.setData(this.data);

            // Convertir el wrapper completo a JSON
            String json = mapper.writeValueAsString(wrapper);
            System.out.println("[DEBUG] JSON generado: " + json);
            return json;
        } catch (JsonProcessingException e) {
            System.out.println("[ERROR] Error al convertir a JSON en SendMessageUserRequest: " + e.getMessage());
            throw new RuntimeException("Error al convertir a JSON la solicitud de envío de mensaje", e);
        }
    }

    public SendMessageUserRequestDto getData() {
        System.out.println("[DEBUG] Clase: SendMessageUserRequest - Método: getData()");
        return data;
    }

    public void setData(SendMessageUserRequestDto data) {
        System.out.println("[DEBUG] Clase: SendMessageUserRequest - Método: setData()");
        System.out.println("[DEBUG] Nuevo DTO asignado: " + data);
        this.data = data;
    }

    // Clase interna para envolver la estructura del JSON
    public static class RequestWrapper {
        private String command;
        private SendMessageUserRequestDto data;

        public String getCommand() { return command; }
        public void setCommand(String command) { this.command = command; }

        public SendMessageUserRequestDto getData() { return data; }
        public void setData(SendMessageUserRequestDto data) { this.data = data; }
    }
}
