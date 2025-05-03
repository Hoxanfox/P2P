package controller.implementaciones;

import dto.implementacion.SendMessageUser.Request.SendMessageUserRequestDto;
import facade.implementaciones.SendMessageUserFacade;
import transport.TransportContext;

import java.sql.SQLException;

public class SendMessageUserController {

    private static final String CLASS_NAME = "SendMessageUserController";

    private final SendMessageUserFacade sendMessageUserFacade;

    public SendMessageUserController(TransportContext transportContext) {
        System.out.println(CLASS_NAME + " -> Constructor: Recibiendo TransportContext y creando fachada");

        // Inicializa la fachada con el contexto recibido
        this.sendMessageUserFacade = new SendMessageUserFacade(transportContext);
    }

    public void enviarMensaje(SendMessageUserRequestDto requestDto) {
        System.out.println(CLASS_NAME + " -> enviarMensaje: Iniciando proceso de envío");

        try {
            sendMessageUserFacade.processMessage(requestDto);
            System.out.println(CLASS_NAME + " -> enviarMensaje: Mensaje enviado y guardado exitosamente");
        } catch (SQLException e) {
            System.err.println(CLASS_NAME + " -> enviarMensaje: Error al guardar mensaje en base de datos");
            e.printStackTrace();
        }
    }
}
