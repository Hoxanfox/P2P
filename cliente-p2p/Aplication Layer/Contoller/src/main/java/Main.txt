import controller.implementaciones.SendMessageUserController;
import dto.implementacion.SendMessageUser.DestinatarioMessageDto;
import dto.implementacion.SendMessageUser.RemitenteMessageDto;
import dto.implementacion.SendMessageUser.Request.MensajeRequestDto;
import dto.implementacion.SendMessageUser.Request.SendMessageUserRequestDto;
import transport.TcpPersistentTransportStrategy;
import transport.TransportContext;

import java.util.UUID;

public class Main {
    public static void main(String[] args) {
        try {
            System.out.println("[DEBUG] Iniciando flujo de prueba...");

            // Crear la estrategia de transporte TCP persistente
            TcpPersistentTransportStrategy estrategia = new TcpPersistentTransportStrategy("localhost", 9000); // Reemplaza con IP y puerto reales
            TransportContext context = new TransportContext(estrategia);

            String emailAna = "ana.torres@correo.com";
            String emailLuis = "luis.mendoza@correo.com";

            System.out.println("[DEBUG] Enviando mensaje desde " + emailAna + " a " + emailLuis + "...");

            SendMessageUserRequestDto mensaje = new SendMessageUserRequestDto();
            mensaje.setMensaje(new MensajeRequestDto());

            UUID remitenteId = UUID.fromString("210c3aea-d243-4b6c-8456-7bb67ff5306e");
            UUID destinatarioId = UUID.fromString("0db20b34-00a6-48d0-8ebb-49de460a99a4");

            RemitenteMessageDto remitente = new RemitenteMessageDto(remitenteId, emailAna);
            DestinatarioMessageDto destinatario = new DestinatarioMessageDto(destinatarioId, emailLuis);

            mensaje.getMensaje().setRemitente(remitente);
            mensaje.getMensaje().setDestinatario(destinatario);
            mensaje.getMensaje().setContenido("Hola Luis, ¿cómo estás?");
            mensaje.getMensaje().setArchivo(null);

            SendMessageUserController mensajeController = new SendMessageUserController(context);
            mensajeController.enviarMensaje(mensaje);

            System.out.println("[INFO] Mensaje enviado exitosamente.");

        } catch (Exception e) {
            System.err.println("[ERROR GENERAL] " + e.getMessage());
            e.printStackTrace();
        }
    }
}
