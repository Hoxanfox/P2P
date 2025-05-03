package dto.implementacion.SendMessageUser.Request;

public class SendMessageUserRequestDto {
    private MensajeRequestDto mensaje;

    // Getters y setters
    public MensajeRequestDto getMensaje() {
        return mensaje;
    }

    public void setMensaje(MensajeRequestDto mensaje) {
        this.mensaje = mensaje;
    }
}
