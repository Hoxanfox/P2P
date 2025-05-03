package dto.implementacion.CreateChannel;

import dto.implementacion.CreateChannel.Channel.MiembroCanalDto;
import dto.implementacion.CreateChannel.Channel.InvitacionDto;
import java.util.List;

public class CreateChannelRequestDto {

    private String nombre;
    private String descripcion;
    private List<MiembroCanalDto> miembros;
    private List<InvitacionDto> invitaciones;
    private String tipo;

    // Getters y Setters
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

    public List<MiembroCanalDto> getMiembros() {
        return miembros;
    }

    public void setMiembros(List<MiembroCanalDto> miembros) {
        this.miembros = miembros;
    }

    public List<InvitacionDto> getInvitaciones() {
        return invitaciones;
    }

    public void setInvitaciones(List<InvitacionDto> invitaciones) {
        this.invitaciones = invitaciones;
    }
    public String getTipo() {
        return tipo;
    }

    public void setTipo(String tipo) {
        this.tipo = tipo;
    }

}
