package dto.implementacion.CreateChannel.Channel;

import java.util.UUID;

public class DestinatarioDto {
    private UUID id;
    private String email;

    public DestinatarioDto() {}

    public DestinatarioDto(UUID id, String email) {
        this.id = id;
        this.email = email;
    }

    // Getters y Setters
    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }
}
