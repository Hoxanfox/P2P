package controller.implementaciones;

import dto.implementacion.ListUsers.ListUsersRequestDto;
import dto.implementacion.ListUsers.ListUsersResponseDto;
import facade.implementaciones.ListUsersFacade;
import facade.interfaces.IListUsersFacade;
import transport.TransportContext;

public class ListUsersController {

    private final IListUsersFacade listUsersFacade;
    private final TransportContext transportContext;

    public ListUsersController(TransportContext transportContext) {
        System.out.println("[ListUsersController] Constructor iniciado...");
        this.transportContext = transportContext;

        if (this.transportContext != null) {
            System.out.println("[ListUsersController] TransportContext recibido: " + this.transportContext);
        } else {
            System.out.println("[ListUsersController] TransportContext es null");
        }

        this.listUsersFacade = new ListUsersFacade(); // Se instancia internamente
        System.out.println("[ListUsersController] ListUsersFacade instanciado: " + this.listUsersFacade);
    }

    public ListUsersResponseDto obtenerUsuarios() {
        System.out.println("[ListUsersController] Método obtenerUsuarios() invocado.");

        // Creamos el DTO de la petición
        ListUsersRequestDto requestDto = new ListUsersRequestDto("list-users");
        System.out.println("[ListUsersController] ListUsersRequestDto creado con acción: " + requestDto);

        // Utilizamos la fachada para obtener la respuesta
        ListUsersResponseDto response = listUsersFacade.obtenerUsuariosDesdeServidor(requestDto, transportContext);

        if (response != null) {
            System.out.println("[ListUsersController] Respuesta recibida. Cantidad de usuarios: "
                    + (response.getUsuarios() != null ? response.getUsuarios().size() : "null"));
        } else {
            System.out.println("[ListUsersController] Respuesta nula recibida.");
        }

        return response;
    }
}
