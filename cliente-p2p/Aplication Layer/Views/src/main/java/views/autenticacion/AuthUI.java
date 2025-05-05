package views.autenticacion;

import views.autenticacion.interfaces.IAuthHandler;
import views.autenticacion.login.LoginPanel;
import views.autenticacion.registrar.RegisterPanel;
import views.autenticacion.implementaciones.AuthHandlerImpl; // Corrección directa aquí

import javax.swing.*;
import java.awt.*;

public class AuthUI extends JFrame {

    public AuthUI() {
        System.out.println("[DEBUG] Creando la ventana de autenticación...");
        setTitle("Autenticación");
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setSize(400, 250);
        setLocationRelativeTo(null);

        // Crear el manejador de autenticación
        IAuthHandler handler = new AuthHandlerImpl(); // Usa el AuthHandler implementado
        System.out.println("[DEBUG] Handler de autenticación creado.");

        JPanel mainPanel = new JPanel(new CardLayout());
        System.out.println("[DEBUG] Panel principal creado con CardLayout.");

        // Crear el panel de inicio de sesión (Login)
        LoginPanel loginPanel = new LoginPanel(handler, () -> {
            System.out.println("[DEBUG] Se presionó 'Registrarse', cambiando a la vista de registro.");
            CardLayout cl = (CardLayout) mainPanel.getLayout();
            cl.show(mainPanel, "register");
        });

        // Crear el panel de registro (Register)
        RegisterPanel registerPanel = new RegisterPanel(handler, () -> {
            System.out.println("[DEBUG] Se presionó 'Iniciar sesión', cambiando a la vista de login.");
            CardLayout cl = (CardLayout) mainPanel.getLayout();
            cl.show(mainPanel, "login");
        });

        // Agregar los paneles al panel principal
        mainPanel.add(loginPanel, "login");
        mainPanel.add(registerPanel, "register");
        System.out.println("[DEBUG] Paneles 'login' y 'register' agregados al panel principal.");

        // Establecer el panel principal en el JFrame
        add(mainPanel);
        System.out.println("[DEBUG] Panel principal añadido a la ventana.");

        // Mostrar la vista de login por defecto
        ((CardLayout) mainPanel.getLayout()).show(mainPanel, "login");
        System.out.println("[DEBUG] Mostrando la vista de login.");
    }

    public static void main(String[] args) {
        System.out.println("[DEBUG] Iniciando la aplicación de autenticación...");
        SwingUtilities.invokeLater(() -> new AuthUI().setVisible(true));
    }
}
