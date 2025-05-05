package dto.implementacion.InvitacionCanal.Response;

public class DestinatarioResponseDto {
    private Long id;
    private String nombre;

    // Constructor por defecto
    public DestinatarioResponseDto() {
    }

    // Getter y Setter para id
    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    // Getter y Setter para nombre
    public String getNombre() {
        return nombre;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }
}
