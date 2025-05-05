package views.autenticacion.login;

import dto.implementacion.login.LoginRequestDto;
import views.autenticacion.interfaces.IAuthHandler;

import javax.swing.*;
import java.awt.*;

public class LoginPanel extends JPanel {
    public LoginPanel(IAuthHandler authHandler, Runnable switchToRegister) {
        setLayout(new GridBagLayout());
        GridBagConstraints gbc = new GridBagConstraints();
        gbc.insets = new Insets(5, 10, 5, 10);

        JLabel emailLabel = new JLabel("Correo:");
        JTextField emailField = new JTextField(20);

        JLabel passwordLabel = new JLabel("Contraseña:");
        JPasswordField passwordField = new JPasswordField(20);

        JButton loginButton = new JButton("Iniciar sesión");
        JButton registerLink = new JButton("Registrarse");

        loginButton.addActionListener(e -> {
            LoginRequestDto dto = new LoginRequestDto(
                    emailField.getText(),
                    new String(passwordField.getPassword())
            );
            authHandler.login(dto);
        });

        registerLink.addActionListener(e -> switchToRegister.run());

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
        add(loginButton, gbc);

        gbc.gridy = 3;
        add(registerLink, gbc);
    }
}
