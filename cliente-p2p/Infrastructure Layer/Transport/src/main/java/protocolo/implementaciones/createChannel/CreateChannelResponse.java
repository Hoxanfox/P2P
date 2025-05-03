package protocolo.implementaciones.createChannel;

import com.fasterxml.jackson.databind.ObjectMapper;
import dto.implementacion.CreateChannel.CreateChannelResponseDto;
import protocolo.interfaces.ResponseRoute;

public class CreateChannelResponse implements ResponseRoute {

    private String status;
    private String message;
    private CreateChannelResponseDto data;

    @Override
    public void fromJson(String jsonResponse) {
        System.out.println("[CreateChannelResponse] === INICIO PARSEO JSON ====");
        System.out.println("[CreateChannelResponse] ==> JSON de entrada:\n" + jsonResponse);

        try {
            ObjectMapper mapper = new ObjectMapper();
            // Se crea una clase auxiliar temporal para mapear toda la respuesta
            ResponseWrapper wrapper = mapper.readValue(jsonResponse, ResponseWrapper.class);

            this.status = wrapper.getStatus();
            this.message = wrapper.getMessage();
            this.data = wrapper.getData();

            System.out.println("[CreateChannelResponse] ==> Status: " + status);
            System.out.println("[CreateChannelResponse] ==> Message: " + message);
            System.out.println("[CreateChannelResponse] ==> Data: " + data);
        } catch (Exception e) {
            System.err.println("[CreateChannelResponse] Error al parsear el JSON: " + e.getMessage());
        }

        System.out.println("[CreateChannelResponse] === FIN PARSEO JSON ====");
    }

    public String getStatus() {
        return status;
    }

    public String getMessage() {
        return message;
    }

    public CreateChannelResponseDto getData() {
        return data;
    }

    // Clase interna para mapear el JSON completo
    public static class ResponseWrapper {
        private String status;
        private String message;
        private CreateChannelResponseDto data;

        public String getStatus() { return status; }
        public void setStatus(String status) { this.status = status; }

        public String getMessage() { return message; }
        public void setMessage(String message) { this.message = message; }

        public CreateChannelResponseDto getData() { return data; }
        public void setData(CreateChannelResponseDto data) { this.data = data; }
    }
}
