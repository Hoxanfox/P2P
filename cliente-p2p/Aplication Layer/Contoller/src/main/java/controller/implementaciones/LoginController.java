package controller.implementaciones;

import dto.implementacion.login.LoginRequestDto;
import facade.implementaciones.LoginFacade;
import transport.TcpPersistentTransportStrategy;
import transport.TransportContext;

public class LoginController {

    private final LoginFacade loginFacade;

    public LoginController() {
        // Instancia la fachada con una estrategia de transporte TCP
        this.loginFacade = new LoginFacade(new TcpPersistentTransportStrategy("localhost", 9000));
    }

    public TransportContext loguearse(LoginRequestDto requestDto) {
        return loginFacade.validarYObtenerInformacion(requestDto);
    }
}
