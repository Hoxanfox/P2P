package dto.implementacion.ListUsers;

public class ListUsersRequestDto {
    private String command;

    public ListUsersRequestDto(String command) {
        this.command = command;
    }

    public String getCommand() {
        return command;
    }
}
