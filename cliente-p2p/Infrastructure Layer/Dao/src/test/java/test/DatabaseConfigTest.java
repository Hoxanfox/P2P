package test;

import dao.DatabaseConfig;

import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

public class DatabaseConfigTest {

    public static void main(String[] args) {
        try {
            System.out.println("[DEBUG] Obteniendo conexión...");
            Connection conn = DatabaseConfig.getConnection();

            System.out.println("[DEBUG] Conexión establecida: " + (conn != null));

            Statement stmt = conn.createStatement();

            System.out.println("[DEBUG] Creando tabla de prueba...");
            stmt.execute("CREATE TABLE IF NOT EXISTS UsuariosServidor2 (" +
                    "id_usuario_servidor TEXT PRIMARY KEY, " +
                    "nombre TEXT, " +
                    "email TEXT, " +
                    "password TEXT, " +
                    "foto BLOB, " +
                    "ip TEXT, " +
                    "fecha_registro DATE, " +
                    "estado BOOLEAN)");

            System.out.println("[DEBUG] Tabla creada (o ya existía).");

            System.out.println("[DEBUG] Insertando registro de prueba...");
            stmt.executeUpdate("INSERT INTO UsuariosServidor2 (id_usuario_servidor, nombre, email, password, ip, fecha_registro, estado) " +
                    "VALUES ('123e4567-e89b-12d3-a456-426614174000', 'Juan Perez', 'juan@example.com', '1234', '127.0.0.1', DATE('now'), 1)");

            System.out.println("[DEBUG] Registro insertado.");

            System.out.println("[DEBUG] Consultando registros...");
            ResultSet rs = stmt.executeQuery("SELECT * FROM UsuariosServidor");

            while (rs.next()) {
                System.out.println("[RESULTADO] Usuario: " + rs.getString("nombre") + ", Email: " + rs.getString("email"));
            }

            conn.close();
            System.out.println("[DEBUG] Conexión cerrada correctamente.");

        } catch (SQLException e) {
            System.err.println("[ERROR] Error al conectar o ejecutar sentencia SQL: " + e.getMessage());
            e.printStackTrace();
        }
    }
}
