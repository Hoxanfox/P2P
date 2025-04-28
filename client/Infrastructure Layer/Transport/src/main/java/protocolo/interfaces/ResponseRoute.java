package protocolo.interfaces;

public interface ResponseRoute {
    /**
     * Procesa la respuesta JSON recibida.
     *
     * @param jsonResponse El JSON recibido como respuesta.
     */
    void fromJson(String jsonResponse);
}