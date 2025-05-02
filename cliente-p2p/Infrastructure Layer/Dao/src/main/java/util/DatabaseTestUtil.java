package util;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.nio.charset.StandardCharsets;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.Statement;
import java.util.stream.Collectors;

public class DatabaseTestUtil {

    private static final String JDBC_URL = "jdbc:h2:mem:testdb;DB_CLOSE_DELAY=-1";
    private static final String USER = "sa";
    private static final String PASSWORD = "";

    public static void inicializarBaseDeDatosTest() throws Exception {
        try (Connection conn = DriverManager.getConnection(JDBC_URL, USER, PASSWORD);
             Statement stmt = conn.createStatement()) {

            String sql = new BufferedReader(
                    new InputStreamReader(DatabaseTestUtil.class.getClassLoader().getResourceAsStream("schema.sql"),
                            StandardCharsets.UTF_8)
            ).lines().collect(Collectors.joining("\n"));

            for (String s : sql.split(";")) {
                if (!s.trim().isEmpty()) {
                    stmt.execute(s.trim());
                }
            }
        }
    }
}
