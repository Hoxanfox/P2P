package protocolo.implementaciones.register;


import dto.implementacion.register.RegisterRequestDto;
import protocolo.interfaces.RequestRoute;

public class RegisterRequest implements RequestRoute {

    private final RegisterRequestDto dto;

    public RegisterRequest(RegisterRequestDto dto) {
        this.dto = dto;
    }

    @Override
    public String toJson() {
        return String.format(
                "{" +
                        "\"command\":\"register\"," +
                        "\"data\":{" +
                        "\"username\":\"%s\"," +  // Se usa "username" aunque en el DTO se llama "nombre"
                        "\"email\":\"%s\"," +
                        "\"password\":\"%s\"," +
                        "\"photo\":\"%s\"," +
                        "\"ip\":\"%s\"" +
                        "}" +
                        "}",
                escapeJson(dto.getNombre()),
                escapeJson(dto.getEmail()),
                escapeJson(dto.getPassword()),
                escapeJson(dto.getFoto()),
                escapeJson(dto.getIp())
        );
    }

    private String escapeJson(String value) {
        return value != null ? value.replace("\"", "\\\"") : "";
    }
}
