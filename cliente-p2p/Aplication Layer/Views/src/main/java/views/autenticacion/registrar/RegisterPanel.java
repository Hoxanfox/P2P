package views.autenticacion.registrar;

import views.autenticacion.interfaces.IAuthHandler;

import javax.swing.*;
import java.awt.*;

public class RegisterPanel extends JPanel {
    public RegisterPanel(IAuthHandler authHandler, Runnable switchToLogin) {
        setLayout(new GridBagLayout());
        GridBagConstraints gbc = new GridBagConstraints();
        gbc.insets = new Insets(5, 10, 5, 10);

        JLabel emailLabel = new JLabel("Correo:");
        JTextField emailField = new JTextField(20);

        JLabel passwordLabel = new JLabel("Contraseña:");
        JPasswordField passwordField = new JPasswordField(20);

        JButton registerButton = new JButton("Registrarse");
        JButton loginLink = new JButton("Volver al login");

        registerButton.addActionListener(e ->
                authHandler.register(emailField.getText(), new String(passwordField.getPassword()))
        );

        loginLink.addActionListener(e -> switchToLogin.run());

        gbc.gridx = 0; gbc.gridy = 0; gbc.anchor = GridBagConstraints.LINE_END;
        add(emailLabel, gbc);
        gbc.gridx = 1; gbc.anchor = GridBagConstraints.LINE_START;
        add(emailField, gbc);

        gbc.gridx = 0; gbc.gridy = 1; gbc.anchor = GridBagConstraints.LINE_END;
        add(passwordLabel, gbc);
        gbc.gridx = 1; gbc.anchor = GridBagConstraints.LINE_START;
        add(passwordField, gbc);

        gbc.gridy = 2; gbc.gridx = 0; gbc.gridwidth = 2;
        gbc.anchor = GridBagConstraints.CENTER;
        add(registerButton, gbc);

        gbc.gridy = 3;
        add(loginLink, gbc);
    }
}
