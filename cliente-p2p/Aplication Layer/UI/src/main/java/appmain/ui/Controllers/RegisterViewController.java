package appmain.ui.Controllers;

import javafx.fxml.FXML;
import javafx.scene.control.*;
import javafx.stage.Stage;
import appmain.ui.MainViewController;

import dto.implementacion.register.RegisterRequestDto;
import controller.implementaciones.RegisterController;

import java.net.InetAddress;
import java.net.UnknownHostException;


public class RegisterViewController {

    @FXML
    private TextField usernameField;
    @FXML
    private TextField emailField;
    @FXML
    private PasswordField passwordField;
    @FXML
    private PasswordField confirmPasswordField;
    @FXML
    private Button registerButton;
    @FXML
    private Button backButton;
    @FXML

    private MainViewController mainController;
    private RegisterController registerController;

    @FXML
    public void initialize() {
        // Configurar eventos
        setupButtons();

        registerController = new RegisterController();
    }

    private void setupButtons() {
        registerButton.setOnAction(event -> handleRegister());
        backButton.setOnAction(event -> closeWindow());
    }

    public static String getLocalIpAddress() {
        try {
            InetAddress inetAddress = InetAddress.getLocalHost();
            return inetAddress.getHostAddress(); // Devuelve la IP en formato String
        } catch (UnknownHostException e) {
            e.printStackTrace();
            return "No se pudo obtener la dirección IP";
        }
    }

    private void handleRegister() {
        String username = usernameField.getText().trim();
        String email = emailField.getText().trim();
        String password = passwordField.getText();
        String confirmPassword = confirmPasswordField.getText();

        if (username.isEmpty() || email.isEmpty() || password.isEmpty() || confirmPassword.isEmpty()) {
            showAlert("Error", "Todos los campos son obligatorios.", Alert.AlertType.ERROR);
            return;
        }

        if (!password.equals(confirmPassword)) {
            showAlert("Error", "Las contraseñas no coinciden.", Alert.AlertType.ERROR);
            return;
        }

        RegisterRequestDto registerRequestDto = new RegisterRequestDto();
        registerRequestDto.setNombre(username);
        registerRequestDto.setEmail(email);
        registerRequestDto.setPassword(password);
        registerRequestDto.setFoto("testImg");
        registerRequestDto.setIp(getLocalIpAddress());

        //TODO: Corregir, no se recibió respuesta del servidor
        //registerController.registrar(registerRequestDto);
        //TODO: Ahora que? como sé si fue exitso?

        //Mostrar Alerta
        showAlert("Registrado", "Registrado exitosamente.", Alert.AlertType.INFORMATION );
        closeWindow();
    }

    private void openLoginForm() {
        closeWindow();
        mainController.openLoginForm();
    }

    private void showAlert(String title, String message, Alert.AlertType type) {
        Alert alert = new Alert(type);
        alert.setTitle(title);
        alert.setHeaderText(null);
        alert.setContentText(message);
        alert.showAndWait();
    }

    private void closeWindow() {
        Stage stage = (Stage) registerButton.getScene().getWindow();
        stage.close();
    }

    public void setMainController(MainViewController controller) {
        this.mainController = controller;
    }
}
