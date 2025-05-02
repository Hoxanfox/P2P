package protocolo.implementaciones.listUsers;

import dto.implementacion.ListUsers.ListUsersResponseDto;
import dto.implementacion.ListUsers.UsuarioResponseDTO;
import protocolo.interfaces.ResponseRoute;

import java.util.ArrayList;
import java.util.UUID;

public class ListUsersResponse implements ResponseRoute {
    private final ListUsersResponseDto dto = new ListUsersResponseDto();

    public ListUsersResponseDto getDto() {
        return dto;
    }

    @Override
    public void fromJson(String jsonResponse) {
        try {
            dto.setStatus(extractField(jsonResponse, "status"));
            dto.setMessage(extractField(jsonResponse, "message"));

            if (jsonResponse.contains("\"data\":")) {
                String dataArray = jsonResponse.substring(jsonResponse.indexOf("[") + 1, jsonResponse.lastIndexOf("]"));
                String[] userEntries = dataArray.split("\\},\\s*\\{");

                for (String rawUser : userEntries) {
                    rawUser = rawUser.replace("{", "").replace("}", "").trim();
                    String[] fields = rawUser.split(",");

                    UUID id = null;
                    String nombre = null;
                    String email = null;
                    boolean isConnected = false;

                    for (String field : fields) {
                        String[] keyValue = field.split(":", 2);
                        String key = keyValue[0].replace("\"", "").trim();
                        String value = keyValue[1].replace("\"", "").trim();

                        switch (key) {
                            case "id":
                                id = UUID.fromString(value);
                                break;
                            case "nombre":
                                nombre = value;
                                break;
                            case "email":
                                email = value;
                                break;
                            case "is_connected":
                                isConnected = Boolean.parseBoolean(value);
                                break;
                        }
                    }

                    if (id != null && nombre != null && email != null) {
                        dto.addUsuario(new UsuarioResponseDTO(id, nombre, email, isConnected));
                    }
                }
            }
        } catch (Exception e) {
            dto.setStatus("error");
            dto.setMessage("Error al parsear la respuesta: " + e.getMessage());
            dto.setUsuarios(new ArrayList<>());
        }
    }

    private String extractField(String json, String fieldName) {
        int index = json.indexOf("\"" + fieldName + "\":");
        if (index == -1) return null;

        int start = json.indexOf("\"", index + fieldName.length() + 3) + 1;
        int end = json.indexOf("\"", start);
        return json.substring(start, end);
    }
}
