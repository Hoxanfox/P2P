package facade.implementaciones;

import dao.implementaciones.register.RegisterDao;
import dto.implementacion.register.RegisterRequestDto;
import dto.implementacion.register.RegisterResponseDto;
import facade.interfaces.IRegisterFacade;
import protocolo.implementaciones.register.RegisterRequest;
import protocolo.implementaciones.register.RegisterResponse;
import transport.TransportContext;
import transport.interfaces.ITransportStrategy;

import java.sql.SQLException;

public class RegisterFacade implements IRegisterFacade {

    private final ITransportStrategy strategy;


    public RegisterFacade(ITransportStrategy strategy) {
        this.strategy = strategy;
    }

    @Override
    public RegisterResponseDto obtenerInformacionDelServidor(RegisterRequestDto requestDto) {
        // 1. Preparar request
        RegisterRequest request = new RegisterRequest(requestDto);
        String jsonToSend = request.toJson();

        // 2. Ejecutar transporte
        TransportContext context = new TransportContext(strategy);
        String jsonResponse = context.executeSend(jsonToSend);

        if (jsonResponse == null || jsonResponse.isBlank()) {
            System.out.println("[ERROR] No se recibió respuesta del servidor.");
            return null;
        }

        // 3. Procesar respuesta
        RegisterResponse response = new RegisterResponse();
        response.fromJson(jsonResponse);

        if ("success".equalsIgnoreCase(response.getStatus())) {
            return response.getData();
        } else {
            System.out.println("[ERROR] Respuesta fallida: " + response.getMessage());
            return null;
        }
    }

    @Override
    public boolean persistirUsuario(RegisterResponseDto responseDto) {
        if (responseDto == null) {
            System.out.println("[ERROR] El DTO de respuesta es nulo. No se puede persistir.");
            return false;
        }

        try {
            RegisterDao dao = new RegisterDao();
            boolean inserted = dao.registrarUsuario(responseDto);

            if (!inserted) {
                System.out.println("[ERROR] No se insertó ninguna fila en la base de datos.");
            }

            return inserted;
        } catch (SQLException e) {
            System.out.println("[ERROR] No se pudo insertar el usuario: " + e.getMessage());
            return false;
        }
    }

    @Override
    public void ejecutarFlujoRegistro(RegisterRequestDto requestDto) {
        if (requestDto == null) {
            System.out.println("[ERROR] El DTO de solicitud es nulo.");
            return;
        }

        RegisterResponseDto responseDto = obtenerInformacionDelServidor(requestDto);

        if (responseDto != null) {
            boolean success = persistirUsuario(responseDto);
            if (success) {
                System.out.println("[INFO] Usuario registrado y persistido exitosamente.");
            }
        } else {
            System.out.println("[ERROR] No se pudo obtener información válida del servidor.");
        }
    }
}
