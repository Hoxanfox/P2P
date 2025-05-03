package protocolo.implementaciones.sendMessageUser;

import com.fasterxml.jackson.databind.ObjectMapper;
import dto.implementacion.SendMessageUser.Response.SendMessageUserResponseDto;
import protocolo.interfaces.ResponseRoute;

public class SendMessageUserResponse implements ResponseRoute {

    private String status;
    private String message;
    private SendMessageUserResponseDto data;

    @Override
    public void fromJson(String jsonResponse) {
        System.out.println("[SendMessageUserResponse] === INICIO PARSEO JSON ====");
        System.out.println("[SendMessageUserResponse] ==> JSON recibido:\n" + jsonResponse);

        ObjectMapper mapper = new ObjectMapper();
        try {
            ResponseWrapper wrapper = mapper.readValue(jsonResponse, ResponseWrapper.class);
            this.status = wrapper.getStatus();
            this.message = wrapper.getMessage();
            this.data = wrapper.getData();
            System.out.println("[SendMessageUserResponse] ==> Status: " + status);
            System.out.println("[SendMessageUserResponse] ==> Message: " + message);
            System.out.println("[SendMessageUserResponse] ==> Data: " + data);
        } catch (Exception e) {
            System.err.println("[SendMessageUserResponse] Error al parsear JSON: " + e.getMessage());
            throw new RuntimeException("Error al procesar el JSON de respuesta de envío de mensaje", e);
        }

        System.out.println("[SendMessageUserResponse] === FIN PARSEO JSON ====");
    }

    public String getStatus() { return status; }
    public String getMessage() { return message; }
    public SendMessageUserResponseDto getData() { return data; }

    // Clase interna para mapear la respuesta
    public static class ResponseWrapper {
        private String status;
        private String message;
        private SendMessageUserResponseDto data;

        public String getStatus() { return status; }
        public void setStatus(String status) { this.status = status; }

        public String getMessage() { return message; }
        public void setMessage(String message) { this.message = message; }

        public SendMessageUserResponseDto getData() { return data; }
        public void setData(SendMessageUserResponseDto data) { this.data = data; }
    }
}
