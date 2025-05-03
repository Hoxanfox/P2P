package dto.implementacion.SendMessageUser;

import java.util.UUID;

public class DestinatarioMessageDto {
    private UUID id;
    private String correo;


    public DestinatarioMessageDto(UUID id, String correo) {
        this.id = id;
        this.correo = correo;
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
