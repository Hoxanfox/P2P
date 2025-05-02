package main;

import controller.implementaciones.LoginController;
import controller.implementaciones.ListUsersController;
import dto.implementacion.ListUsers.ListUsersResponseDto;
import dto.implementacion.login.LoginRequestDto;
import facade.implementaciones.ListUsersFacade;
import transport.TransportContext;

public class MainListUsers {
    public static void main(String[] args) {
        // 1. Login para obtener TransportContext
        LoginRequestDto loginDto = new LoginRequestDto();
        loginDto.setEmail("juan.perez@ejeasdasdplo.cm");
        loginDto.setPassword("123adsasd456");

        LoginController loginController = new LoginController();
        TransportContext context = loginController.loguearse(loginDto);

        if (context == null) {
            System.out.println("❌ Login fallido o conexión rechazada.");
            return;
        }

        System.out.println("✅ Login exitoso. Consultando usuarios...");

        // 2. Obtener usuarios con el contexto de conexión
        ListUsersController usersController = new ListUsersController(new ListUsersFacade(), context);
        ListUsersResponseDto response = usersController.obtenerUsuarios();

        // 3. Validar y mostrar resultados
        if ("success".equalsIgnoreCase(response.getStatus())) {
            System.out.println("📋 Lista de usuarios obtenida correctamente:");
            response.getUsuarios().forEach(user -> {
                System.out.println(" - ID: " + user.getId());
                System.out.println("   Nombre: " + user.getNombre());
                System.out.println("   Email: " + user.getEmail());
                System.out.println("   Conectado: " + user.isIs_connected());
                System.out.println("---------------------------");
            });
        } else {
            System.out.println("⚠️ Error: " + response.getMessage());
        }
    }
}
