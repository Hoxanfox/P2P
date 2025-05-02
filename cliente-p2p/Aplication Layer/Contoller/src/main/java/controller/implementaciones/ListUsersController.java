package controller.implementaciones;

import dto.implementacion.ListUsers.ListUsersRequestDto;
import dto.implementacion.ListUsers.ListUsersResponseDto;
import facade.interfaces.IListUsersFacade;
import transport.TransportContext;

public class ListUsersController {

    private final IListUsersFacade listUsersFacade;
    private final TransportContext transportContext;

    public ListUsersController(IListUsersFacade listUsersFacade, TransportContext transportContext) {
        this.listUsersFacade = listUsersFacade;
        this.transportContext = transportContext;
    }

    public ListUsersResponseDto obtenerUsuarios() {
        // Creamos el DTO de la petición
        ListUsersRequestDto requestDto = new ListUsersRequestDto("list-users");

        // Utilizamos la fachada para obtener la respuesta
        return listUsersFacade.obtenerUsuariosDesdeServidor(requestDto, transportContext);
    }
}
