package protocolo;

import dto.implementacion.InvitacionCanal.InvitacionCanalRequestDto;
import dto.implementacion.InvitacionCanal.Request.CanalInvitacion;
import dto.implementacion.InvitacionCanal.Request.DestinatarioNotificacion;
import dto.implementacion.InvitacionCanal.Request.NotificacionDto;
import protocolo.implementaciones.InvitacionCanal.InvitacionCanalRequest;

import java.time.LocalDateTime;

public class Main {
    public static void main(String[] args) {
        // Crear los objetos correspondientes para el DTO
        CanalInvitacion canal = new CanalInvitacion();
        canal.setId(12); // Asignamos el ID al canal

        DestinatarioNotificacion destinatario = new DestinatarioNotificacion();
        destinatario.setId(4); // Asignamos el ID al destinatario

        NotificacionDto notificacion = new NotificacionDto();
        notificacion.setContenido("Has sido invitado al usuario 'usuario.id'");

        // Crear el DTO con los valores correspondientes
        InvitacionCanalRequestDto dto = new InvitacionCanalRequestDto();
        dto.setCanal(canal);
        dto.setDestinatario(destinatario);
        dto.setFechaEnvio(LocalDateTime.of(2025, 4, 28, 21, 30, 0, 0)); // Fecha de envío
        dto.setEstado(true); // El estado es null
        dto.setNotificacion(notificacion);

        // Crear el objeto InvitacionCanalRequest a partir del DTO
        InvitacionCanalRequest request = new InvitacionCanalRequest(dto);

        // Obtener el JSON del objeto
        String jsonRequest = request.toJson();

        // Imprimir el JSON
        System.out.println(jsonRequest);
    }
}
