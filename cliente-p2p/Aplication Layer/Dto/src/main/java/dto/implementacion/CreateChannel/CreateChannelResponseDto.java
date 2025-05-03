package dto.implementacion.CreateChannel;

import dto.implementacion.CreateChannel.Channel.ChatDto;
import dto.implementacion.CreateChannel.Channel.MiembroCanalDto;
import dto.implementacion.CreateChannel.Channel.InvitacionDto;
import java.util.List;
import java.util.UUID;

public class CreateChannelResponseDto {

    private UUID id;
    private String nombre;
    private String descripcion;
    private List<MiembroCanalDto> miembros;
    private List<InvitacionDto> invitaciones;
    private ChatDto chat;

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

    public ChatDto getChat() {
        return chat;
    }

    public void setChat(ChatDto chat) {
        this.chat = chat;
    }
}
