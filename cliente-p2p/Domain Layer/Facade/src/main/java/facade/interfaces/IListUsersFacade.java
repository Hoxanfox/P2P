package facade.interfaces;

import dto.implementacion.ListUsers.ListUsersRequestDto;
import dto.implementacion.ListUsers.ListUsersResponseDto;
import transport.TransportContext;

public interface IListUsersFacade {
    ListUsersResponseDto obtenerUsuariosDesdeServidor(ListUsersRequestDto request, TransportContext context);
}
