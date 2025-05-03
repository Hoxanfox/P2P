package dto.implementacion.CreateChannel.Channel;

import java.util.List;
import java.util.UUID;

public class ChatDto {
    private UUID id;
    private String tipo;
    private List<MiembroCanalDto> miembros;

    public ChatDto() {}

    public ChatDto(UUID id, String tipo, List<MiembroCanalDto> miembros) {
        this.id = id;
        this.tipo = tipo;
        this.miembros = miembros;
    }

    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }


    public List<MiembroCanalDto> getMiembros() {
        return miembros;
    }

    public void setMiembros(List<MiembroCanalDto> miembros) {
        this.miembros = miembros;
    }

    public void setTipo(String tipo) {
        this.tipo = tipo;
    }
    public String getTipo() {
        return tipo;
    }
}
