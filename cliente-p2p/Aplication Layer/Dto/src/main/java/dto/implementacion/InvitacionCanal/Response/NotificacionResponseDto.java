package dto.implementacion.InvitacionCanal.Response;

public class NotificacionResponseDto {
    private Long id;
    private String contenido;
    private InvitacionResponseDto invitacion;

    // Constructor por defecto
    public NotificacionResponseDto() {
    }

    // Getter y Setter para id
    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    // Getter y Setter para contenido
    public String getContenido() {
        return contenido;
    }

    public void setContenido(String contenido) {
        this.contenido = contenido;
    }

    // Getter y Setter para invitacion
    public InvitacionResponseDto getInvitacion() {
        return invitacion;
    }

    public void setInvitacion(InvitacionResponseDto invitacion) {
        this.invitacion = invitacion;
    }
}
