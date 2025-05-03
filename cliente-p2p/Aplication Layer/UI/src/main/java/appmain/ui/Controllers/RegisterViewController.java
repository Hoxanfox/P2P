package appmain.ui.Controllers;

import javafx.fxml.FXML;
import javafx.scene.control.*;
import javafx.stage.Stage;
import appmain.ui.MainViewController;

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
    private Hyperlink loginLink;

    private MainViewController mainController;

    @FXML
    public void initialize() {
        // Configurar eventos
        setupButtons();
    }

    private void setupButtons() {
        registerButton.setOnAction(event -> handleRegister());
        backButton.setOnAction(event -> closeWindow());
        loginLink.setOnAction(event -> openLoginForm());
    }

    private void handleRegister() {
        String username = usernameField.getText().trim();
        String email = emailField.getText().trim();
        String password = passwordField.getText();
        String confirmPassword = confirmPasswordField.getText();

        if (username.isEmpty() || email.isEmpty() || password.isEmpty() || confirmPassword.isEmpty()) {
            showAlert("Error", "Todos los campos son obligatorios.");
            return;
        }

        if (!password.equals(confirmPassword)) {
            showAlert("Error", "Las contraseñas no coinciden.");
            return;
        }

        // Simulación de registro exitoso
        mainController.onLoginSuccess(username, email, username); // Usamos username como fullName temporalmente
        closeWindow();
    }

    private void openLoginForm() {
        closeWindow();
        mainController.openLoginForm();
    }

    private void showAlert(String title, String message) {
        Alert alert = new Alert(Alert.AlertType.ERROR);
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
