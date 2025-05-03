package dto.implementacion.CreateChannel.Channel;

import java.time.LocalDateTime;
import java.util.UUID;

public class InvitacionDto {
    private UUID id;
    private DestinatarioDto destinatario;
    private String fechaEnvio;
    private String estado;

    public InvitacionDto() {}

    public InvitacionDto( DestinatarioDto destinatario) {
        this.destinatario = destinatario;

    }

    // Getters y Setters
    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public DestinatarioDto getDestinatario() {
        return destinatario;
    }

    public void setDestinatario(DestinatarioDto destinatario) {
        this.destinatario = destinatario;
    }

    public String getFechaEnvio() {
        return fechaEnvio;
    }

    public void setFechaEnvio(String fechaEnvio) {
        this.fechaEnvio = fechaEnvio;
    }

    public String getEstado() {
        return estado;
    }

    public void setEstado(String estado) {
        this.estado = estado;
    }
}
