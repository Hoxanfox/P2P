package facade.implementaciones;

import dto.implementacion.ListUsers.ListUsersRequestDto;
import dto.implementacion.ListUsers.ListUsersResponseDto;
import protocolo.implementaciones.listUsers.ListUsersRequest;
import protocolo.implementaciones.listUsers.ListUsersResponse;
import transport.TransportContext;
import facade.interfaces.IListUsersFacade;

public class ListUsersFacade implements IListUsersFacade {

    @Override
    public ListUsersResponseDto obtenerUsuariosDesdeServidor(ListUsersRequestDto request, TransportContext context) {
        System.out.println("[ListUsersFacade] Iniciando obtención de usuarios desde el servidor...");

        // 1. Empaquetamos la petición en un objeto protocolo
        System.out.println("[ListUsersFacade] Empaquetando petición...");
        ListUsersRequest protocoloRequest = new ListUsersRequest(request);

        // 2. Serializamos y enviamos por el contexto
        String jsonRequest = protocoloRequest.toJson();
        System.out.println("[ListUsersFacade] JSON enviado: " + jsonRequest);

        String jsonResponse = context.executeSend(jsonRequest);
        System.out.println("[ListUsersFacade] JSON recibido: " + jsonResponse);

        // 3. Procesamos la respuesta
        System.out.println("[ListUsersFacade] Procesando respuesta...");
        ListUsersResponse protocoloResponse = new ListUsersResponse();
        protocoloResponse.fromJson(jsonResponse);

        // 4. Retornamos el DTO de respuesta
        ListUsersResponseDto responseDto = protocoloResponse.getDto();
        System.out.println("[ListUsersFacade] DTO de respuesta generado con " +
                (responseDto.getUsuarios() != null ? responseDto.getUsuarios().size() : 0) + " usuarios.");

        return responseDto;
    }
}
