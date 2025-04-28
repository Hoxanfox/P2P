package protocolo.interfaces;

public interface RequestRoute {
    /**
     * Devuelve el JSON del mensaje a enviar.
     *
     * @return String en formato JSON.
     */
    String toJson();
}
