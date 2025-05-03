package facade.interfaces;

import transport.TransportContext;
import dto.implementacion.CreateChannel.CreateChannelRequestDto;
import dto.implementacion.CreateChannel.CreateChannelResponseDto;
import java.sql.SQLException;

public interface IChannelFacade {

    /**
     * Función para obtener información del servidor y procesar la solicitud de creación de canal.
     * @param request La solicitud de creación de canal.
     * @param context El contexto de transporte para conectar con el servidor.
     * @return Un DTO con la información procesada.
     * @throws SQLException Si ocurre un error al interactuar con la base de datos.
     */
    CreateChannelResponseDto obtenerInformacionDelServidor(CreateChannelRequestDto request, TransportContext context) throws SQLException;

    /**
     * Función para crear un flujo en el que se persiste la información en la base de datos.
     *
     * @param responseDto La respuesta con la información procesada para el canal.
     * @return
     * @throws SQLException Si ocurre un error al persistir en la base de datos.
     */
    CreateChannelResponseDto crearFlujoYPersistir(CreateChannelResponseDto responseDto) throws SQLException;
}
