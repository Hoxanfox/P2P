package dto.implementacion.CreateChannel.Channel;

import java.util.List;

public class ChatDto {
    private Long id;
    private String tipo;
    private List<MiembroCanalDto> miembros;

    public ChatDto() {}

    public ChatDto(Long id, String tipo, List<MiembroCanalDto> miembros) {
        this.id = id;
        this.tipo = tipo;
        this.miembros = miembros;
    }

    // Getters y Setters
}
