package protocolo.implementaciones.login;

import protocolo.interfaces.ResponseRoute;
import dto.implementacion.login.LoginResponseDto;

public class LoginResponse implements ResponseRoute {

    private String status;
    private String message;
    private String id;
    private String nombre;
    private String email;
    private String photo;
    private String ip;
    private String createdAt;
    private boolean isConnected;

    @Override
    public void fromJson(String json) {
        if (json == null || json.isBlank()) {
            System.out.println("[ERROR] JSON nulo o vacío en LoginResponse");
            return;
        }

        System.out.println("[INFO] Iniciando parseo de JSON en LoginResponse");

        this.status = extractSimple(json, "status");
        this.message = extractSimple(json, "message");

        System.out.println("[DEBUG] status: " + status);
        System.out.println("[DEBUG] message: " + message);

        if ("success".equalsIgnoreCase(status)) {
            String dataJson = extractBlock(json, "data");

            if (dataJson != null) {
                System.out.println("[INFO] Bloque 'data' extraído correctamente");
                this.id          = extractSimple(dataJson, "id");
                this.nombre      = extractSimple(dataJson, "nombre");
                this.email       = extractSimple(dataJson, "email");
                this.photo       = extractSimple(dataJson, "photo");
                this.ip          = extractSimple(dataJson, "ip");
                this.createdAt   = extractSimple(dataJson, "created_at");
                String connStr   = extractSimple(dataJson, "is_connected");
                this.isConnected = Boolean.parseBoolean(connStr);

                System.out.println("[DEBUG] id: " + id);
                System.out.println("[DEBUG] nombre: " + nombre);
                System.out.println("[DEBUG] email: " + email);
                System.out.println("[DEBUG] photo: " + photo);
                System.out.println("[DEBUG] ip: " + ip);
                System.out.println("[DEBUG] createdAt: " + createdAt);
                System.out.println("[DEBUG] isConnected: " + isConnected);
            } else {
                System.out.println("[ERROR] No se pudo extraer el bloque 'data' del JSON");
            }
        } else {
            System.out.println("[INFO] Estado no es 'success', no se procesarán más campos");
        }
    }

    private String extractSimple(String json, String key) {
        String pattern = "\"" + key + "\":";
        int index = json.indexOf(pattern);
        if (index == -1) {
            System.out.println("[WARN] Clave '" + key + "' no encontrada en el JSON");
            return null;
        }

        index += pattern.length();
        while (index < json.length() && Character.isWhitespace(json.charAt(index))) index++;

        if (index >= json.length()) {
            System.out.println("[WARN] Fin del JSON alcanzado al buscar valor de '" + key + "'");
            return null;
        }

        char ch = json.charAt(index);
        if (ch == '"') {
            int end = json.indexOf('"', index + 1);
            if (end > index) {
                String result = json.substring(index + 1, end);
                System.out.println("[DEBUG] Valor extraído de '" + key + "': " + result);
                return result;
            }
        } else {
            int end = findEndOfPrimitive(json, index);
            if (end > index) {
                String result = json.substring(index, end).trim();
                System.out.println("[DEBUG] Valor extraído de '" + key + "': " + result);
                return result;
            }
        }

        System.out.println("[WARN] No se pudo extraer valor para la clave '" + key + "'");
        return null;
    }

    private int findEndOfPrimitive(String json, int start) {
        int end = start;
        while (end < json.length()) {
            char c = json.charAt(end);
            if (c == ',' || c == '}' || Character.isWhitespace(c)) break;
            end++;
        }
        return end;
    }

    private String extractBlock(String json, String key) {
        String pattern = "\"" + key + "\":";
        int start = json.indexOf(pattern);
        if (start == -1) {
            System.out.println("[WARN] Bloque '" + key + "' no encontrado");
            return null;
        }

        int braceStart = json.indexOf('{', start);
        if (braceStart == -1) {
            System.out.println("[WARN] No se encontró apertura de bloque '{' para '" + key + "'");
            return null;
        }

        int braceCount = 0;
        for (int i = braceStart; i < json.length(); i++) {
            char c = json.charAt(i);
            if (c == '{') braceCount++;
            else if (c == '}') braceCount--;

            if (braceCount == 0) {
                String block = json.substring(braceStart, i + 1);
                System.out.println("[DEBUG] Bloque extraído para '" + key + "': " + block);
                return block;
            }
        }

        System.out.println("[ERROR] No se cerró correctamente el bloque para '" + key + "'");
        return null;
    }

    public LoginResponseDto toDto() {
        System.out.println("[INFO] Convirtiendo LoginResponse a DTO");
        return new LoginResponseDto(
                status,
                message,
                id != null ? Integer.valueOf(id) : null,
                nombre,
                email,
                photo,
                ip,
                createdAt,
                isConnected
        );
    }

    public boolean isConnected() {
        return isConnected;
    }
}
