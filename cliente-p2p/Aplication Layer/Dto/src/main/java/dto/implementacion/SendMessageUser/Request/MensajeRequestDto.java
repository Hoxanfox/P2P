package dto.implementacion.SendMessageUser.Request;

import dto.implementacion.SendMessageUser.DestinatarioMessageDto;
import dto.implementacion.SendMessageUser.RemitenteMessageDto;

public class MensajeRequestDto {
    private RemitenteMessageDto remitente;
    private DestinatarioMessageDto destinatario;
    private String contenido;
    private String fechaEnvio; // ISO-8601: "2025-04-28T14:30:00"
    private String archivo;
    // Puede ser null o base64

    // Getters y setters
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
}
