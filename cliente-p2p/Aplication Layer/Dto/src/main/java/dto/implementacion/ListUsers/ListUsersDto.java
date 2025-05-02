package dto.implementacion.ListUsers;

import dto.implementacion.ListUsers.UsuarioResponseDTO;

import java.util.ArrayList;
import java.util.List;

public class ListUsersDto {
    private List<UsuarioResponseDTO> usuarios;

    public ListUsersDto() {
        this.usuarios = new ArrayList<>();
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
