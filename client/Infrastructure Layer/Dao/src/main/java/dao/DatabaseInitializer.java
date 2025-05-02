package dao;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.nio.charset.StandardCharsets;
import java.sql.*;
import java.util.stream.Collectors;

public class DatabaseInitializer {

    private static final String JDBC_URL = "jdbc:h2:./data/mydb"; // o jdbc:h2:mem:testdb
    private static final String USER = "sa";
    private static final String PASSWORD = "";

    public static void main(String[] args) {
        try (Connection conn = DriverManager.getConnection(JDBC_URL, USER, PASSWORD)) {
            if (!tableExists(conn, "USUARIOSSERVIDOR")) {
                System.out.println("Base de datos no encontrada, inicializando...");
                executeSchema(conn);
                System.out.println("Base de datos inicializada correctamente.");
            } else {
                System.out.println("Base de datos ya inicializada.");
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private static boolean tableExists(Connection conn, String tableName) throws SQLException {
        DatabaseMetaData meta = conn.getMetaData();
        try (ResultSet rs = meta.getTables(null, null, tableName.toUpperCase(), null)) {
            return rs.next();
        }
    }

    private static void executeSchema(Connection conn) throws Exception {
        String sql;
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(
                DatabaseInitializer.class.getClassLoader().getResourceAsStream("schema.sql"),
                StandardCharsets.UTF_8))) {
            sql = reader.lines().collect(Collectors.joining("\n"));
        }

        try (Statement stmt = conn.createStatement()) {
            for (String statement : sql.split(";")) {
                if (!statement.trim().isEmpty()) {
                    stmt.execute(statement.trim());
                }
            }
        }
    }
}
