package dao.implementaciones.register;

import dao.DatabaseConfig;
import dto.implementacion.register.RegisterResponseDto;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;

public class RegisterDao {

    /**
     * Inserta un nuevo usuario en la tabla UsuariosServidor usando la conexión
     * gestionada por DatabaseConfig.
     *
     * @param usuario DTO con los datos del usuario a registrar
     * @return true si se insertó exactamente una fila; false en caso contrario
     * @throws SQLException si ocurre un error al ejecutar el INSERT
     */
    public boolean registrarUsuario(RegisterResponseDto usuario) throws SQLException {
        String sql = """
            INSERT INTO UsuariosServidor (
              id_usuario_servidor,
              nombre,
              email,
              password,
              foto,
              ip,
              estado
            ) VALUES (?, ?, ?, ?, ?, ?,?)
            """;

        // Obtenemos la conexión singleton de DatabaseConfig
        try (Connection conn = DatabaseConfig.getConnection();
             PreparedStatement stmt = conn.prepareStatement(sql)) {

            stmt.setString(1, usuario.getId().toString());
            stmt.setString(2, usuario.getNombre());
            stmt.setString(3, usuario.getEmail());
            stmt.setString(4, usuario.getPassword());
            stmt.setBytes(5, usuario.getFoto()); // byte[] -> BLOB
            stmt.setString(6, usuario.getIp());


            // Boolean -> entero 0/1
            stmt.setInt(7, usuario.isEstado() ? 1 : 0);

            int rowsInserted = stmt.executeUpdate();
            return rowsInserted == 1;
        }
    }
}
