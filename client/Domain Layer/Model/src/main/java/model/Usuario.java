package model;

import java.time.LocalDateTime;
import java.util.UUID;

public class Usuario {

    private UUID id;
    private String nombre;
    private String email;
    private String password;
    private String foto;
    private String estado;
    private LocalDateTime fechaRegistro;
    private Chat chat;

    // Constructor vacío
    public Usuario() {
    }

    // Constructor con todos los campos (excepto id si quieres que sea opcional)
    public Usuario(UUID id, String nombre, String email, String password, String foto, String estado, LocalDateTime fechaRegistro) {
        this.id = UUID.fromString(String.valueOf(id));
        this.nombre = nombre;
        this.email = email;
        this.password = password;
        this.foto = foto;
        this.estado = estado;
        this.fechaRegistro = fechaRegistro;
    }

    // Getters y Setters

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

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getFoto() {
        return foto;
    }

    public void setFoto(String foto) {
        this.foto = foto;
    }

    public String getEstado() {
        return estado;
    }

    public void setEstado(String estado) {
        this.estado = estado;
    }

    public LocalDateTime getFechaRegistro() {
        return fechaRegistro;
    }

    public void setFechaRegistro(LocalDateTime fechaRegistro) {
        this.fechaRegistro = fechaRegistro;
    }

    public Chat getChat() {
        return chat;
    }
    public void setChat(Chat chat) {
        this.chat = chat;
    }
}
