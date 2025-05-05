package facade.implementaciones;

import dao.implementaciones.invitacionCanal.InvitacionCanalDao;  // Importar el DAO
import dto.implementacion.InvitacionCanal.InvitacionCanalRequestDto;
import dto.implementacion.InvitacionCanal.InvitacionCanalResponseDto;
import facade.interfaces.IInvitacionCanalFacade;
import protocolo.implementaciones.InvitacionCanal.InvitacionCanalRequest;
import protocolo.implementaciones.InvitacionCanal.InvitacionCanalResponse;
import transport.TransportContext;

public class InvitacionCanalFacade implements IInvitacionCanalFacade {

    private final InvitacionCanalDao invitacionCanalDao = new InvitacionCanalDao(); // Crear una instancia del DAO

    @Override
    public InvitacionCanalResponseDto obtenerInformacionDesdeServidor(InvitacionCanalRequestDto requestDto, TransportContext context) {
        try {
            // Crear objeto de protocolo y convertirlo a JSON
            InvitacionCanalRequest request = new InvitacionCanalRequest(requestDto);
            String jsonToSend = request.toJson();

            // Enviar JSON al servidor y obtener respuesta
            String jsonResponse = context.executeSend(jsonToSend);

            // Parsear respuesta del servidor
            InvitacionCanalResponse response = new InvitacionCanalResponse();
            response.fromJson(jsonResponse);

            return response.getData(); // DTO con la información deserializada
        } catch (Exception e) {
            System.err.println("[InvitacionCanalFacade] Error al obtener información desde el servidor: " + e.getMessage());
            throw new RuntimeException("Fallo al ejecutar la solicitud al servidor", e);
        }
    }

    @Override
    public void persistirInformacion(InvitacionCanalResponseDto responseDto) {
        // Usamos el DAO para persistir la invitación y la notificación
        try {
            System.out.println("[InvitacionCanalFacade] Persistiendo información...");

            // Aquí pasamos los datos necesarios al DAO para que persista
            invitacionCanalDao.guardarInvitacionYNotificacion(
                    responseDto.getInvitacion().getDestinatario().getId().toString(),
                    responseDto.getInvitacion().getCanal().getId().toString(),
                    responseDto.getNotificacion().getContenido()
            );

            System.out.println("[InvitacionCanalFacade] Invitación ID: " + responseDto.getInvitacion().getId());
            System.out.println("[InvitacionCanalFacade] Notificación ID: " + responseDto.getNotificacion().getId());

        } catch (Exception e) {
            System.err.println("[InvitacionCanalFacade] Error al persistir la información: " + e.getMessage());
            throw new RuntimeException("Error al persistir la invitación y notificación", e);
        }
    }

    @Override
    public void ejecutarFlujo(InvitacionCanalRequestDto requestDto, TransportContext context) {
        System.out.println("[InvitacionCanalFacade] === INICIO FLUJO COMPLETO ===");
        InvitacionCanalResponseDto responseDto = obtenerInformacionDesdeServidor(requestDto, context);
        persistirInformacion(responseDto);
        System.out.println("[InvitacionCanalFacade] === FIN FLUJO COMPLETO ===");
    }
}
