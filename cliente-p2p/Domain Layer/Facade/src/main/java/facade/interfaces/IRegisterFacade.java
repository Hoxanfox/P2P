package facade.interfaces;

import dto.implementacion.register.RegisterRequestDto;
import dto.implementacion.register.RegisterResponseDto;

public interface IRegisterFacade {

    // Obtener datos del servidor usando un DTO de entrada
    RegisterResponseDto obtenerInformacionDelServidor(RegisterRequestDto requestDto);

    // Persistir los datos recibidos
    boolean persistirUsuario(RegisterResponseDto responseDto);

    // Ejecutar todo el flujo: obtener + persistir
    void ejecutarFlujoRegistro(RegisterRequestDto requestDto);
}
