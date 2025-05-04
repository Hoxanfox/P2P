package appmain.ui.Controllers;

import javafx.fxml.FXML;
import javafx.scene.control.*;
import javafx.stage.Stage;
import appmain.ui.MainViewController;

import controller.implementaciones.LoginController;
import dto.implementacion.login.LoginRequestDto;

public class LoginViewController {

    @FXML
    private TextField emailField;
    @FXML
    private PasswordField passwordField;
    @FXML
    private Button loginButton;
    @FXML
    private Button backButton;
    @FXML
    private Hyperlink forgotPasswordLink;

    private MainViewController mainController;

    private LoginController loginController;

    @FXML
    public void initialize() {
        // Configurar eventos para los botones y enlaces
        setupButtons();

        loginController = new LoginController();
    }

    private void setupButtons() {
        loginButton.setOnAction(event -> handleLogin());
        backButton.setOnAction(event -> closeWindow());
        forgotPasswordLink.setOnAction(event -> showForgotPasswordDialog());
    }

    private void handleLogin() {
        String email = emailField.getText().trim();
        String password = passwordField.getText().trim();

        // Validación básica (en una app real se haría contra una base de datos)
        if (email.isEmpty() || password.isEmpty()) {
            showAlert("Error", "Por favor, ingresa tu correo y contraseña.");
            return;
        }

        //TODO: Probar con el servidor
        LoginRequestDto loginRequestDto = new LoginRequestDto(email, password);
        //loginController.loguearse(loginRequestDto);
        System.out.println("He logueado: " + email);

        //TODO: Extracción de datos del response, para continuar ???
        String username = email.split("@")[0];

        mainController.onLoginSuccess(username, email);

        closeWindow();
    }

    private void showForgotPasswordDialog() {
        Alert alert = new Alert(Alert.AlertType.INFORMATION);
        alert.setTitle("Recuperar contraseña");
        alert.setHeaderText(null);
        alert.setContentText("Funcionalidad no implementada todavía.");
        alert.showAndWait();
    }

    private void showAlert(String title, String message) {
        Alert alert = new Alert(Alert.AlertType.ERROR);
        alert.setTitle(title);
        alert.setHeaderText(null);
        alert.setContentText(message);
        alert.showAndWait();
    }

    private void closeWindow() {
        Stage stage = (Stage) loginButton.getScene().getWindow();
        stage.close();
    }

    public void setMainController(MainViewController controller) {
        this.mainController = controller;
    }
}
