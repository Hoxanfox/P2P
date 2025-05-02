package controller.implementaciones;

import dto.implementacion.register.RegisterRequestDto;
import facade.implementaciones.RegisterFacade;
import transport.TcpTransportStrategy;

public class RegisterController {

    private final RegisterFacade facade;

    public RegisterController() {
        // Estrategia TCP apuntando al host/puerto de tu servidor
        this.facade = new RegisterFacade(new TcpTransportStrategy("localhost", 9000));
    }

    /**
     * Recibe el DTO de petición, ejecuta todo el flujo:
     *   1) Llamada al servidor remoto (REGISTER)
     *   2) Persiste el usuario localmente si el servidor respondió "success"
     *
     * @param requestDto DTO con los datos para registrar al usuario
     */
    public void registrar(RegisterRequestDto requestDto) {
        if (requestDto == null) {
            System.err.println("[ERROR] DTO de petición nulo. Abortando registro.");
            return;
        }

        System.out.println("[INFO] Controlador: iniciando registro de " + requestDto.getNombre());
        facade.ejecutarFlujoRegistro(requestDto);
    }
}
