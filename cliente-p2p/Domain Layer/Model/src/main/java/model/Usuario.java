package model;

import java.time.LocalDate;
import java.util.UUID;

public class Usuario {
    private UUID id;
    private String nombre;
    private String email;
    private String password;
    private byte[] foto; // Cambiado a byte[] por BLOB
    private String ip;
    private boolean estado; // BOOLEAN en BD/ DATE en BD

    // Constructor vacío
    public Usuario() {}

    // Constructor completo
    public Usuario(UUID id, String nombre, String email, String password, byte[] foto, String ip, boolean estado, LocalDate fechaRegistro) {
        this.id = id;
        this.nombre = nombre;
        this.email = email;
        this.password = password;
        this.foto = foto;
        this.ip = ip;
        this.estado = estado;

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

    public byte[] getFoto() {
        return foto;
    }

    public void setFoto(byte[] foto) {
        this.foto = foto;
    }

    public String getIp() {
        return ip;
    }

    public void setIp(String ip) {
        this.ip = ip;
    }

    public boolean isEstado() {
        return estado;
    }

    public void setEstado(boolean estado) {
        this.estado = estado;
    }


}
