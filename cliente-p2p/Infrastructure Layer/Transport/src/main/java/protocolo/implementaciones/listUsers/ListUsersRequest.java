package protocolo.implementaciones.listUsers;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ObjectNode;
import dto.implementacion.ListUsers.ListUsersRequestDto;
import protocolo.interfaces.RequestRoute;

public class ListUsersRequest implements RequestRoute {
    private final ListUsersRequestDto dto;

    public ListUsersRequest() {
        this.dto = new ListUsersRequestDto("list-users");
        System.out.println("[ListUsersRequest] Constructor sin parámetros, dto inicializado con: " + dto.getCommand());
    }

    public ListUsersRequest(ListUsersRequestDto dto) {
        this.dto = dto;
        System.out.println("[ListUsersRequest] Constructor con parámetros, dto recibido con: " + dto.getCommand());
    }

    public ListUsersRequestDto getDto() {
        return dto;
    }

    @Override
    public String toJson() {
        ObjectMapper objectMapper = new ObjectMapper();

        try {
            // Crear manualmente el objeto con la estructura requerida
            ObjectNode root = objectMapper.createObjectNode();
            root.put("command", dto.getCommand());
            root.set("data", objectMapper.createObjectNode()); // objeto vacío

            String json = objectMapper.writeValueAsString(root);

            System.out.println("[ListUsersRequest] JSON estructurado: " + json);
            return json;

        } catch (JsonProcessingException e) {
            System.err.println("[ListUsersRequest] Error al generar el JSON.");
            e.printStackTrace();
            return "{}";
        }
    }
}
