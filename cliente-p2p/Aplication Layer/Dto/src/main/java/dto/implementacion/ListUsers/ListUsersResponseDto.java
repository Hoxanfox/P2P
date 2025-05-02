package dto.implementacion.ListUsers;

import java.util.ArrayList;
import java.util.List;

public class ListUsersResponseDto {
    private String status;
    private String message;
    private List<UsuarioResponseDTO> usuarios = new ArrayList<>();

    public ListUsersResponseDto() {}

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public List<UsuarioResponseDTO> getUsuarios() {
        return usuarios;
    }

    public void setUsuarios(List<UsuarioResponseDTO> usuarios) {
        this.usuarios = usuarios;
    }

    public void addUsuario(UsuarioResponseDTO usuario) {
        this.usuarios.add(usuario);
    }
}
