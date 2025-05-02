package dao.test;

import dao.DatabaseConfig;

import java.sql.Connection;
import java.sql.Statement;

public class DatabaseTestUtil {

    public static void inicializarBaseDeDatosTest() throws Exception {
        try (Connection conn = DatabaseConfig.getConnection(); Statement stmt = conn.createStatement()) {
            stmt.execute("""
                CREATE TABLE IF NOT EXISTS usuarios (
                    id UUID PRIMARY KEY,
                    nombre VARCHAR(100),
                    email VARCHAR(100),
                    password VARCHAR(100),
                    foto VARCHAR(255),
                    estado VARCHAR(50),
                    fecha_registro TIMESTAMP
                )
            """);
        }
    }
}
