package dao;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.nio.charset.StandardCharsets;
import java.sql.*;
import java.util.stream.Collectors;

public class DatabaseInitializer {

    // URL de SQLite, el archivo se crea en ./data/mydatabase.db
    private static final String JDBC_URL = "jdbc:sqlite:./data/mydatabase.db";

    public static void main(String[] args) {
        try (Connection conn = DriverManager.getConnection(JDBC_URL)) {
            // Verifica si la tabla ya existe
            if (!tableExists(conn, "UsuariosServidor")) {
                System.out.println("Tabla UsuariosServidor no encontrada, inicializando esquema...");
                executeSchema(conn);
                System.out.println("Base de datos inicializada correctamente.");
            } else {
                System.out.println("La tabla ya existe. No se requiere inicialización.");
            }
        } catch (Exception e) {
            // Muestra el error si ocurre alguna excepción
            System.err.println("Error durante la inicialización de la base de datos:");
            e.printStackTrace();
        }
    }

    /** Comprueba si la tabla existe en la base de datos */
    private static boolean tableExists(Connection conn, String tableName) throws SQLException {
        String sql = "SELECT name FROM sqlite_master WHERE type='table' AND name=?";
        try (PreparedStatement stmt = conn.prepareStatement(sql)) {
            stmt.setString(1, tableName);
            try (ResultSet rs = stmt.executeQuery()) {
                return rs.next();
            }
        }
    }

    /** Ejecuta el archivo schema.sql para crear las tablas */
    private static void executeSchema(Connection conn) throws Exception {
        // Carga todo el contenido de schema.sql desde recursos
        String sql;
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(
                DatabaseInitializer.class.getClassLoader().getResourceAsStream("schema.sql"),
                StandardCharsets.UTF_8))) {
            sql = reader.lines().collect(Collectors.joining("\n"));
        }

        // Ejecuta cada sentencia SQL separada por ';'
        try (Statement stmt = conn.createStatement()) {
            for (String statement : sql.split(";")) {
                String trimmed = statement.trim();
                if (!trimmed.isEmpty()) {
                    stmt.execute(trimmed);
                }
            }
        }
    }
}
