package appmain.ui.Controllers;

import javafx.fxml.FXML;
import javafx.scene.control.*;
import javafx.scene.layout.*;
import javafx.scene.shape.Circle;
import javafx.stage.Stage;
import appmain.ui.MainViewController;

public class ProfileViewController {

    // Header
    @FXML private Button closeButton;

    // Personal Information Fields
    @FXML private TextField usernameField;
    @FXML private TextField emailField;

    // Account Settings Buttons
    @FXML private Button changePasswordButton;

    // Footer Buttons
    @FXML private Button cancelButton;
    @FXML private Button changePicButton;
    @FXML private Button saveButton;

    @FXML
    private void initialize() {
        setupButtons();
    }

    private void setupButtons() {
        closeButton.setOnAction(event -> handleCloseButton());
        cancelButton.setOnAction(event -> handleCloseButton());
        changePasswordButton.setOnAction(event -> handleChangePasswordButton());
        changePicButton.setOnAction(event -> handleChangePicButton());
        saveButton.setOnAction(event -> handleSaveButton());
    }

    private void handleCloseButton() {
        // Obtiene el StackPane raíz
        Stage stage = (Stage) closeButton.getScene().getWindow();
        stage.close();
        System.out.println("Vista de perfil cerrada.");
    }

    private void handleChangePasswordButton(){
        Alert alert = new Alert(Alert.AlertType.INFORMATION);
        alert.setTitle("ChangePassword");
        alert.setHeaderText(null);
        alert.setContentText("Funcionalidad no disponible");

        alert.showAndWait();
    }

    private void handleSaveButton(){
        Alert alert = new Alert(Alert.AlertType.INFORMATION);
        alert.setTitle("Guardar información");
        alert.setHeaderText(null);
        alert.setContentText("Funcionalidad no disponible");

        alert.showAndWait();
    }


    private void handleChangePicButton() {

        Alert alert = new Alert(Alert.AlertType.INFORMATION);
        alert.setTitle("Cambiar foto de perfil");
        alert.setHeaderText(null);
        alert.setContentText("Funcionalidad no disponible");

        alert.showAndWait();
    }

    public void setUserData(String username, String email, String fullName, String status) {
    }

    public void setMainController(MainViewController mainViewController) {
    }
}
