package model;

import java.time.LocalDateTime;

public class Archivo {

    private Long id;
    private String nombre;
    private byte[] binario;


    // Constructor vacío
    public Archivo() {
    }

    // Constructor con todos los campos
    public Archivo(Long id, String nombre, byte[] binario) {
        this.id = id;
        this.nombre = nombre;
        this.binario = binario;

    }

    // Getters y Setters

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getNombre() {
        return nombre;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }

    public byte[] getBinario() {
        return binario;
    }

    public void setBinario(byte[] binario) {
        this.binario = binario;
    }

}
