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
        TcpPersistentTransportStrategy estrategia = null;

        try {
            System.out.println("[DEBUG] Iniciando flujo de prueba...");

            // Crear estrategia de transporte
            estrategia = new TcpPersistentTransportStrategy("localhost", 9000);
            TransportContext context = new TransportContext(estrategia);

            // Crear los datos del mensaje
            String emailAna = "juan@example.com";
            String emailLuis = "juan@example.es";

            UUID remitenteId = UUID.fromString("123e4567-e89b-12d3-a456-426614174000");
            UUID destinatarioId = UUID.fromString("0fd3d585-4144-476f-b4ff-d4cf643671d7");

            RemitenteMessageDto remitente = new RemitenteMessageDto(remitenteId, emailAna);
            DestinatarioMessageDto destinatario = new DestinatarioMessageDto(destinatarioId, emailLuis);

            MensajeRequestDto mensajeDto = new MensajeRequestDto();
            mensajeDto.setRemitente(remitente);
            mensajeDto.setDestinatario(destinatario);
            mensajeDto.setContenido("Hola Luis, ¿cómo estás?");
            mensajeDto.setArchivo(null); // Si no hay archivo, puede ir null

            SendMessageUserRequestDto mensaje = new SendMessageUserRequestDto();
            mensaje.setMensaje(mensajeDto);

            // Enviar el mensaje
            SendMessageUserController mensajeController = new SendMessageUserController(context);
            mensajeController.enviarMensaje(mensaje);

            System.out.println("[INFO] Mensaje enviado exitosamente.");

        } catch (Exception e) {
            System.err.println("[ERROR GENERAL] " + e.getMessage());
            e.printStackTrace();
        } finally {
            // Cerrar la conexión TCP
            if (estrategia != null) {
                estrategia.close();
                System.out.println("[DEBUG] Estrategia de transporte cerrada correctamente.");
            }
        }
    }
}
