package facade.implementaciones;

import dto.implementacion.login.LoginRequestDto;
import dto.implementacion.login.LoginResponseDto;
import facade.interfaces.ILoginFacade;
import protocolo.implementaciones.login.LoginRequest;
import protocolo.implementaciones.login.LoginResponse;
import transport.TransportContext;
import transport.interfaces.ITransportStrategy;

public class LoginFacade implements ILoginFacade {

    private final TransportContext context;
    private LoginRequestDto requestDto = new LoginRequestDto();
    private LoginResponseDto responseDto = new LoginResponseDto();

    public LoginFacade(TransportContext context) {
        this.context = context;
        log("[INFO] LoginFacade creado con contexto de transporte usando estrategia: "
                + context.getStrategy().getClass().getSimpleName());
    }

    @Override
    public boolean validarRequest(LoginRequestDto requestDto) {
        this.requestDto = requestDto;
        if (requestDto == null) {
            log("[ERROR] El objeto LoginRequestDto es nulo");
            return false;
        }

        boolean emailValido = requestDto.getEmail() != null && !requestDto.getEmail().isBlank();
        boolean passValida = requestDto.getPassword() != null && !requestDto.getPassword().isBlank();

        boolean esValido = emailValido && passValida;
        log("[DEBUG] Validación de request: " + esValido);
        return esValido;
    }

    @Override
    public LoginResponseDto obtenerInformacionDelServidor(LoginRequestDto requestDto) {
        if (!validarRequest(requestDto)) {
            log("[ERROR] Request inválido. Abortando.");
            return null;
        }

        String jsonRequest = new LoginRequest(requestDto).toJson();
        log("[DEBUG] JSON a enviar: " + jsonRequest);

        String jsonResponse;
        try {
            jsonResponse = context.executeSend(jsonRequest);
            log("[DEBUG] JSON recibido: " + jsonResponse);
        } catch (Exception e) {
            log("[ERROR] Error al comunicarse con el servidor: " + e.getMessage());
            return null;
        }

        LoginResponse response = new LoginResponse();
        response.fromJson(jsonResponse);
        return response.toDto();
    }

    @Override
    public void cerrarConexion() {
        log("[INFO] Intentando cerrar conexión...");

        ITransportStrategy strategy = context.getStrategy();
        if (strategy instanceof transport.TcpPersistentTransportStrategy tcp) {
            tcp.close();
            log("[INFO] Conexión cerrada correctamente.");
        } else {
            log("[WARN] Esta estrategia no soporta cierre explícito.");
        }
    }

    public TransportContext validarYObtenerInformacion(LoginRequestDto requestDto) {
        log("[INFO] Iniciando login...");

        if (!validarRequest(requestDto)) {
            log("[ERROR] Datos incompletos o inválidos.");
            return null;
        }

        String jsonRequest = new LoginRequest(requestDto).toJson();
        log("[DEBUG] Enviando login: " + jsonRequest);

        String jsonResponse;
        try {
            jsonResponse = context.executeSend(jsonRequest);
            log("[DEBUG] Respuesta recibida: " + jsonResponse);
        } catch (Exception e) {
            log("[ERROR] Fallo al recibir respuesta: " + e.getMessage());
            return null;
        }

        LoginResponse response = new LoginResponse();
        response.fromJson(jsonResponse);

        if (!response.isConnected()) {
            log("[WARN] Login fallido. Cerrando conexión.");
            cerrarConexion();
            return null;
        }
        this.responseDto = response.toDto(); // Aquí se guarda correctamente
        log("[INFO] Login exitoso. Conexión válida.");
        return context;

    }

    private void log(String mensaje) {
        System.out.println(mensaje);
    }
    public LoginResponseDto getResponseDto() {
        return responseDto;
    }

    public void setResponseDto(LoginResponseDto responseDto) {
        this.responseDto = responseDto;
    }

}
