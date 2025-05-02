package dto.implementacion.ListUsers;

import java.util.UUID;

public class UsuarioResponseDTO {
    private UUID id;
    private String nombre;
    private String email;
    private boolean is_connected;

    public UsuarioResponseDTO(UUID id, String nombre, String email, boolean is_connected) {
        this.id = id;
        this.nombre = nombre;
        this.email = email;
        this.is_connected = is_connected;
    }

    public UUID getId() {
        return id;
    }

    public String getNombre() {
        return nombre;
    }

    public String getEmail() {
        return email;
    }

    public boolean isIs_connected() {
        return is_connected;
    }
}
