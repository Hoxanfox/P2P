package dao.implementaciones.sendMessageUser;

import dao.DatabaseConfig;
import dto.implementacion.SendMessageUser.Response.SendMessageUserResponseDto;
import dto.implementacion.SendMessageUser.Response.ChatResponseDto;
import dto.implementacion.SendMessageUser.RemitenteMessageDto;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

public class SendMessageUserDao {

    public void saveMessage(SendMessageUserResponseDto responseDto) throws SQLException {
        Connection conn = DatabaseConfig.getConnection();
        conn.setAutoCommit(false); // Iniciar transacción

        try {
            // Obtener los datos del remitente y chat
            RemitenteMessageDto remitente = responseDto.getRemitente();
            ChatResponseDto chat = responseDto.getChat();
            String contenido = responseDto.getContenido();
            String archivo = responseDto.getArchivo();

            // Mostrar todos los usuarios activos antes de intentar insertar el nuevo usuario
            printActiveUsuarios(conn);

            // Asegurarse de que el remitente esté presente en la base de datos
            String queryUsuario = "SELECT id_usuario_servidor FROM UsuariosServidor WHERE id_usuario_servidor = ?";
            try (PreparedStatement stmt = conn.prepareStatement(queryUsuario)) {
                stmt.setString(1, remitente.getId().toString());
                ResultSet rs = stmt.executeQuery();

                // Si el usuario no existe, insertarlo
                if (!rs.next()) {
                    String insertUsuario = "INSERT INTO UsuariosServidor (id_usuario_servidor, nombre, email, password, foto, ip, estado) VALUES (?, ?, ?, ?, ?, ?, ?)";
                    try (PreparedStatement insertStmt = conn.prepareStatement(insertUsuario)) {
                        insertStmt.setString(1, remitente.getId().toString());
           
                        insertStmt.setString(3, remitente.getCorreo()); // Correo del remitente
                      // IP del remitente (si la hay)
                        insertStmt.setInt(7, 1); // Estado 1 para activo (puedes ajustarlo según tu lógica)
                        insertStmt.executeUpdate();
                    }
                }
            }

            // Asegurarse de que el chat esté presente en la base de datos
            String queryChat = "SELECT id_chat FROM Chat WHERE id_chat = ?";
            try (PreparedStatement stmt = conn.prepareStatement(queryChat)) {
                stmt.setString(1, chat.getId().toString());
                ResultSet rs = stmt.executeQuery();

                // Si el chat no existe, insertarlo
                if (!rs.next()) {
                    String insertChat = "INSERT INTO Chat (id_chat, id_chat_tipo) VALUES (?, ?)";
                    try (PreparedStatement insertStmt = conn.prepareStatement(insertChat)) {
                        insertStmt.setString(1, chat.getId().toString());
                        insertStmt.setString(2, chat.getTipoChatId().toString()); // Tipo de chat
                        insertStmt.executeUpdate();
                    }
                }
            }

            // Insertar el archivo si existe (opcional)
            String archivoId = null;
            if (archivo != null && !archivo.isEmpty()) {
                archivoId = UUID.randomUUID().toString();
                String insertArchivo = "INSERT INTO Archivos (id_archivo, nombre, binario) VALUES (?, ?, ?)";
                try (PreparedStatement stmt = conn.prepareStatement(insertArchivo)) {
                    stmt.setString(1, archivoId);
                    stmt.setString(2, "Archivo adjunto"); // Descripción del archivo
                    stmt.setBytes(3, archivo.getBytes()); // Asumiendo que el archivo es una cadena Base64
                    stmt.executeUpdate();
                }
            }

            // Insertar el mensaje en la tabla MensajeServidor
            String mensajeId = UUID.randomUUID().toString(); // Generar un ID para el mensaje
            String insertMensaje = "INSERT INTO MensajeServidor (id_mensaje_servidor, id_chat, id_usuario, contenido, id_archivo, fecha_envio) VALUES (?, ?, ?, ?, ?, ?)";
            try (PreparedStatement stmt = conn.prepareStatement(insertMensaje)) {
                stmt.setString(1, mensajeId);
                stmt.setString(2, chat.getId().toString());
                stmt.setString(3, remitente.getId().toString());
                stmt.setString(4, contenido);
                stmt.setString(5, archivoId); // Si no hay archivo, se puede pasar NULL
                stmt.setString(6, responseDto.getFechaEnvio());
                stmt.executeUpdate();
                System.out.println("[SendMessageUserDao - saveMessage] Mensaje insertado correctamente.");
            }

            conn.commit(); // Confirmar la transacción
        } catch (SQLException e) {
            conn.rollback(); // Rollback en caso de error
            System.err.println("[SendMessageUserDao - saveMessage] Error en transacción, rollback: " + e.getMessage());
            throw e;
        } finally {
            conn.setAutoCommit(true); // Restaurar configuración por defecto
        }
    }

    // Método para imprimir todos los usuarios activos de la tabla UsuariosServidor
    private void printActiveUsuarios(Connection conn) throws SQLException {
        String query = "SELECT id_usuario_servidor, nombre, email, password, foto, ip, estado FROM UsuariosServidor WHERE estado = 1"; // Solo usuarios activos
        try (PreparedStatement stmt = conn.prepareStatement(query)) {
            ResultSet rs = stmt.executeQuery();
            System.out.println("Usuarios activos en la tabla UsuariosServidor:");
            while (rs.next()) {
                String idUsuario = rs.getString("id_usuario_servidor");
                String nombre = rs.getString("nombre");
                String email = rs.getString("email");
                String password = rs.getString("password"); // Contraseña
                byte[] foto = rs.getBytes("foto");
                String ip = rs.getString("ip");
                int estado = rs.getInt("estado");
                System.out.println("ID: " + idUsuario + ", Nombre: " + nombre + ", Email: " + email + ", Contraseña: " + password + ", Foto: " + (foto != null ? "Disponible" : "No disponible") + ", IP: " + ip + ", Estado: " + estado);
            }
        }
    }
}
