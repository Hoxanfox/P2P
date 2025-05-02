package main;

import dto.implementacion.register.RegisterRequestDto;
import facade.implementaciones.RegisterFacade;
import transport.TcpTransportStrategy;
import java.util.Base64;
import java.net.InetAddress;

public class MainRegistro {

    public static void main(String[] args) {
        try {
            // Crear el DTO con datos de prueba
            RegisterRequestDto requestDto = new RegisterRequestDto();
            requestDto.setNombre("Juanasdasd");
            requestDto.setEmail("juan.perez@ejeasdasdplo.cm");
            requestDto.setPassword("123adsasd456"); // ejemplo
            requestDto.setFoto(Base64.getEncoder().encodeToString("aG9sYQ==".getBytes())); // simula una imagen
            requestDto.setIp(InetAddress.getLocalHost().getHostAddress());

            // Instanciar la fachada con el transporte TCP
            RegisterFacade facade = new RegisterFacade(new TcpTransportStrategy("localhost", 9000));

            // Ejecutar el flujo completo
            facade.ejecutarFlujoRegistro(requestDto);

        } catch (Exception e) {
            System.out.println("[ERROR] Fallo en ejecución del main: " + e.getMessage());
            e.printStackTrace();
        }
    }
}
