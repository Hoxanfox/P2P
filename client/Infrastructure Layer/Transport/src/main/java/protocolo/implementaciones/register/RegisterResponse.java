package protocolo.implementaciones.register;

import protocolo.interfaces.ResponseRoute;

public class RegisterResponse implements ResponseRoute {

    private String status;
    private String message;
    private String userId; // Ahora UUID en formato string
    private String username;

    @Override
    public void fromJson(String jsonResponse) {
        System.out.println("[DEBUG] Respuesta recibida: " + jsonResponse);

        // Asignar status
        if (jsonResponse.contains("\"status\":\"success\"")) {
            status = "success";
        } else {
            status = "error";
        }

        // Extraer message
        message = extractValue(jsonResponse, "message");

        // Extraer userId como UUID (string)
        userId = extractValue(jsonResponse, "userId");

        // Extraer username
        username = extractValue(jsonResponse, "username");
    }

    private String extractValue(String json, String key) {
        try {
            String search = "\"" + key + "\":";
            int index = json.indexOf(search);
            if (index == -1) {
                return null;
            }

            int start = index + search.length();

            // Verificar si es String o número
            if (json.charAt(start) == '\"') {
                start++; // Saltar comilla inicial
                int end = json.indexOf("\"", start);
                return json.substring(start, end);
            } else {
                int end = json.indexOf(",", start);
                if (end == -1) {
                    end = json.indexOf("}", start);
                }
                return json.substring(start, end).trim();
            }

        } catch (Exception e) {
            System.out.println("[ERROR] No se pudo extraer la key: " + key);
            return null;
        }
    }

    // Getters
    public String getStatus() {
        return status;
    }

    public String getMessage() {
        return message;
    }

    public String getUserId() {
        return userId;
    }

    public String getUsername() {
        return username;
    }
}
