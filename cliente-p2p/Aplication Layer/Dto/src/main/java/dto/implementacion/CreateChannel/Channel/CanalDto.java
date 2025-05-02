package dto.implementacion.CreateChannel.Channel;

import dto.implementacion.CreateChannel.Channel.ChatDto;
import dto.implementacion.CreateChannel.Channel.InvitacionDto;
import dto.implementacion.CreateChannel.Channel.MiembroCanalDto;

import java.util.List;

public class CanalDto {
    private Long id;
    private String nombre;
    private String descripcion;
    private List<InvitacionDto> invitaciones;
    private List<MiembroCanalDto> miembros;
    private ChatDto chat;

    public CanalDto() {}

    public CanalDto(Long id, String nombre, String descripcion, List<InvitacionDto> invitaciones, List<MiembroCanalDto> miembros, ChatDto chat) {
        this.id = id;
        this.nombre = nombre;
        this.descripcion = descripcion;
        this.invitaciones = invitaciones;
        this.miembros = miembros;
        this.chat = chat;
    }

    // Getters y Setters
    // (Genera todos)
}
