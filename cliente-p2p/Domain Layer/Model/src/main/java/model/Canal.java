package model;

import java.util.ArrayList;
import java.util.List;

public class Canal {

    private Long id;
    private String nombre;
    private String descripcion;
    private List<Invitacion> invitaciones;
    private List<Usuario> miembros; // Lista de usuarios (miembros del canal)
    private Chat chat;

    // Constructor vacío
    public Canal() {
        this.invitaciones = new ArrayList<>();
        this.miembros = new ArrayList<>();
    }

    // Constructor con todos los campos
    public Canal(Long id, String nombre, String descripcion, List<Invitacion> invitaciones, List<Usuario> miembros) {
        this.id = id;
        this.nombre = nombre;
        this.descripcion = descripcion;
        this.invitaciones = invitaciones;
        this.miembros = miembros;
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

    public String getDescripcion() {
        return descripcion;
    }

    public void setDescripcion(String descripcion) {
        this.descripcion = descripcion;
    }

    public List<Invitacion> getInvitaciones() {
        return invitaciones;
    }

    public void setInvitaciones(List<Invitacion> invitaciones) {
        this.invitaciones = invitaciones;
    }

    public List<Usuario> getMiembros() {
        return miembros;
    }

    public void setMiembros(List<Usuario> miembros) {
        this.miembros = miembros;
    }

    public Chat getChat() {
        return chat;
    }

    public void setChat(Chat chat) {
        this.chat = chat;
    }

    // Métodos para agregar/eliminar usuarios (miembros del canal)

    public void agregarMiembro(Usuario usuario) {
        if (usuario != null && !this.miembros.contains(usuario)) {
            this.miembros.add(usuario);
        }
    }

    public void eliminarMiembro(Long idUsuario) {
        this.miembros.removeIf(usuario -> usuario.getId().equals(idUsuario));
    }

    // Métodos para agregar/eliminar invitaciones

    public void agregarInvitacion(Invitacion invitacion) {
        if (invitacion != null) {
            this.invitaciones.add(invitacion);
        }
    }

    public void eliminarInvitacion(Long idInvitacion) {
        this.invitaciones.removeIf(invitacion -> invitacion.getId().equals(idInvitacion));
    }
}
