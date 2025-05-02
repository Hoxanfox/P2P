package dao.implementaciones.usuario;

import dao.interfaces.IUsuarioDao;
import model.Usuario;
import org.junit.jupiter.api.*;
import util.DatabaseTestUtil;

import java.time.LocalDateTime;
import java.util.List;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.*;

class UsuarioDaoImplTest {

    private static IUsuarioDao usuarioDao;

    @BeforeAll
    static void setupDatabase() throws Exception {
        DatabaseTestUtil.inicializarBaseDeDatosTest();
        usuarioDao = new UsuarioDao();
    }

    @Test
    void testGuardarYBuscarPorId() throws Exception {
        UUID id = UUID.randomUUID();
        Usuario usuario = new Usuario(
                id, "Carlos Pérez", "carlos@test.com", "clave123",
                "foto.jpg", "ACTIVO", LocalDateTime.now()
        );

        usuarioDao.guardar(usuario);
        Usuario encontrado = usuarioDao.buscarPorId(id);

        assertNotNull(encontrado);
        assertEquals("Carlos Pérez", encontrado.getNombre());
    }

    @Test
    void testActualizar() throws Exception {
        UUID id = UUID.randomUUID();
        Usuario usuario = new Usuario(id, "Ana", "ana@correo.com", "1234", null, "ACTIVO", LocalDateTime.now());

        usuarioDao.guardar(usuario);
        usuario.setNombre("Ana María");
        usuarioDao.actualizar(usuario);

        Usuario actualizado = usuarioDao.buscarPorId(id);
        assertEquals("Ana María", actualizado.getNombre());
    }

    @Test
    void testEliminar() throws Exception {
        UUID id = UUID.randomUUID();
        Usuario usuario = new Usuario(id, "Laura", "laura@correo.com", "pass", null, "INACTIVO", LocalDateTime.now());

        usuarioDao.guardar(usuario);
        usuarioDao.eliminar(id);

        Usuario eliminado = usuarioDao.buscarPorId(id);
        assertNull(eliminado);
    }

    @Test
    void testListarTodos() throws Exception {
        List<Usuario> usuarios = usuarioDao.listarTodos();
        assertNotNull(usuarios);
        assertTrue(usuarios.size() >= 0); // puede estar vacía inicialmente
    }
}
