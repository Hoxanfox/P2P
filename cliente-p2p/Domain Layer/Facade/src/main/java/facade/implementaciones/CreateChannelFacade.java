package facade.implementaciones;

import dao.implementaciones.creactChannel.CreactChannelDao;
import dto.implementacion.CreateChannel.CreateChannelRequestDto;
import dto.implementacion.CreateChannel.CreateChannelResponseDto;
import facade.interfaces.IChannelFacade;
import protocolo.implementaciones.createChannel.CreateChannelRequest;
import protocolo.implementaciones.createChannel.CreateChannelResponse;
import transport.TransportContext;

import java.sql.SQLException;

public class CreateChannelFacade implements IChannelFacade {

    private final CreactChannelDao creactChannelDao;
    private final TransportContext transportContext;

    public CreateChannelFacade(TransportContext transportContext) {
        this.creactChannelDao = new CreactChannelDao();
        this.transportContext = transportContext;
        System.out.println("[DEBUG] CreateChannelFacade inicializado.");
    }

    @Override
    public CreateChannelResponseDto obtenerInformacionDelServidor(CreateChannelRequestDto requestDto, TransportContext context) throws SQLException {
        System.out.println("[DEBUG] [CreateChannelFacade] Iniciando obtención de información del servidor...");

        CreateChannelRequest protocoloRequest = new CreateChannelRequest(requestDto);
        System.out.println("[DEBUG] [CreateChannelFacade] Objeto protocoloRequest creado: " + protocoloRequest);

        String jsonRequest = protocoloRequest.toJson();
        System.out.println("[DEBUG] [CreateChannelFacade] JSON enviado al servidor: " + jsonRequest);

        String jsonResponse = context.executeSend(jsonRequest);
        System.out.println("[DEBUG] [CreateChannelFacade] JSON recibido del servidor: " + jsonResponse);

        CreateChannelResponse protocoloResponse = new CreateChannelResponse();
        protocoloResponse.fromJson(jsonResponse);
        System.out.println("[DEBUG] [CreateChannelFacade] Objeto protocoloResponse procesado: " + protocoloResponse);

        CreateChannelResponseDto responseDto = protocoloResponse.getData();
        System.out.println("[DEBUG] [CreateChannelFacade] DTO de respuesta retornado: " + responseDto);
        return responseDto;
    }

    @Override
    public CreateChannelResponseDto crearFlujoYPersistir(CreateChannelResponseDto responseDto) throws SQLException {
        System.out.println("[DEBUG] [CreateChannelFacade] Iniciando persistencia en base de datos...");
        System.out.println("[DEBUG] [CreateChannelFacade] Datos a persistir: " + responseDto);

        creactChannelDao.save(responseDto);

        System.out.println("[DEBUG] [CreateChannelFacade] Persistencia completada.");
        return responseDto;
    }

    // NUEVO MÉTODO: Ejecuta todo el flujo (obtener y persistir)
    public CreateChannelResponseDto crearFlujo(CreateChannelRequestDto requestDto) throws SQLException {
        System.out.println("[DEBUG] [CreateChannelFacade] Iniciando flujo completo de creación de canal...");

        CreateChannelResponseDto responseDto = obtenerInformacionDelServidor(requestDto, transportContext);
        System.out.println("[DEBUG] [CreateChannelFacade] Información del servidor obtenida correctamente.");

        CreateChannelResponseDto result = crearFlujoYPersistir(responseDto);
        System.out.println("[DEBUG] [CreateChannelFacade] Flujo completo finalizado con éxito.");

        return result;
    }
    public TransportContext getContext() {
        return this.transportContext;
    }

}
