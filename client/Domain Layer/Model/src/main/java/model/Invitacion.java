package model;

import java.time.LocalDateTime;

public class Invitacion {

    private Long id;
    private Usuario destinatario; // Usuario destinatario
    private LocalDateTime fechaEnvio;
    private String estado;

    // Constructor vacío
    public Invitacion() {
    }

    // Constructor con todos los campos
    public Invitacion(Long id, Usuario destinatario, LocalDateTime fechaEnvio, String estado) {
        this.id = id;
        this.destinatario = destinatario;
        this.fechaEnvio = fechaEnvio;
        this.estado = estado;
    }

    // Getters y Setters

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Usuario getDestinatario() {
        return destinatario;
    }

    public void setDestinatario(Usuario destinatario) {
        this.destinatario = destinatario;
    }

    public LocalDateTime getFechaEnvio() {
        return fechaEnvio;
    }

    public void setFechaEnvio(LocalDateTime fechaEnvio) {
        this.fechaEnvio = fechaEnvio;
    }

    public String getEstado() {
        return estado;
    }

    public void setEstado(String estado) {
        this.estado = estado;
    }
}
