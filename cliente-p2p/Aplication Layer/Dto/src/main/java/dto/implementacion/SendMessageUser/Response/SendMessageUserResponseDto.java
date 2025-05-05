package dto.implementacion.SendMessageUser.Response;

import dto.implementacion.SendMessageUser.DestinatarioMessageDto;
import dto.implementacion.SendMessageUser.RemitenteMessageDto;

import java.util.UUID;

public class SendMessageUserResponseDto {

    private UUID id;
    private RemitenteMessageDto remitente;
    private DestinatarioMessageDto destinatario;
    private String contenido;
    private String fechaEnvio;
    private String archivo;
    private ChatResponseDto chat;
   // Este puede ser null si no hubo error

    // Getters and setters

    public UUID getId() { return id; }
    public void setId(UUID id) { this.id = id; }

    public RemitenteMessageDto getRemitente() { return remitente; }
    public void setRemitente(RemitenteMessageDto remitente) { this.remitente = remitente; }

    public DestinatarioMessageDto getDestinatario() { return destinatario; }
    public void setDestinatario(DestinatarioMessageDto destinatario) { this.destinatario = destinatario; }

    public String getContenido() { return contenido; }
    public void setContenido(String contenido) { this.contenido = contenido; }

    public String getFechaEnvio() { return fechaEnvio; }
    public void setFechaEnvio(String fechaEnvio) { this.fechaEnvio = fechaEnvio; }

    public String getArchivo() { return archivo; }
    public void setArchivo(String archivo) { this.archivo = archivo; }

    public ChatResponseDto getChat() { return chat; }
    public void setChat(ChatResponseDto chat) { this.chat = chat; }

}
