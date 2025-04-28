package model;

public class Notificacion {

    private Long id;
    private Invitacion invitacion;
    private String contenido; // Mensaje o contenido de la notificación

    // Constructor vacío
    public Notificacion() {
    }

    // Constructor con todos los campos
    public Notificacion(Long id, Invitacion invitacion, String contenido) {
        this.id = id;
        this.invitacion = invitacion;
        this.contenido = contenido;
    }

    // Getters y Setters

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Invitacion getInvitacion() {
        return invitacion;
    }

    public void setInvitacion(Invitacion invitacion) {
        this.invitacion = invitacion;
    }

    public String getContenido() {
        return contenido;
    }

    public void setContenido(String contenido) {
        this.contenido = contenido;
    }
}
