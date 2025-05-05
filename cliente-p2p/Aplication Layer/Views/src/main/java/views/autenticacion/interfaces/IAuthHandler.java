package views.autenticacion.interfaces;

import dto.implementacion.login.LoginRequestDto;

public interface IAuthHandler {
    void login(LoginRequestDto loginData);
    void register(String email, String password); // puedes definir otro DTO si lo necesitas luego
}
