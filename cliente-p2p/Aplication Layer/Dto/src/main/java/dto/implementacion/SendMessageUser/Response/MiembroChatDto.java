package dto.implementacion.SendMessageUser.Response;

import java.util.UUID;

public class MiembroChatDto {
    private UUID id;
    private String email;

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
