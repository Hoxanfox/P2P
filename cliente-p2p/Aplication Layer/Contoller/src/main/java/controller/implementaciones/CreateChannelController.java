package controller.implementaciones;

import dto.implementacion.CreateChannel.CreateChannelRequestDto;
import dto.implementacion.CreateChannel.CreateChannelResponseDto;
import facade.implementaciones.CreateChannelFacade;
import transport.TransportContext;

import java.sql.SQLException;

public class CreateChannelController {

    private final CreateChannelFacade channelFacade;
    private final CreateChannelRequestDto requestDto;

    public CreateChannelController(TransportContext context, CreateChannelRequestDto requestDto) {
        this.channelFacade = new CreateChannelFacade(context);
        this.requestDto = requestDto;
        System.out.println("[DEBUG] CreateChannelController inicializado.");
    }

    /**
     * Ejecuta el flujo completo: obtiene datos del servidor y los guarda en la base de datos.
     *
     * @return DTO con la respuesta procesada o null si ocurre un error.
     */
    public CreateChannelResponseDto crearCanal() {
        System.out.println("[DEBUG] Iniciando creación de canal desde el controlador...");

        try {
            CreateChannelResponseDto responseDto = channelFacade.crearFlujo(requestDto);
            System.out.println("[DEBUG] Canal creado y persistido correctamente.");
            return responseDto;

        } catch (SQLException e) {
            System.err.println("[ERROR] Error durante la creación del canal: " + e.getMessage());
            e.printStackTrace();
            return null;
        }
    }
}
