package protocolo.implementaciones.InvitacionCanal;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import dto.implementacion.InvitacionCanal.InvitacionCanalResponseDto;
import protocolo.interfaces.ResponseRoute;

public class InvitacionCanalResponse implements ResponseRoute {

    private String status;
    private String message;
    private InvitacionCanalResponseDto data;

    @Override
    public void fromJson(String jsonResponse) {
        System.out.println("[InvitacionCanalResponse] === INICIO PARSEO JSON ====");
        System.out.println("[InvitacionCanalResponse] ==> JSON recibido:\n" + jsonResponse);

        ObjectMapper mapper = new ObjectMapper();
        mapper.registerModule(new JavaTimeModule()); // Para LocalDateTime
        try {
            ResponseWrapper wrapper = mapper.readValue(jsonResponse, ResponseWrapper.class);
            this.status = wrapper.getStatus();
            this.message = wrapper.getMessage();
            this.data = wrapper.getData();

            System.out.println("[InvitacionCanalResponse] ==> Status: " + status);
            System.out.println("[InvitacionCanalResponse] ==> Message: " + message);
            System.out.println("[InvitacionCanalResponse] ==> Data: " + data);
        } catch (Exception e) {
            System.err.println("[InvitacionCanalResponse] Error al parsear JSON: " + e.getMessage());
            throw new RuntimeException("Error al procesar el JSON de respuesta de invitación al canal", e);
        }

        System.out.println("[InvitacionCanalResponse] === FIN PARSEO JSON ====");
    }

    public String getStatus() { return status; }
    public String getMessage() { return message; }
    public InvitacionCanalResponseDto getData() { return data; }

    // Clase interna para mapear la estructura JSON
    public static class ResponseWrapper {
        private String status;
        private String message;
        private InvitacionCanalResponseDto data;

        public String getStatus() { return status; }
        public void setStatus(String status) { this.status = status; }

        public String getMessage() { return message; }
        public void setMessage(String message) { this.message = message; }

        public InvitacionCanalResponseDto getData() { return data; }
        public void setData(InvitacionCanalResponseDto data) { this.data = data; }
    }
}
