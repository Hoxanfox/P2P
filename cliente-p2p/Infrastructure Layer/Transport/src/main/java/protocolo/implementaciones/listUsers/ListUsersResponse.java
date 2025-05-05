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
            System.out.println("[INFO] Comenzando a procesar la respuesta JSON...");
            dto.setStatus(extractField(jsonResponse, "status"));
            System.out.println("[DEBUG] Estado extraído: " + dto.getStatus());

            dto.setMessage(extractField(jsonResponse, "message"));
            System.out.println("[DEBUG] Mensaje extraído: " + dto.getMessage());

            if (jsonResponse.contains("\"data\":")) {
                String dataArray = jsonResponse.substring(jsonResponse.indexOf("[") + 1, jsonResponse.lastIndexOf("]"));
                System.out.println("[DEBUG] Array de datos extraído: " + dataArray);

                String[] userEntries = dataArray.split("\\},\\s*\\{");
                System.out.println("[DEBUG] Número de usuarios encontrados: " + userEntries.length);

                for (String rawUser : userEntries) {
                    rawUser = rawUser.replace("{", "").replace("}", "").trim();
                    System.out.println("[DEBUG] Usuario procesado: " + rawUser);

                    String[] fields = rawUser.split(",");
                    UUID id = null;
                    String nombre = null;
                    String email = null;
                    boolean isConnected = false;

                    for (String field : fields) {
                        String[] keyValue = field.split(":", 2);
                        String key = keyValue[0].replace("\"", "").trim();
                        String value = keyValue[1].replace("\"", "").trim();

                        System.out.println("[DEBUG] Clave: " + key + " | Valor: " + value);

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
                        System.out.println("[INFO] Usuario válido encontrado. Agregando...");
                        dto.addUsuario(new UsuarioResponseDTO(id, nombre, email, isConnected));
                    } else {
                        System.out.println("[WARNING] Usuario inválido, faltan campos.");
                    }
                }
            } else {
                System.out.println("[INFO] No se encontraron datos de usuarios en la respuesta.");
            }
        } catch (Exception e) {
            System.out.println("[ERROR] Excepción durante el parseo de la respuesta: " + e.getMessage());
            dto.setStatus("error");
            dto.setMessage("Error al parsear la respuesta: " + e.getMessage());
            dto.setUsuarios(new ArrayList<>());
        }
    }

    private String extractField(String json, String fieldName) {
        System.out.println("[DEBUG] Extrayendo campo: " + fieldName);
        int index = json.indexOf("\"" + fieldName + "\":");
        if (index == -1) return null;

        int start = json.indexOf("\"", index + fieldName.length() + 3) + 1;
        int end = json.indexOf("\"", start);
        String fieldValue = json.substring(start, end);
        System.out.println("[DEBUG] Valor del campo " + fieldName + ": " + fieldValue);
        return fieldValue;
    }
}
