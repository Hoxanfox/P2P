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
        // 1. Empaquetamos la petición en un objeto protocolo
        ListUsersRequest protocoloRequest = new ListUsersRequest(request);

        // 2. Serializamos y enviamos por el contexto
        String jsonRequest = protocoloRequest.toJson();
        String jsonResponse = context.executeSend(jsonRequest);

        // 3. Procesamos la respuesta
        ListUsersResponse protocoloResponse = new ListUsersResponse();
        protocoloResponse.fromJson(jsonResponse);

        // 4. Retornamos el DTO de respuesta
        return protocoloResponse.getDto();
    }
}
