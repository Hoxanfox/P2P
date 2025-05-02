package dto.implementacion.CreateChannel.Channel;

import java.util.UUID;

public class MiembroCanalDto {
    private UUID id;
    private String nombre;
    private String email;
    private boolean estado;

    public MiembroCanalDto() {}

    public MiembroCanalDto(UUID id, String nombre, String email, boolean estado) {
        this.id = id;
        this.nombre = nombre;
        this.email = email;
        this.estado = estado;
    }

    // Getters y Setters
}
