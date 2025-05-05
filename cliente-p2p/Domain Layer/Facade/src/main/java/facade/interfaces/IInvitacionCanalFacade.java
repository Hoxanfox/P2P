package facade.interfaces;

import dto.implementacion.InvitacionCanal.InvitacionCanalRequestDto;
import dto.implementacion.InvitacionCanal.InvitacionCanalResponseDto;
import transport.TransportContext;

public interface IInvitacionCanalFacade {

    /**
     * Obtiene la información desde el servidor a partir del DTO enviado.
     *
     * @param requestDto DTO con los datos para enviar
     * @param context contexto de transporte para realizar la petición
     * @return respuesta del servidor en forma de DTO
     */
    InvitacionCanalResponseDto obtenerInformacionDesdeServidor(InvitacionCanalRequestDto requestDto, TransportContext context);

    /**
     * Persiste la información recibida del servidor.
     *
     * @param responseDto DTO con los datos a persistir
     */
    void persistirInformacion(InvitacionCanalResponseDto responseDto);

    /**
     * Ejecuta todo el flujo: obtiene la información del servidor y luego la persiste.
     *
     * @param requestDto DTO con la solicitud
     * @param context contexto de transporte
     */
    void ejecutarFlujo(InvitacionCanalRequestDto requestDto, TransportContext context);
}
