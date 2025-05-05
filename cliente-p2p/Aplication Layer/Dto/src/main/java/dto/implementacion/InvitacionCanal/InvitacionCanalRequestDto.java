package dto.implementacion.InvitacionCanal;

import dto.implementacion.InvitacionCanal.Request.CanalInvitacion;
import dto.implementacion.InvitacionCanal.Request.DestinatarioNotificacion;
import dto.implementacion.InvitacionCanal.Request.NotificacionDto;

import java.time.LocalDateTime;

public class InvitacionCanalRequestDto {

    private CanalInvitacion canal;
    private DestinatarioNotificacion destinatario;
    private LocalDateTime fechaEnvio;
    private boolean estado;
    private NotificacionDto notificacion;

    // Getters y setters
    public CanalInvitacion getCanal() {
        return canal;
    }

    public void setCanal(CanalInvitacion canal) {
        this.canal = canal;
    }

    public DestinatarioNotificacion getDestinatario() {
        return destinatario;
    }

    public void setDestinatario(DestinatarioNotificacion destinatario) {
        this.destinatario = destinatario;
    }

    public LocalDateTime getFechaEnvio() {
        return fechaEnvio;
    }

    public void setFechaEnvio(LocalDateTime fechaEnvio) {
        this.fechaEnvio = fechaEnvio;
    }

    public boolean getEstado() {
        return estado;
    }

    public void setEstado(boolean estado) {
        this.estado = estado;
    }

    public NotificacionDto getNotificacion() {
        return notificacion;
    }

    public void setNotificacion(NotificacionDto notificacion) {
        this.notificacion = notificacion;
    }

}
