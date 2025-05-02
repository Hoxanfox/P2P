package dao.interfaces;

import model.Usuario;
import java.util.List;
import java.util.UUID;

public interface IUsuarioDao {
    void guardar(Usuario usuario) throws Exception;
    Usuario buscarPorId(UUID id) throws Exception;
    List<Usuario> listarTodos() throws Exception;
    void actualizar(Usuario usuario) throws Exception;
    void eliminar(UUID id) throws Exception;
}
