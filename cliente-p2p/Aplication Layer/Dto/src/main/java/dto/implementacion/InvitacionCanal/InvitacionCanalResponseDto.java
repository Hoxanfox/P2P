package dto.implementacion.InvitacionCanal;

import dto.implementacion.InvitacionCanal.Response.InvitacionResponseDto;
import dto.implementacion.InvitacionCanal.Response.NotificacionResponseDto;

public class InvitacionCanalResponseDto {
    private InvitacionResponseDto invitacion;
    private NotificacionResponseDto notificacion;

    // Constructor por defecto
    public InvitacionCanalResponseDto() {
    }

    // Getter y Setter para invitacion
    public InvitacionResponseDto getInvitacion() {
        return invitacion;
    }

    public void setInvitacion(InvitacionResponseDto invitacion) {
        this.invitacion = invitacion;
    }

    // Getter y Setter para notificacion
    public NotificacionResponseDto getNotificacion() {
        return notificacion;
    }

    public void setNotificacion(NotificacionResponseDto notificacion) {
        this.notificacion = notificacion;
    }
}
