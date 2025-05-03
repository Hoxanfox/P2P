package dto.implementacion.SendMessageUser.Response;

import java.util.List;
import java.util.UUID;

public class ChatResponseDto {
    private UUID id;
    private String tipo;
    private List<MiembroChatDto> miembros;
    private UUID tipoChatId ;

    public ChatResponseDto() {}

    public ChatResponseDto(UUID id, String tipo, List<MiembroChatDto> miembros,UUID tipoChatId) {
        this.id = id;
        this.tipo = tipo;
        this.miembros = miembros;
        this.tipoChatId = tipoChatId;
    }

    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }
    public UUID getTipoChatId() {
        return tipoChatId;
    }

    public void setTipoChatId(UUID tipoChatId) {
        this.tipoChatId = tipoChatId;
    }

    public List<MiembroChatDto> getMiembros() {
        return miembros;
    }

    public void setMiembros(List<MiembroChatDto> miembros) {
        this.miembros = miembros;
    }

    public void setTipo(String tipo) {
        this.tipo = tipo;
    }
    public String getTipo() {
        return tipo;
    }
}
