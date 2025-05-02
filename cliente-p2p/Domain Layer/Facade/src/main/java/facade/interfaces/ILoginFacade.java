// src/main/java/facade/interfaces/ILoginFacade.java
package facade.interfaces;

import dto.implementacion.login.LoginRequestDto;
import dto.implementacion.login.LoginResponseDto;

public interface ILoginFacade {

    /**
     * Valida que el DTO de request tenga todos los campos mínimos
     * necesarios para enviar la petición de login.
     *
     * @param requestDto el DTO de petición de login
     * @return true si es válido; false en caso contrario
     */
    boolean validarRequest(LoginRequestDto requestDto);

    /**
     * Envía la petición al servidor y parsea la respuesta en un DTO.
     *
     * @param requestDto el DTO con email y password
     * @return un LoginResponseDto con los datos devueltos por el servidor,
     *         o null si hubo error o status="error"
     */
    LoginResponseDto obtenerInformacionDelServidor(LoginRequestDto requestDto);

    /**
     * Si la autenticación es exitosa, deja la conexión abierta;
     * si la autenticación falla o se solicita cerrar, cierra la conexión.
     */
    void cerrarConexion();
}
