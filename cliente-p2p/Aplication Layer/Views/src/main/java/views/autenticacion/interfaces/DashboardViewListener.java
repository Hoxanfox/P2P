package views.autenticacion.interfaces;

import dto.implementacion.ListUsers.UsuarioResponseDTO;

import java.util.List;

public interface DashboardViewListener {
    void onUsuariosCargados(List<UsuarioResponseDTO> usuarios);
    void onErrorCargarUsuarios(String mensaje);
}
