package dao.implementaciones.sendMessageUser;

import dao.DatabaseConfig;
import dto.implementacion.SendMessageUser.Response.SendMessageUserResponseDto;
import dto.implementacion.SendMessageUser.Response.MensajeResponseDto;
import dto.implementacion.SendMessageUser.Response.ChatResponseDto;
import dto.implementacion.SendMessageUser.RemitenteMessageDto;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

public class SendMessageUserDao {

    public void saveMessage(SendMessageUserResponseDto responseDto) throws SQLException {
        Connection conn = DatabaseConfig.getConnection();
        conn.setAutoCommit(false); // Iniciar transacción

        try {
            MensajeResponseDto mensaje = responseDto.getMensaje();
            RemitenteMessageDto remitente = mensaje.getRemitente();
            ChatResponseDto chat = mensaje.getChat();

            // Asegurarse de que el remitente esté presente en la base de datos
            String queryUsuario = "SELECT id_usuario_servidor FROM UsuariosServidor WHERE id_usuario_servidor = ?";
            try (PreparedStatement stmt = conn.prepareStatement(queryUsuario)) {
                stmt.setString(1, remitente.getId().toString());
                ResultSet rs = stmt.executeQuery();

                // Si el usuario no existe, podemos insertarlo (suponiendo que el DTO de usuario tiene el nombre y el email)
                if (!rs.next()) {
                    String insertUsuario = "INSERT INTO UsuariosServidor (id_usuario_servidor, nombre, email, password, estado) VALUES (?, ?, ?, ?, ?)";
                    try (PreparedStatement insertStmt = conn.prepareStatement(insertUsuario)) {
                        insertStmt.setString(1, remitente.getId().toString());
 // Asumiendo que 'nombre' está en el DTO de remitente
                        insertStmt.setString(3, remitente.getCorreo()); // Asumiendo que 'email' está en el DTO de remitente
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
                        insertStmt.setString(2, chat.getTipoChatId().toString()); // Asumiendo que el tipo de chat está en el DTO
                        insertStmt.executeUpdate();
                    }
                }
            }

            // Insertar el archivo si existe (opcional)
            String archivoId = null;
            if (mensaje.getArchivo() != null && !mensaje.getArchivo().isEmpty()) {
                archivoId = UUID.randomUUID().toString();
                String insertArchivo = "INSERT INTO Archivos (id_archivo, nombre, binario) VALUES (?, ?, ?)";
                try (PreparedStatement stmt = conn.prepareStatement(insertArchivo)) {
                    stmt.setString(1, archivoId);
                    stmt.setString(2, "Archivo adjunto"); // Aquí puedes ajustar según la información del archivo
                    stmt.setBytes(3, mensaje.getArchivo().getBytes()); // Asumimos que 'archivo' es una cadena Base64
                    stmt.executeUpdate();
                }
            }

            // Insertar el mensaje en la tabla MensajeServidor
            String mensajeId = mensaje.getId().toString();
            String insertMensaje = "INSERT INTO MensajeServidor (id_mensaje_servidor, id_chat, id_usuario, contenido, id_archivo, fecha_envio) VALUES (?, ?, ?, ?, ?, ?)";
            try (PreparedStatement stmt = conn.prepareStatement(insertMensaje)) {
                stmt.setString(1, mensajeId);
                stmt.setString(2, chat.getId().toString());
                stmt.setString(3, remitente.getId().toString());
                stmt.setString(4, mensaje.getContenido());
                stmt.setString(5, archivoId); // Si no hay archivo, se puede pasar NULL
                stmt.setString(6, mensaje.getFechaEnvio());
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
}
