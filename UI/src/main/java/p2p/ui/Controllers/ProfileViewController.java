package p2p.ui.Controllers;

import javafx.fxml.FXML;
import javafx.scene.control.*;
import javafx.scene.layout.*;
import javafx.scene.shape.Circle;
import javafx.stage.Stage;
import p2p.ui.MainViewController;

public class ProfileViewController {

    // Header
    @FXML private Button closeButton;

    // Profile Picture Section
    @FXML private Circle profileImage;
    @FXML private Button changePicButton;
    @FXML private Label fullNameLabel;
    @FXML private Label userStatusLabel;

    // Personal Information Fields
    @FXML private TextField usernameField;
    @FXML private TextField firstNameField;
    @FXML private TextField lastNameField;
    @FXML private TextField emailField;
    @FXML private ComboBox<String> statusComboBox;
    @FXML private TextArea bioTextArea;

    // Account Settings Buttons
    @FXML private Button changePasswordButton;
    @FXML private Button notificationsButton;
    @FXML private Button privacyButton;

    // Footer Buttons
    @FXML private Button cancelButton;
    @FXML private Button saveButton;

    @FXML
    private void initialize() {
        // Aquí puedes inicializar cosas como valores iniciales del ComboBox, listeners, etc.
        statusComboBox.getItems().addAll("Disponible", "Ocupado", "Ausente");
        setupButtons();
    }

    private void setupButtons() {
        closeButton.setOnAction(event -> handleCloseButton());
    }

    // Puedes agregar métodos para manejar los eventos de los botones, por ejemplo:
    @FXML
    private void handleSaveButton() {
        System.out.println("Guardar cambios:");
        System.out.println("Nombre de usuario: " + usernameField.getText());
        System.out.println("Nombre: " + firstNameField.getText());
        System.out.println("Apellido: " + lastNameField.getText());
        System.out.println("Email: " + emailField.getText());
        System.out.println("Estado: " + statusComboBox.getValue());
        System.out.println("Biografía: " + bioTextArea.getText());
    }

    private void handleCancelButton() {
        System.out.println("Cancelar cambios");
    }

    private void handleCloseButton() {
        // Obtiene el StackPane raíz
        Stage stage = (Stage) closeButton.getScene().getWindow();
        stage.close();


        System.out.println("Vista de perfil cerrada.");
    }


    private void handleChangePicButton() {
        System.out.println("Cambiar foto de perfil");
    }

    private void handleChangePasswordButton() {
        System.out.println("Cambiar contraseña");
    }

    private void handleNotificationsButton() {
        System.out.println("Configurar notificaciones");
    }

    private void handlePrivacyButton() {
        System.out.println("Privacidad y seguridad");
    }

    public void setUserData(String username, String email, String fullName, String status) {
    }

    public void setMainController(MainViewController mainViewController) {
    }
}
