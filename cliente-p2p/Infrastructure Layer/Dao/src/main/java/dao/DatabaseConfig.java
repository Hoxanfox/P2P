package dao;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;

public class DatabaseConfig {
    // Ruta del archivo SQLite. Se creará ./data/mydatabase.db si no existe.
    private static final String JDBC_URL = "jdbc:sqlite:./data/mydatabase.db";

    private static Connection connection;
    private static Connection testConnection = null;

    /**
     * Obtiene una conexión singleton a la base de datos SQLite.
     * Si no existe o está cerrada, la abre.
     */
    public static Connection getConnection() throws SQLException {
        if (connection == null || connection.isClosed()) {
            connection = DriverManager.getConnection(JDBC_URL);
        }
        return connection;
    }

    public static void setConnectionForTesting(Connection conn) {
        testConnection = conn;
    }
    public static Connection getConnectionTest() throws SQLException {
        if (testConnection != null) {
            return testConnection;
        }
        return DriverManager.getConnection("jdbc:sqlite:./data/mydatabase.db");
    }
}
