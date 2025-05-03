package repository;

import dao.implementaciones.models.UsuarioDao;
import model.Usuario;

public class UsuarioRepository {

    private final UsuarioDao usuarioDao = new UsuarioDao();

    public String registrarUsuario(String nombre, String email, String password, String foto, String estado) {
        Usuario usuario = new Usuario();
        usuario.setNombre(nombre);
        usuario.setEmail(email);
        usuario.setPassword(password);
        usuario.setFoto(foto.getBytes());
        usuario.setEstado(Boolean.parseBoolean(estado));

        Long id = usuarioDao.guardarUsuario(usuario);

        if (id != null) {
            return "Usuario registrado con ID: " + id;
        } else {
            return "Error al registrar usuario";
        }
    }
}
