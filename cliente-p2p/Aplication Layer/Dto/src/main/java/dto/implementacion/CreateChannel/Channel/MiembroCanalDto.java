package dto.implementacion.CreateChannel.Channel;

import java.util.UUID;

public class MiembroCanalDto {
    private UUID id;
    private String nombre;
    private String email;
    private boolean estado;

    public MiembroCanalDto() {}

    public MiembroCanalDto(UUID id, String email) {
        this.id = id;

        this.email = email;
    }

    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public String getNombre() {
        return nombre;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public boolean isEstado() {
        return estado;
    }

    public void setEstado(boolean estado) {
        this.estado = estado;
    }
}
