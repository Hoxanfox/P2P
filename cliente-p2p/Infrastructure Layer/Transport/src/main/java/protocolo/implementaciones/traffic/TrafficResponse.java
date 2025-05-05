package protocolo.implementaciones.traffic;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import dto.implementacion.Traffic.*;

import dto.implementacion.Traffic.RefreshUser.RefreshUsersDataDTO;
import protocolo.interfaces.ResponseRoute;

public class TrafficResponse implements ResponseRoute {

    private TrafficResponseDto responseDTO;

    public TrafficResponseDto getResponseDTO() {
        return responseDTO;
    }

    @Override
    public void fromJson(String jsonResponse) {
        try {
            ObjectMapper mapper = new ObjectMapper();
            JsonNode root = mapper.readTree(jsonResponse);

            String command = root.get("command").asText();
            JsonNode dataNode = root.get("data");

            Object dataDto;

            switch (command) {
                case "refresh-users":
                    dataDto = mapper.treeToValue(dataNode, RefreshUsersDataDTO.class);
                    break;
                // Aquí podrías agregar más comandos como:
                // case "update-profile":
                //     dataDto = mapper.treeToValue(dataNode, UpdateProfileDTO.class);
                //     break;
                default:
                    dataDto = null;
                    break;
            }

            this.responseDTO = new TrafficResponseDto(command, dataDto);

        } catch (Exception e) {
            e.printStackTrace();
            this.responseDTO = null;
        }
    }

    @Override
    public String toString() {
        return (responseDTO != null) ? responseDTO.toString() : "Respuesta vacía o inválida";
    }
}
