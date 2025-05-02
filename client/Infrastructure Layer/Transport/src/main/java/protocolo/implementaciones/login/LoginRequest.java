package protocolo.implementaciones.login;

import protocolo.interfaces.RequestRoute;

public class LoginRequest implements RequestRoute {

    private final String email;
    private final String password;

    public LoginRequest(String email, String password) {
        this.email = email;
        this.password = password;
    }

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

    private String escapeJson(String value) {
        return value.replace("\"", "\\\"");
    }
}
