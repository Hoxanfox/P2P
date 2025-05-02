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
    // getters & setters omitted for brevity
}
