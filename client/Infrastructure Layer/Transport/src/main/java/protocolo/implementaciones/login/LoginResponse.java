package protocolo.implementaciones.login;

import protocolo.interfaces.ResponseRoute;

public class LoginResponse implements ResponseRoute {

    private String status;
    private String message;

    // Campos si la respuesta es exitosa
    private String id; // UUID en formato string
    private String nombre;
    private String email;
    private String photo;
    private String ip;
    private String createdAt;
    private boolean isConnected;

    @Override
    public void fromJson(String jsonResponse) {
        this.status = extractValue(jsonResponse, "status");
        this.message = extractValue(jsonResponse, "message");

        if ("success".equals(status)) {
            this.id = extractValue(jsonResponse, "id");
            this.nombre = extractValue(jsonResponse, "nombre");
            this.email = extractValue(jsonResponse, "email");
            this.photo = extractValue(jsonResponse, "photo");
            this.ip = extractValue(jsonResponse, "ip");
            this.createdAt = extractValue(jsonResponse, "created_at");
            this.isConnected = Boolean.parseBoolean(extractValue(jsonResponse, "is_connected"));
        }
    }

    private String extractValue(String json, String key) {
        String pattern = "\"" + key + "\":";
        int index = json.indexOf(pattern);
        if (index == -1) return null;

        index += pattern.length();

        // Salta espacios
        while (index < json.length() && Character.isWhitespace(json.charAt(index))) {
            index++;
        }

        char startChar = json.charAt(index);
        if (startChar == '"') {
            index++;
            int endIndex = json.indexOf('"', index);
            return json.substring(index, endIndex);
        } else if (startChar == 't' || startChar == 'f') {
            // booleano
            int endIndex = json.indexOf(",", index);
            if (endIndex == -1) endIndex = json.indexOf("}", index);
            return json.substring(index, endIndex).trim();
        } else {
            // número o UUID (sin comillas, pero puede contener guiones)
            int endIndex = json.indexOf(",", index);
            if (endIndex == -1) endIndex = json.indexOf("}", index);
            return json.substring(index, endIndex).replace("\"", "").trim();
        }
    }

    // Getters

    public String getStatus() {
        return status;
    }

    public String getMessage() {
        return message;
    }

    public String getId() {
        return id;
    }

    public String getNombre() {
        return nombre;
    }

    public String getEmail() {
        return email;
    }

    public String getPhoto() {
        return photo;
    }

    public String getIp() {
        return ip;
    }

    public String getCreatedAt() {
        return createdAt;
    }

    public boolean isConnected() {
        return isConnected;
    }
}
