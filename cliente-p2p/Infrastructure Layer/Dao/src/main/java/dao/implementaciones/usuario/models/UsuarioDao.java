package dao.implementaciones.usuario.models;

import dao.DatabaseConfig;
import dao.interfaces.IUsuarioDao;
import model.Usuario;

import java.sql.*;
import java.time.Instant;
import java.time.ZoneId;
import java.util.*;

public class UsuarioDao implements IUsuarioDao {

    @Override
    public void guardar(Usuario usuario) throws Exception {
        String sql = """
            INSERT INTO UsuariosServidor(
                id_usuario_servidor, nombre, email, password, foto, ip, fecha_registro, estado
            ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
        """;

        try (Connection conn = DatabaseConfig.getConnection();
             PreparedStatement stmt = conn.prepareStatement(sql)) {

            System.out.println("[DEBUG] Guardando usuario: " + usuario);

            stmt.setObject(1, usuario.getId());
            stmt.setString(2, usuario.getNombre());
            stmt.setString(3, usuario.getEmail());
            stmt.setString(4, usuario.getPassword());

            if (usuario.getFoto() != null) {
                stmt.setBytes(5, usuario.getFoto());
            } else {
                stmt.setNull(5, Types.BLOB);
            }

            stmt.setString(6, usuario.getIp());
            // → Store date as ISO text

            stmt.setBoolean(8, usuario.isEstado());

            int rows = stmt.executeUpdate();
            System.out.println("[DEBUG] Filas insertadas: " + rows);

        } catch (SQLException e) {
            System.err.println("[ERROR] Error al guardar usuario: " + e.getMessage());
            throw new Exception("Error al guardar el usuario: " + e.getMessage(), e);
        }
    }

    @Override
    public Usuario buscarPorId(UUID id) throws Exception {
        String sql = "SELECT * FROM UsuariosServidor WHERE id_usuario_servidor = ?";

        try (Connection conn = DatabaseConfig.getConnection();
             PreparedStatement stmt = conn.prepareStatement(sql)) {

            System.out.println("[DEBUG] Buscando usuario por ID: " + id);
            stmt.setObject(1, id);

            try (ResultSet rs = stmt.executeQuery()) {
                if (rs.next()) {
                    Usuario usuario = mapUsuario(rs);
                    System.out.println("[DEBUG] Usuario encontrado: " + usuario);
                    return usuario;
                } else {
                    System.out.println("[DEBUG] No se encontró usuario con ID: " + id);
                }
            }
        } catch (SQLException e) {
            System.err.println("[ERROR] Error al buscar usuario por ID: " + e.getMessage());
            throw new Exception("Error al buscar usuario por ID: " + e.getMessage(), e);
        }
        return null;
    }

    @Override
    public List<Usuario> listarTodos() throws Exception {
        String sql = "SELECT * FROM UsuariosServidor";
        List<Usuario> usuarios = new ArrayList<>();

        try (Connection conn = DatabaseConfig.getConnection();
             PreparedStatement stmt = conn.prepareStatement(sql);
             ResultSet rs = stmt.executeQuery()) {

            System.out.println("[DEBUG] Listando todos los usuarios");
            while (rs.next()) {
                usuarios.add(mapUsuario(rs));
            }
            System.out.println("[DEBUG] Total usuarios listados: " + usuarios.size());

        } catch (SQLException e) {
            System.err.println("[ERROR] Error al listar usuarios: " + e.getMessage());
            throw new Exception("Error al listar usuarios: " + e.getMessage(), e);
        }
        return usuarios;
    }

    @Override
    public void actualizar(Usuario usuario) throws Exception {
        String sql = """
            UPDATE UsuariosServidor
            SET nombre = ?, email = ?, password = ?, foto = ?, ip = ?, fecha_registro = ?, estado = ?
            WHERE id_usuario_servidor = ?
        """;

        try (Connection conn = DatabaseConfig.getConnection();
             PreparedStatement stmt = conn.prepareStatement(sql)) {

            System.out.println("[DEBUG] Actualizando usuario: " + usuario);

            stmt.setString(1, usuario.getNombre());
            stmt.setString(2, usuario.getEmail());
            stmt.setString(3, usuario.getPassword());

            if (usuario.getFoto() != null) {
                stmt.setBytes(4, usuario.getFoto());
            } else {
                stmt.setNull(4, Types.BLOB);
            }

            stmt.setString(5, usuario.getIp());
            // → Store date as ISO text

            stmt.setBoolean(7, usuario.isEstado());
            stmt.setObject(8, usuario.getId());

            int rows = stmt.executeUpdate();
            System.out.println("[DEBUG] Filas actualizadas: " + rows);

        } catch (SQLException e) {
            System.err.println("[ERROR] Error al actualizar usuario: " + e.getMessage());
            throw new Exception("Error al actualizar usuario: " + e.getMessage(), e);
        }
    }

    @Override
    public void eliminar(UUID id) throws Exception {
        String sql = "DELETE FROM UsuariosServidor WHERE id_usuario_servidor = ?";

        try (Connection conn = DatabaseConfig.getConnection();
             PreparedStatement stmt = conn.prepareStatement(sql)) {

            System.out.println("[DEBUG] Eliminando usuario con ID: " + id);
            stmt.setObject(1, id);

            int rows = stmt.executeUpdate();
            System.out.println("[DEBUG] Filas eliminadas: " + rows);

        } catch (SQLException e) {
            System.err.println("[ERROR] Error al eliminar usuario: " + e.getMessage());
            throw new Exception("Error al eliminar usuario: " + e.getMessage(), e);
        }
    }

    private Usuario mapUsuario(ResultSet rs) throws SQLException {
        Usuario u = new Usuario();

        u.setId(UUID.fromString(rs.getString("id_usuario_servidor")));
        u.setNombre(rs.getString("nombre"));
        u.setEmail(rs.getString("email"));
        u.setPassword(rs.getString("password"));
        u.setFoto(rs.getBytes("foto"));
        u.setIp(rs.getString("ip"));

        // Si la fecha es un timestamp en milisegundos
        long timestamp = rs.getLong("fecha_registro");


        u.setEstado(rs.getBoolean("estado"));

        System.out.println("[DEBUG] Mapeando usuario desde ResultSet: " + u);
        return u;
    }

}
