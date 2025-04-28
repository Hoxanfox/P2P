package protocolo.implementaciones.register;

import protocolo.interfaces.RequestRoute;

public class RegisterRequest implements RequestRoute {

    private final String username;
    private final String email;
    private final String password;
    private final String photo;
    private final String ip;

    public RegisterRequest(String username, String email, String password, String photo, String ip) {
        this.username = username;
        this.email = email;
        this.password = password;
        this.photo = photo;
        this.ip = ip;
    }

    @Override
    public String toJson() {
        return String.format(
                "{" +
                        "\"command\":\"register\"," +
                        "\"data\":{" +
                        "\"username\":\"%s\"," +
                        "\"email\":\"%s\"," +
                        "\"password\":\"%s\"," +
                        "\"photo\":\"%s\"," +
                        "\"ip\":\"%s\"" +
                        "}" +
                        "}",
                escapeJson(username),
                escapeJson(email),
                escapeJson(password),
                escapeJson(photo),
                escapeJson(ip)
        );
    }

    private String escapeJson(String value) {
        return value.replace("\"", "\\\""); // Por si vienen comillas en los campos
    }
}
