package dto.implementacion.login;

public class LoginResponseDto {
    private String status;
    private String message;
    private Integer id;
    private String nombre;
    private String email;
    private String photo;
    private String ip;
    private String createdAt;
    private boolean isConnected;

    public LoginResponseDto() {}

    public LoginResponseDto(
            String status, String message, Integer id,
            String nombre, String email, String photo,
            String ip, String createdAt, boolean isConnected
    ) {
        this.status = status;
        this.message = message;
        this.id = id;
        this.nombre = nombre;
        this.email = email;
        this.photo = photo;
        this.ip = ip;
        this.createdAt = createdAt;
        this.isConnected = isConnected;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getNombre() {
        return nombre;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPhoto() {
        return photo;
    }

    public void setPhoto(String photo) {
        this.photo = photo;
    }

    public String getIp() {
        return ip;
    }

    public void setIp(String ip) {
        this.ip = ip;
    }

    public String getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(String createdAt) {
        this.createdAt = createdAt;
    }

    public boolean isConnected() {
        return isConnected;
    }

    public void setConnected(boolean connected) {
        isConnected = connected;
    }

    // getters & setters omitted for brevity
}
