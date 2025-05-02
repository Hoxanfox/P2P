package model;

import java.time.LocalDateTime;

public class Mensaje {

    private Long id;
    private Usuario remitente;
    private String contenido;
    private LocalDateTime fechaEnvio;
    private Archivo archivo;
    private Chat chat;

    // Constructor vacío
    public Mensaje() {
    }

    // Constructor con todos los campos
    public Mensaje(Long id, Usuario remitente, String contenido, LocalDateTime fechaEnvio, Archivo archivo, Chat chat) {
        this.id = id;
        this.remitente = remitente;
        this.contenido = contenido;
        this.fechaEnvio = fechaEnvio;
        this.archivo = archivo;
        this.chat = chat;
    }

    // Getters y Setters

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Usuario getRemitente() {
        return remitente;
    }

    public void setRemitente(Usuario remitente) {
        this.remitente = remitente;
    }

    public String getContenido() {
        return contenido;
    }

    public void setContenido(String contenido) {
        this.contenido = contenido;
    }

    public LocalDateTime getFechaEnvio() {
        return fechaEnvio;
    }

    public void setFechaEnvio(LocalDateTime fechaEnvio) {
        this.fechaEnvio = fechaEnvio;
    }

    public Archivo getArchivo() {
        return archivo;
    }

    public void setArchivo(Archivo archivo) {
        this.archivo = archivo;
    }

    public Chat getChat() {
        return chat;
    }

    public void setChat(Chat chat) {
        this.chat = chat;
    }
}
