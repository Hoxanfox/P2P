package views.autenticacion.implementaciones;

import controller.implementaciones.LoginController;
import dto.implementacion.login.LoginRequestDto;
import dto.implementacion.login.LoginResponseDto;
import transport.TransportContext;
import transport.TcpPersistentTransportStrategy;
import views.autenticacion.interfaces.IAuthHandler;
import views.autenticacion.mainView.DashboardView;

import javax.swing.*;

public class AuthHandlerImpl implements IAuthHandler {

    private final LoginController loginController;
    private final TransportContext context = new TransportContext(
            new TcpPersistentTransportStrategy("localhost", 9000)
    );

    public AuthHandlerImpl() {
        this.loginController = new LoginController(context);
    }

    @Override
    public void login(LoginRequestDto loginData) {
        System.out.println("[DEBUG] Iniciando login en hilo separado...");

        new Thread(() -> {
            try {
                System.out.println("[DEBUG] Llamando a loguearse...");
                TransportContext resultado = loginController.loguearse(loginData);

                SwingUtilities.invokeLater(() -> {
                    if (resultado != null) {
                        // Obtener el DTO de la respuesta desde el controlador
                        LoginResponseDto responseDto = loginController.getResponseDto();

                        if (responseDto != null && responseDto.isConnected()) {
                            System.out.println("[DEBUG] Login exitoso. Usuario: " + responseDto.getNombre());
                            JOptionPane.showMessageDialog(null, "Bienvenido " + responseDto.getNombre());

                            // Lanzar vista principal con el DTO
                            new DashboardView(resultado);
                        } else {
                            JOptionPane.showMessageDialog(null, "Credenciales incorrectas", "Error", JOptionPane.ERROR_MESSAGE);
                        }

                    } else {
                        JOptionPane.showMessageDialog(null, "Error al conectar con el servidor", "Error", JOptionPane.ERROR_MESSAGE);
                    }
                });

            } catch (Exception e) {
                e.printStackTrace();
                SwingUtilities.invokeLater(() ->
                        JOptionPane.showMessageDialog(null, "Error inesperado durante el login", "Error", JOptionPane.ERROR_MESSAGE)
                );
            }
        }).start();
    }

    @Override
    public void register(String email, String password) {
        // Aquí irá la lógica de registro cuando la implementes
    }
}
