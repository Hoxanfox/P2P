package dto.implementacion.SendMessageUser.Response;

import dto.implementacion.SendMessageUser.Response.ChatResponseDto;
import dto.implementacion.SendMessageUser.DestinatarioMessageDto;
import dto.implementacion.SendMessageUser.RemitenteMessageDto;

import java.util.UUID;

public class MensajeResponseDto {
    private UUID id;
    private RemitenteMessageDto remitente;
    private DestinatarioMessageDto destinatario;
    private String contenido;
    private String fechaEnvio; // Formato ISO-8601: "2025-04-28T14:30:00"

    /**
     * Puede ser null o una cadena en Base64 si se incluye archivo adjunto.
     */
    private String archivo;

    private ChatResponseDto chat;

    // Constructor vacío
    public MensajeResponseDto() {}

    // Constructor completo (opcional)
    public MensajeResponseDto(RemitenteMessageDto remitente, DestinatarioMessageDto destinatario, String contenido, String fechaEnvio, String archivo, ChatResponseDto chat) {
        this.remitente = remitente;
        this.destinatario = destinatario;
        this.contenido = contenido;
        this.fechaEnvio = fechaEnvio;
        this.archivo = archivo;
        this.chat = chat;
    }

    // Getters y setters

    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }
    public RemitenteMessageDto getRemitente() {
        return remitente;
    }

    public void setRemitente(RemitenteMessageDto remitente) {
        this.remitente = remitente;
    }

    public DestinatarioMessageDto getDestinatario() {
        return destinatario;
    }

    public void setDestinatario(DestinatarioMessageDto destinatario) {
        this.destinatario = destinatario;
    }

    public String getContenido() {
        return contenido;
    }

    public void setContenido(String contenido) {
        this.contenido = contenido;
    }

    public String getFechaEnvio() {
        return fechaEnvio;
    }

    public void setFechaEnvio(String fechaEnvio) {
        this.fechaEnvio = fechaEnvio;
    }

    public String getArchivo() {
        return archivo;
    }

    public void setArchivo(String archivo) {
        this.archivo = archivo;
    }

    public ChatResponseDto getChat() {
        return chat;
    }

    public void setChat(ChatResponseDto chat) {
        this.chat = chat;
    }
}
