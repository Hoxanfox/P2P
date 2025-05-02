package protocolo.implementaciones.register;

import protocolo.interfaces.ResponseRoute;
import dto.implementacion.register.RegisterResponseDto;

import java.time.LocalDate;
import java.util.Base64;
import java.util.UUID;

public class RegisterResponse implements ResponseRoute {

    private String status;
    private String message;
    private RegisterResponseDto data;

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
        message = extractValue(jsonResponse, "message").replaceAll("\"", "");

        // Solo si success y existe data, parsear DTO completo
        if ("success".equals(status) && jsonResponse.contains("\"data\":")) {
            // Extraer bloque data
            String dataBlock = extractDataBlock(jsonResponse);

            // Campos del dataBlock
            String idStr         = extractValue(dataBlock, "id");
            String nombre        = extractValue(dataBlock, "nombre");
            String email         = extractValue(dataBlock, "email");
            String password      = extractValue(dataBlock, "password");
            String fotoBase64    = extractValue(dataBlock, "foto");
            String ip            = extractValue(dataBlock, "ip");
            String estadoStr     = extractValue(dataBlock, "estado");
            String fechaStr      = extractValue(dataBlock, "fechaRegistro");

            // Limpiar comillas
            idStr      = trimQuotes(idStr);
            nombre     = trimQuotes(nombre);
            email      = trimQuotes(email);
            password   = trimQuotes(password);
            fotoBase64 = trimQuotes(fotoBase64);
            ip         = trimQuotes(ip);
            estadoStr  = trimQuotes(estadoStr);
            fechaStr   = trimQuotes(fechaStr);

            // Convertir tipos
            UUID id = UUID.fromString(idStr);
            byte[] foto = fotoBase64 != null && !fotoBase64.isEmpty()
                    ? Base64.getDecoder().decode(fotoBase64)
                    : null;
            boolean estado = "true".equalsIgnoreCase(estadoStr) || "1".equals(estadoStr);
            LocalDate fechaRegistro = fechaStr != null && !fechaStr.isEmpty()
                    ? LocalDate.parse(fechaStr)
                    : null;

            // Crear DTO
            data = new RegisterResponseDto(id, nombre, email, password, foto, ip, estado, fechaRegistro);
        }
    }

    /** Extrae el bloque JSON de data incluyendo llaves {} */
    private String extractDataBlock(String json) {
        int idx = json.indexOf("\"data\":");
        int start = json.indexOf('{', idx + 6);
        int braceCount = 0;
        for (int i = start; i < json.length(); i++) {
            if (json.charAt(i) == '{') braceCount++;
            else if (json.charAt(i) == '}') {
                braceCount--;
                if (braceCount == 0) {
                    return json.substring(start, i + 1);
                }
            }
        }
        return "";
    }

    /**
     * Extrae valor raw (con comillas si es string) para clave dada
     */
    private String extractValue(String json, String key) {
        String search = "\"" + key + "\":";
        int index = json.indexOf(search);
        if (index == -1) return null;
        int start = index + search.length();
        // saltar espacios
        while (start < json.length() && Character.isWhitespace(json.charAt(start))) start++;
        if (json.charAt(start) == '"') {
            int end = json.indexOf('"', start + 1);
            return json.substring(start, end + 1);
        } else {
            int end = json.indexOf(',', start);
            if (end == -1) end = json.indexOf('}', start);
            return json.substring(start, end);
        }
    }

    private String trimQuotes(String value) {
        if (value == null) return null;
        if (value.startsWith("\"") && value.endsWith("\"")) {
            return value.substring(1, value.length() - 1);
        }
        return value;
    }

    // Getters
    public String getStatus() {
        return status;
    }

    public String getMessage() {
        return message;
    }

    public RegisterResponseDto getData() {
        return data;
    }
}
