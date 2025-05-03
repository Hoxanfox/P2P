package dao.implementaciones.creactChannel;

import dao.DatabaseConfig;
import dto.implementacion.CreateChannel.CreateChannelResponseDto;
import dto.implementacion.CreateChannel.Channel.InvitacionDto;
import dto.implementacion.CreateChannel.Channel.MiembroCanalDto;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.UUID;

public class CreactChannelDao {

    // Método para insertar un nuevo tipo de chat
    private void insertChatTipo(Connection conn, String tipoChatId, String tipoChat) throws SQLException {
        // Validación de tipoChat
        if (tipoChat == null || tipoChat.isEmpty()) {
            throw new IllegalArgumentException("El tipo de chat no puede ser nulo o vacío");
        }

        // Verificamos si el tipo de chat ya existe
        if (getChatTipoId(conn, tipoChat) == null) {
            // Si no existe, lo insertamos
            String query = "INSERT INTO ChatTipo (id_chat_tipo, nombre) VALUES (?, ?)";
            try (PreparedStatement stmt = conn.prepareStatement(query)) {
                stmt.setString(1, tipoChatId);
                stmt.setString(2, tipoChat);
                stmt.executeUpdate();
                System.out.println("[CreactChannelDao - insertChatTipo] Tipo de chat insertado: " + tipoChat);
            }
        } else {
            System.out.println("[CreactChannelDao - insertChatTipo] El tipo de chat ya existe: " + tipoChat);
        }
    }

    private String getChatTipoId(Connection conn, String tipoChat) throws SQLException {
        String tipoChatId = null;
        String query = "SELECT id_chat_tipo FROM ChatTipo WHERE nombre = ?";
        try (PreparedStatement stmt = conn.prepareStatement(query)) {
            stmt.setString(1, tipoChat);
            try (ResultSet rs = stmt.executeQuery()) {
                if (rs.next()) {
                    tipoChatId = rs.getString("id_chat_tipo");
                    System.out.println("[CreactChannelDao - getChatTipoId] Tipo de chat encontrado con ID: " + tipoChatId);
                }
            }
        }
        return tipoChatId;
    }

    public void save(CreateChannelResponseDto responseDto) throws SQLException {
        Connection conn = DatabaseConfig.getConnection();
        conn.setAutoCommit(false); // Iniciamos una transacción

        try {
            // 1. Insertar en CanalesServidor
            String canalId = responseDto.getId().toString();
            System.out.println("[CreactChannelDao - save] Insertando canal: " + canalId);
            try (PreparedStatement stmt = conn.prepareStatement(
                    "INSERT INTO CanalesServidor (id_canal_servidor, nombre, descripcion) VALUES (?, ?, ?)")) {
                stmt.setString(1, canalId);
                stmt.setString(2, responseDto.getNombre());
                stmt.setString(3, responseDto.getDescripcion());
                stmt.executeUpdate();
                System.out.println("[CreactChannelDao - save] Canal insertado correctamente.");
            }

            // 2. Verificar e insertar tipos de chat
            String tipoChat = responseDto.getChat().getTipo(); // Tipo de chat: "privado" o "público"
            String chatId = responseDto.getChat().getId().toString();
            System.out.println("[CreactChannelDao - save] Verificando tipo de chat: " + tipoChat);

            // Validar que el tipo de chat no sea null ni vacío
            if (tipoChat == null || tipoChat.isEmpty()) {
                throw new IllegalArgumentException("El tipo de chat no puede ser nulo o vacío");
            }

            // Verificar si el tipo de chat existe
            String tipoChatId = getChatTipoId(conn, tipoChat);
            if (tipoChatId == null) {
                // Si no existe, insertar el tipo de chat
                tipoChatId = UUID.randomUUID().toString();
                insertChatTipo(conn, tipoChatId, tipoChat); // Insertar el tipo de chat
            }

            // 3. Insertar el chat
            System.out.println("[CreactChannelDao - save] Insertando chat con ID: " + chatId);
            try (PreparedStatement insertChat = conn.prepareStatement(
                    "INSERT INTO Chat (id_chat, id_chat_tipo) VALUES (?, ?)")) {
                insertChat.setString(1, chatId);
                insertChat.setString(2, tipoChatId);
                insertChat.executeUpdate();
                System.out.println("[CreactChannelDao - save] Chat insertado correctamente.");
            }

            // 4. Insertar en ChatMiembrosPublico
            String chatMiembroId = UUID.randomUUID().toString();
            System.out.println("[CreactChannelDao - save] Insertando en ChatMiembrosPublico con ID: " + chatMiembroId);
            try (PreparedStatement stmt = conn.prepareStatement(
                    "INSERT INTO ChatMiembrosPublico (id_chat_miembros, id_canal, id_chat) VALUES (?, ?, ?)")) {
                stmt.setString(1, chatMiembroId);
                stmt.setString(2, canalId);
                stmt.setString(3, chatId);
                stmt.executeUpdate();
                System.out.println("[CreactChannelDao - save] Miembros públicos insertados.");
            }

            // 5. Insertar miembros en CanalMiembros
            if (responseDto.getMiembros() != null) {
                for (MiembroCanalDto miembro : responseDto.getMiembros()) {
                    String miembroId = UUID.randomUUID().toString();
                    System.out.println("[CreactChannelDao - save] Insertando miembro al canal: " + miembro.getId());
                    try (PreparedStatement stmt = conn.prepareStatement(
                            "INSERT INTO CanalMiembros (id_canal_miembro, id_usuario_servidor, id_canal_servidor) VALUES (?, ?, ?)")) {
                        stmt.setString(1, miembroId);
                        stmt.setString(2, String.valueOf(miembro.getId()));
                        stmt.setString(3, canalId);
                        stmt.executeUpdate();
                        System.out.println("[CreactChannelDao - save] Miembro insertado: " + miembro.getId());
                    }
                }
            }

            // 6. Insertar invitaciones y relación con canal
            if (responseDto.getInvitaciones() != null) {
                for (InvitacionDto invitacion : responseDto.getInvitaciones()) {
                    String invitacionId = invitacion.getId() != null
                            ? String.valueOf(invitacion.getId())
                            : UUID.randomUUID().toString();
                    System.out.println("[CreactChannelDao - save] Insertando invitación para usuario: " + invitacion.getDestinatario().getId());

                    try (PreparedStatement stmt = conn.prepareStatement(
                            "INSERT INTO Invitacion (id_invitacion, id_usuario, fecha_envio, estado) VALUES (?, ?, ?, ?)")) {
                        stmt.setString(1, invitacionId);
                        stmt.setString(2, String.valueOf(invitacion.getDestinatario().getId()));
                        stmt.setString(3, invitacion.getFechaEnvio());
                        stmt.setObject(4, invitacion.getEstado());
                        stmt.executeUpdate();
                        System.out.println("[CreactChannelDao - save] Invitación insertada: " + invitacionId);
                    }

                    String invitacionCanalId = UUID.randomUUID().toString();
                    try (PreparedStatement stmt = conn.prepareStatement(
                            "INSERT INTO InvitacionCanal (id_invitacion_canal, id_canal_servidor, id_invitacion) VALUES (?, ?, ?)")) {
                        stmt.setString(1, invitacionCanalId);
                        stmt.setString(2, canalId);
                        stmt.setString(3, invitacionId);
                        stmt.executeUpdate();
                        System.out.println("[CreactChannelDao - save] Relación invitación-canal insertada.");
                    }
                }
            }

            conn.commit(); // Confirmar transacción
            System.out.println("[CreactChannelDao - save] Transacción completada correctamente.");
        } catch (SQLException e) {
            conn.rollback(); // En caso de error, revertir
            System.err.println("[CreactChannelDao - save] Error durante la transacción, haciendo rollback: " + e.getMessage());
            throw e;
        } finally {
            conn.setAutoCommit(true); // Restaurar comportamiento por defecto
            System.out.println("[CreactChannelDao - save] AutoCommit restaurado.");
        }
    }
}
