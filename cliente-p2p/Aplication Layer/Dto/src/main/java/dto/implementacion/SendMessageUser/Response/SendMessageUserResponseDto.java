package dto.implementacion.SendMessageUser.Response;

public class SendMessageUserResponseDto {
    private MensajeResponseDto mensaje; // null si es error// solo usado si hay error

    // Getters y setters


    public MensajeResponseDto getMensaje() {
        return mensaje;
    }

    public void setMensaje(MensajeResponseDto mensaje) {
        this.mensaje = mensaje;
    }

}
