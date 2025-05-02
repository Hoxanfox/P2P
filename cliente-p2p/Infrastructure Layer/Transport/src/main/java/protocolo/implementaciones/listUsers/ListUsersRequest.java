package protocolo.implementaciones.listUsers;

import dto.implementacion.ListUsers.ListUsersRequestDto;
import protocolo.interfaces.RequestRoute;

public class ListUsersRequest implements RequestRoute {
    private final ListUsersRequestDto dto;

    public ListUsersRequest() {
        this.dto = new ListUsersRequestDto("list-users"); // default
    }

    public ListUsersRequest(ListUsersRequestDto dto) {
        this.dto = dto;
    }

    public ListUsersRequestDto getDto() {
        return dto;
    }

    @Override
    public String toJson() {
        return "{ \"command\": \"" + dto.getCommand() + "\" }";
    }
}
