package dao.implementaciones.invitacionCanal;

import dao.DatabaseConfig;
import dto.implementacion.InvitacionCanal.InvitacionCanalResponseDto;
import dto.implementacion.InvitacionCanal.Response.CanalResponseDto;
import dto.implementacion.InvitacionCanal.Response.DestinatarioResponseDto;
import dto.implementacion.InvitacionCanal.Response.InvitacionResponseDto;
import dto.implementacion.InvitacionCanal.Response.NotificacionResponseDto;

import java.sql.*;
import java.util.UUID;
import java.time.LocalDateTime;

public class InvitacionCanalDao {

    private static final String INSERT_INVI = "INSERT INTO Invitacion (id_invitacion, id_usuario, fecha_envio, estado) VALUES (?, ?, ?, ?)";
    private static final String INSERT_IC = "INSERT INTO InvitacionCanal (id_invitacion_canal, id_canal_servidor, id_invitacion) VALUES (?, ?, ?)";
    private static final String INSERT_NOTI = "INSERT INTO Notificacion (id_notificacion, id_invitacion, contenido) VALUES (?, ?, ?)";

    /**
     * Guarda una invitación y su notificación asociada en la base de datos.
     * @param idUsuario El ID del usuario que recibe la invitación.
     * @param idCanalServidor El ID del canal de servidor relacionado con la invitación.
     * @param contenidoNotificacion El contenido de la notificación.
     * @return Un objeto de respuesta que incluye la invitación y la notificación.
     */
    public InvitacionCanalResponseDto guardarInvitacionYNotificacion(String idUsuario, String idCanalServidor, String contenidoNotificacion) {
        Connection conn = null;
        InvitacionCanalResponseDto respuesta = new InvitacionCanalResponseDto();

        try {
            conn = DatabaseConfig.getConnection();
            conn.setAutoCommit(false); // Comienza la transacción
            System.out.println("Conexión establecida, comenzando la transacción...");

            // Generar IDs únicos
            String idInvitacion = UUID.randomUUID().toString();
            String idNotificacion = UUID.randomUUID().toString();
            String idInvitacionCanal = UUID.randomUUID().toString();
            String fechaEnvio = LocalDateTime.now().toString();

            // 1. Insertar en la tabla Invitacion
            try (PreparedStatement stmt = conn.prepareStatement(INSERT_INVI)) {
                System.out.println("Preparando INSERT en Invitacion...");
                stmt.setString(1, idInvitacion);
                stmt.setString(2, idUsuario);
                stmt.setString(3, fechaEnvio);
                stmt.setInt(4, 0); // estado pendiente
                stmt.executeUpdate();
                System.out.println("INSERT en Invitacion ejecutado.");
            }

            // 2. Insertar en la tabla InvitacionCanal
            try (PreparedStatement stmt = conn.prepareStatement(INSERT_IC)) {
                System.out.println("Preparando INSERT en InvitacionCanal...");
                stmt.setString(1, idInvitacionCanal);
                stmt.setString(2, idCanalServidor);
                stmt.setString(3, idInvitacion);
                stmt.executeUpdate();
                System.out.println("INSERT en InvitacionCanal ejecutado.");
            }

            // 3. Insertar en la tabla Notificacion
            try (PreparedStatement stmt = conn.prepareStatement(INSERT_NOTI)) {
                System.out.println("Preparando INSERT en Notificacion...");
                stmt.setString(1, idNotificacion);
                stmt.setString(2, idInvitacion);
                stmt.setString(3, contenidoNotificacion);
                stmt.executeUpdate();
                System.out.println("INSERT en Notificacion ejecutado.");
            }

            conn.commit(); // Confirmar transacción
            System.out.println("Transacción confirmada.");

            // 4. Construir la respuesta DTO
            CanalResponseDto canalDto = new CanalResponseDto();
            canalDto.setId(Long.parseLong(idCanalServidor)); // Suponiendo que el ID del canal es de tipo Long

            DestinatarioResponseDto destinatarioDto = new DestinatarioResponseDto();
            destinatarioDto.setId(Long.parseLong(idUsuario)); // Suponiendo que el ID del usuario es de tipo Long
            destinatarioDto.setNombre(obtenerNombreUsuario(conn, idUsuario)); // Obtener el nombre del usuario desde la base de datos

            InvitacionResponseDto invitacionDto = new InvitacionResponseDto();
            invitacionDto.setId(Long.parseLong(idInvitacion)); // Suponiendo que el ID de la invitación es de tipo Long
            invitacionDto.setFechaEnvio(LocalDateTime.parse(fechaEnvio));
            invitacionDto.setEstado("Pendiente");
            invitacionDto.setCanal(canalDto);
            invitacionDto.setDestinatario(destinatarioDto);

            NotificacionResponseDto notiDto = new NotificacionResponseDto();
            notiDto.setId(Long.parseLong(idNotificacion)); // Suponiendo que el ID de la notificación es de tipo Long
            notiDto.setContenido(contenidoNotificacion);
            notiDto.setInvitacion(invitacionDto);

            // Asignar la invitación y notificación a la respuesta DTO
            respuesta.setInvitacion(invitacionDto);
            respuesta.setNotificacion(notiDto);
            System.out.println("Respuesta DTO construida exitosamente.");

        } catch (Exception e) {
            if (conn != null) {
                try {
                    conn.rollback(); // Rollback si hay un error
                    System.out.println("Error ocurrido, realizando rollback...");
                } catch (SQLException ex) {
                    ex.printStackTrace();
                }
            }
            e.printStackTrace();
        } finally {
            try {
                if (conn != null) conn.close(); // Cerrar la conexión
                System.out.println("Conexión cerrada.");
            } catch (SQLException e) {
                e.printStackTrace();
            }
        }

        return respuesta;
    }

    /**
     * Obtiene el nombre del usuario desde la base de datos.
     * @param conn La conexión a la base de datos.
     * @param idUsuario El ID del usuario.
     * @return El nombre del usuario.
     * @throws SQLException Si ocurre un error en la consulta.
     */
    private String obtenerNombreUsuario(Connection conn, String idUsuario) throws SQLException {
        String sql = "SELECT nombre FROM UsuariosServidor WHERE id_usuario_servidor = ?";
        try (PreparedStatement stmt = conn.prepareStatement(sql)) {
            stmt.setString(1, idUsuario);
            ResultSet rs = stmt.executeQuery();
            if (rs.next()) {
                String nombre = rs.getString("nombre");
                System.out.println("Nombre del usuario encontrado: " + nombre);
                return nombre;
            }
        }
        System.out.println("Usuario no encontrado, retornando 'Desconocido'.");
        return "Desconocido"; // Si no se encuentra el usuario, retornamos "Desconocido"
    }
}
