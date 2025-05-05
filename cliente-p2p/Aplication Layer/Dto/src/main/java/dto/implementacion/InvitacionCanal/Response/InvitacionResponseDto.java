package dto.implementacion.InvitacionCanal.Response;

import java.time.LocalDateTime;

public class InvitacionResponseDto {
    private Long id;
    private LocalDateTime fechaEnvio;
    private String estado;
    private CanalResponseDto canal;
    private DestinatarioResponseDto destinatario;

    // Constructor por defecto
    public InvitacionResponseDto() {
    }

    // Getter y Setter para id
    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    // Getter y Setter para fechaEnvio
    public LocalDateTime getFechaEnvio() {
        return fechaEnvio;
    }

    public void setFechaEnvio(LocalDateTime fechaEnvio) {
        this.fechaEnvio = fechaEnvio;
    }

    // Getter y Setter para estado
    public String getEstado() {
        return estado;
    }

    public void setEstado(String estado) {
        this.estado = estado;
    }

    // Getter y Setter para canal
    public CanalResponseDto getCanal() {
        return canal;
    }

    public void setCanal(CanalResponseDto canal) {
        this.canal = canal;
    }

    // Getter y Setter para destinatario
    public DestinatarioResponseDto getDestinatario() {
        return destinatario;
    }

    public void setDestinatario(DestinatarioResponseDto destinatario) {
        this.destinatario = destinatario;
    }
}
