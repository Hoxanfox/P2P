package dto.implementacion.SendMessageUser;

import java.util.UUID;

public class RemitenteMessageDto {
    private UUID id;
    private String correo;
    public RemitenteMessageDto() {
    }
    public RemitenteMessageDto(UUID id , String corre) {
        this.id = id;
        this.correo = corre;
    }
    // Getters y setters
    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public String getCorreo() {
        return correo;
    }

    public void setCorreo(String correo) {
        this.correo = correo;
    }
}
