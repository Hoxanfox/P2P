package controller.implementaciones;

import dto.implementacion.login.LoginRequestDto;
import dto.implementacion.login.LoginResponseDto;
import facade.implementaciones.LoginFacade;
import transport.TransportContext;

public class LoginController {

    private final LoginFacade loginFacade;
    private TransportContext context;
    private LoginResponseDto responseDto;
    private LoginRequestDto requestDto;

    public LoginController(TransportContext context) {
        this.context = context;
        this.loginFacade = new LoginFacade(context);
    }

    public TransportContext loguearse(LoginRequestDto loginData) {
        this.requestDto = loginData; // Guarda localmente el request
        TransportContext resultado = loginFacade.validarYObtenerInformacion(loginData);
        this.responseDto = loginFacade.getResponseDto(); // Captura el response desde el facade
        return resultado;
    }

    // Getters y Setters

    public TransportContext getContext() {
        return context;
    }

    public void setContext(TransportContext context) {
        this.context = context;
    }

    public LoginResponseDto getResponseDto() {
        return responseDto;
    }

    public void setResponseDto(LoginResponseDto responseDto) {
        this.responseDto = responseDto;
    }

    public LoginRequestDto getRequestDto() {
        return requestDto;
    }

    public void setRequestDto(LoginRequestDto requestDto) {
        this.requestDto = requestDto;
    }
}
