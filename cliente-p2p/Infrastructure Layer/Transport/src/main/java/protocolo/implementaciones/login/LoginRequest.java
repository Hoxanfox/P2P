package protocolo.implementaciones.login;

import dto.implementacion.login.LoginRequestDto;
import protocolo.interfaces.RequestRoute;

public class LoginRequest implements RequestRoute {

    private final String email;
    private final String password;

    // Constructor que recibe un LoginRequestDto
    public LoginRequest(LoginRequestDto requestDto) {
        this.email = requestDto.getEmail();
        this.password = requestDto.getPassword();
    }

    // Método que convierte el objeto a formato JSON para la solicitud
    @Override
    public String toJson() {
        return String.format(
                "{" +
                        "\"command\":\"login\"," +
                        "\"data\":{" +
                        "\"email\":\"%s\"," +
                        "\"password\":\"%s\"" +
                        "}" +
                        "}",
                escapeJson(email),
                escapeJson(password)
        );
    }

    // Método para escapar las comillas en el valor de los strings para el JSON
    private String escapeJson(String value) {
        return value.replace("\"", "\\\"");
    }
}
