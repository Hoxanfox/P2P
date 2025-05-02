package dao.implementaciones.usuario;

import dao.DatabaseConfig;
import dao.interfaces.IUsuarioDao;
import model.Usuario;

import java.sql.*;
import java.util.*;

public class UsuarioDao implements IUsuarioDao {

    @Override
    public void guardar(Usuario usuario) throws Exception {
        String sql = """
            INSERT INTO usuarios (id, nombre, email, password, foto, estado, fecha_registro)
            VALUES (?, ?, ?, ?, ?, ?, ?)
        """;

        try (Connection connection = DatabaseConfig.getConnection();
             PreparedStatement stmt = connection.prepareStatement(sql)) {

            stmt.setObject(1, usuario.getId());
            stmt.setString(2, usuario.getNombre());
            stmt.setString(3, usuario.getEmail());
            stmt.setString(4, usuario.getPassword());
            stmt.setString(5, usuario.getFoto());
            stmt.setString(6, usuario.getEstado());
            stmt.setTimestamp(7, Timestamp.valueOf(usuario.getFechaRegistro()));

            stmt.executeUpdate();
        }
    }

    @Override
    public Usuario buscarPorId(UUID id) throws Exception {
        String sql = "SELECT * FROM usuarios WHERE id = ?";

        try (Connection connection = DatabaseConfig.getConnection();
             PreparedStatement stmt = connection.prepareStatement(sql)) {

            stmt.setObject(1, id);

            try (ResultSet rs = stmt.executeQuery()) {
                if (rs.next()) {
                    return mapUsuario(rs);
                }
            }
        }

        return null;
    }

    @Override
    public List<Usuario> listarTodos() throws Exception {
        String sql = "SELECT * FROM usuarios";
        List<Usuario> usuarios = new ArrayList<>();

        try (Connection connection = DatabaseConfig.getConnection();
             PreparedStatement stmt = connection.prepareStatement(sql);
             ResultSet rs = stmt.executeQuery()) {

            while (rs.next()) {
                usuarios.add(mapUsuario(rs));
            }
        }

        return usuarios;
    }

    @Override
    public void actualizar(Usuario usuario) throws Exception {
        String sql = """
            UPDATE usuarios SET nombre = ?, email = ?, password = ?, foto = ?, estado = ?, fecha_registro = ?
            WHERE id = ?
        """;

        try (Connection connection = DatabaseConfig.getConnection();
             PreparedStatement stmt = connection.prepareStatement(sql)) {

            stmt.setString(1, usuario.getNombre());
            stmt.setString(2, usuario.getEmail());
            stmt.setString(3, usuario.getPassword());
            stmt.setString(4, usuario.getFoto());
            stmt.setString(5, usuario.getEstado());
            stmt.setTimestamp(6, Timestamp.valueOf(usuario.getFechaRegistro()));
            stmt.setObject(7, usuario.getId());

            stmt.executeUpdate();
        }
    }

    @Override
    public void eliminar(UUID id) throws Exception {
        String sql = "DELETE FROM usuarios WHERE id = ?";

        try (Connection connection = DatabaseConfig.getConnection();
             PreparedStatement stmt = connection.prepareStatement(sql)) {

            stmt.setObject(1, id);
            stmt.executeUpdate();
        }
    }

    private Usuario mapUsuario(ResultSet rs) throws SQLException {
        Usuario u = new Usuario();
        u.setId((UUID) rs.getObject("id"));
        u.setNombre(rs.getString("nombre"));
        u.setEmail(rs.getString("email"));
        u.setPassword(rs.getString("password"));
        u.setFoto(rs.getString("foto"));
        u.setEstado(rs.getString("estado"));
        u.setFechaRegistro(rs.getTimestamp("fecha_registro").toLocalDateTime());
        return u;
    }
}
