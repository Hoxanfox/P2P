package model;

import java.util.List;

public class Chat {

    private Long id;
    private String tipo; // "privado" o "publico"
    private List<Usuario> miembros; // Lista de miembros del chat
    private Mensaje mensajes; // Objeto de mensajes (puede tener varios mensajes dentro)

    // Constructor vacío
    public Chat() {
    }

    // Constructor con todos los campos
    public Chat(Long id, String tipo, List<Usuario> miembros, Mensaje mensajes) {
        this.id = id;
        this.tipo = tipo;
        this.miembros = miembros;
        this.mensajes = mensajes;
    }

    // Getters y Setters

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getTipo() {
        return tipo;
    }

    public void setTipo(String tipo) {
        this.tipo = tipo;
    }

    public List<Usuario> getMiembros() {
        return miembros;
    }

    public void setMiembros(List<Usuario> miembros) {
        this.miembros = miembros;
    }

    public Mensaje getMensajes() {
        return mensajes;
    }

    public void setMensajes(Mensaje mensajes) {
        this.mensajes = mensajes;
    }
}
