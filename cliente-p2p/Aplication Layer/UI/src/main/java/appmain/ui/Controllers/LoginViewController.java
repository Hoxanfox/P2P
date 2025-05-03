package p2p.ui.Controllers;

import javafx.fxml.FXML;
import javafx.scene.control.*;
import javafx.stage.Stage;
import p2p.ui.MainViewController;

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
    private Hyperlink registerLink;
    @FXML
    private Hyperlink forgotPasswordLink;

    private MainViewController mainController;

    @FXML
    public void initialize() {
        // Configurar eventos para los botones y enlaces
        setupButtons();
    }

    private void setupButtons() {
        loginButton.setOnAction(event -> handleLogin());
        backButton.setOnAction(event -> closeWindow());
        registerLink.setOnAction(event -> openRegisterForm());
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

        // Simulación de autenticación exitosa
        String username = email.split("@")[0]; // Por ejemplo, "usuario" de "usuario@gmail.com"
        String fullName = "Nombre de Ejemplo";

        mainController.onLoginSuccess(username, email, fullName);

        closeWindow();
    }

    private void openRegisterForm() {
        closeWindow();
        mainController.openRegisterForm();
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
