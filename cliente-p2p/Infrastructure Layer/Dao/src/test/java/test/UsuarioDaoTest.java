package test;

import dao.DatabaseConfig;
import dao.implementaciones.usuario.models.UsuarioDao;
import model.Usuario;

import java.sql.Connection;
import java.sql.Statement;
import java.time.LocalDate;
import java.util.List;
import java.util.UUID;

public class UsuarioDaoTest {

    public static void main(String[] args) {
        try {
            System.out.println("[DEBUG] Conectando a la base de datos...");
            Connection conn = DatabaseConfig.getConnection();

            // Crear tabla si no existe
            try (Statement stmt = conn.createStatement()) {
                stmt.execute("""
                    CREATE TABLE IF NOT EXISTS UsuariosServidor (
                        id_usuario_servidor TEXT PRIMARY KEY,
                        nombre TEXT,
                        email TEXT,
                        password TEXT,
                        foto BLOB,
                        ip TEXT,
                        fecha_registro DATE,
                        estado BOOLEAN
                    )
                """);
                System.out.println("[DEBUG] Tabla UsuariosServidor verificada o creada.");
            }

            UsuarioDao usuarioDao = new UsuarioDao();

            // Crear usuario de prueba
            UUID id = UUID.randomUUID();
            Usuario usuario = new Usuario();
            usuario.setId(id);
            usuario.setNombre("Juan Pérez");
            usuario.setEmail("juan@example.deivid.esssssss");
            usuario.setPassword("1234");
            usuario.setFoto(null);
            usuario.setIp("127.0.0.1");
            usuario.setEstado(true);

            // Guardar usuario
            usuarioDao.guardar(usuario);

            // Buscar por ID
            Usuario buscado = usuarioDao.buscarPorId(id);
            System.out.println("[RESULTADO] Usuario encontrado: " + buscado);

            // Listar todos
            List<Usuario> usuarios = usuarioDao.listarTodos();
            System.out.println("[RESULTADO] Total usuarios listados: " + usuarios.size());

            // Actualizar
            usuario.setNombre("Juan Actualizado");
            usuario.setEstado(false);
            usuarioDao.actualizar(usuario);

            Usuario actualizado = usuarioDao.buscarPorId(id);
            System.out.println("[RESULTADO] Usuario actualizado: " + actualizado);

            // Eliminar
            usuarioDao.eliminar(id);

            Usuario eliminado = usuarioDao.buscarPorId(id);
            System.out.println("[RESULTADO] Usuario tras eliminación: " + eliminado);

            conn.close();
            System.out.println("[DEBUG] Conexión cerrada.");

        } catch (Exception e) {
            System.err.println("[ERROR] Error durante la prueba: " + e.getMessage());
            e.printStackTrace();
        }
    }
}
