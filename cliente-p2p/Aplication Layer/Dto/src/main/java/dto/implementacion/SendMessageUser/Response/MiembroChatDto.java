package dto.implementacion.SendMessageUser.Response;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.UUID;

public class MiembroChatDto {
    private UUID id;

    @JsonProperty("correo")  // Mapea el campo "correo" del JSON a "email" en la clase
    private String email;

    public MiembroChatDto () {
    }

    // Getter para 'id'
    public UUID getId() {
        return id;
    }

    // Setter para 'id'
    public void setId(UUID id) {
        this.id = id;
    }

    // Getter para 'email'
    public String getEmail() {
        return email;
    }

    // Setter para 'email'
    public void setEmail(String email) {
        this.email = email;
    }
}
