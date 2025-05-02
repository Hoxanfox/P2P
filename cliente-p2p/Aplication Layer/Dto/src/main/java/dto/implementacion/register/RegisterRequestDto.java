package dto.implementacion.register;

public class RegisterRequestDto {
    private String nombre;
    private String email;
    private String password;
    private String foto; // En Base64
    private String ip;

    public RegisterRequestDto() {}

    public RegisterRequestDto(String nombre, String email, String password, String foto, String ip) {
        this.nombre = nombre;
        this.email = email;
        this.password = password;
        this.foto = foto;
        this.ip = ip;
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

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getFoto() {
        return foto;
    }

    public void setFoto(String foto) {
        this.foto = foto;
    }

    public String getIp() {
        return ip;
    }

    public void setIp(String ip) {
        this.ip = ip;
    }
}
