package protocolo.implementaciones.InvitacionCanal;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import dto.implementacion.InvitacionCanal.InvitacionCanalRequestDto;
import protocolo.interfaces.RequestRoute;

public class InvitacionCanalRequest implements RequestRoute {

    private String command;
    private InvitacionCanalRequestDto data;

    // Constructor por defecto
    public InvitacionCanalRequest() {
        this.command = "invite-to-channel"; // Asignar el comando fijo
    }

    // Constructor que recibe un DTO y lo mapea a este objeto
    public InvitacionCanalRequest(InvitacionCanalRequestDto dto) {
        this.command = "invite-to-channel"; // Asignar el comando
        this.data = dto; // Asignar directamente el DTO recibido
    }

    // Getter para 'command'
    public String getCommand() {
        return command;
    }

    public void setCommand(String command) {
        this.command = command;
    }

    // Getter para 'data'
    public InvitacionCanalRequestDto getData() {
        return data;
    }

    public void setData(InvitacionCanalRequestDto data) {
        this.data = data;
    }

    // Implementación del método toJson() de la interfaz RequestRoute
    @Override
    public String toJson() {
        try {
            ObjectMapper objectMapper = new ObjectMapper();
            objectMapper.registerModule(new JavaTimeModule()); // 👈 Importante
            objectMapper.disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS); // 👈 Para obtener ISO 8601
            return objectMapper.writeValueAsString(this);
        } catch (Exception e) {
            e.printStackTrace();
            return "{}";
        }
    }
}
