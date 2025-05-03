package facade.implementaciones;

import dao.implementaciones.sendMessageUser.SendMessageUserDao;
import dto.implementacion.SendMessageUser.Request.SendMessageUserRequestDto;
import dto.implementacion.SendMessageUser.Response.SendMessageUserResponseDto;
import protocolo.implementaciones.sendMessageUser.SendMessageUserRequest;
import protocolo.implementaciones.sendMessageUser.SendMessageUserResponse;
import transport.TransportContext;

import java.sql.SQLException;

public class SendMessageUserFacade {

    private static final String CLASS_NAME = "SendMessageUserFacade";

    private TransportContext transportContext;
    private SendMessageUserDao sendMessageUserDao;

    public SendMessageUserFacade(TransportContext transportContext) {
        System.out.println(CLASS_NAME + " -> Constructor: Inicializando DAO y contexto de transporte");
        this.transportContext = transportContext;
        this.sendMessageUserDao = new SendMessageUserDao();
    }

    // Coordinador del flujo completo
    public void processMessage(SendMessageUserRequestDto requestDto) throws SQLException {
        System.out.println(CLASS_NAME + " -> processMessage: Iniciando proceso");

        // Paso 1: Obtener información del servidor
        SendMessageUserResponseDto responseDto = obtainInformationFromServer(requestDto);

        // Paso 2: Persistir la información
        persistMessage(responseDto);

        System.out.println(CLASS_NAME + " -> processMessage: Proceso completado");
    }

    // Obtiene la información del servidor
    private SendMessageUserResponseDto obtainInformationFromServer(SendMessageUserRequestDto requestDto) {
        System.out.println(CLASS_NAME + " -> obtainInformationFromServer: Preparando solicitud");

        // Convertir DTO a JSON
        SendMessageUserRequest sendMessageUserRequest = new SendMessageUserRequest(requestDto);
        String jsonRequest = sendMessageUserRequest.toJson();
        System.out.println(CLASS_NAME + " -> obtainInformationFromServer: JSON de solicitud: " + jsonRequest);

        // Ejecutar transporte
        String jsonResponse = transportContext.executeSend(jsonRequest);
        System.out.println(CLASS_NAME + " -> obtainInformationFromServer: JSON de respuesta: " + jsonResponse);

        // Procesar respuesta
        SendMessageUserResponse sendMessageUserResponse = new SendMessageUserResponse();
        sendMessageUserResponse.fromJson(jsonResponse);

        SendMessageUserResponseDto responseDto = sendMessageUserResponse.getData();
        System.out.println(CLASS_NAME + " -> obtainInformationFromServer: Datos extraídos del servidor");

        return responseDto;
    }

    // Persiste la información usando el DAO
    private void persistMessage(SendMessageUserResponseDto responseDto) throws SQLException {
        System.out.println(CLASS_NAME + " -> persistMessage: Guardando mensaje en base de datos");

        sendMessageUserDao.saveMessage(responseDto);

        System.out.println(CLASS_NAME + " -> persistMessage: Mensaje guardado correctamente");
    }
}
