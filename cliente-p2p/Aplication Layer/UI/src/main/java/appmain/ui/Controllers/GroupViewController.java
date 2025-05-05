package appmain.ui.Controllers;

import javafx.fxml.FXML;
import javafx.scene.control.*;
import javafx.stage.Stage;
import appmain.ui.MainViewController;


public class GroupViewController {

    @FXML
    private TextField nameField;
    @FXML
    private TextField descriptionfield;
    @FXML
    private Button createButton;
    @FXML
    private Button backButton;
    @FXML

    private MainViewController mainController;


    @FXML
    public void initialize() {
        // Configurar eventos para los botones y enlaces
        setupButtons();

    }

    private void setupButtons() {
        createButton.setOnAction(event -> handleCreateGroup());
        backButton.setOnAction(event -> closeWindow());
    }

    private void handleCreateGroup() {
        String name = nameField.getText();
        String description = descriptionfield.getText();

        //TODO: Implementar logica para crear grupo
        showAlert("Exitoso", "Tu grupo ha sido creado exitosamente", Alert.AlertType.INFORMATION);
        closeWindow();
    }

    private void showAlert(String title, String message, Alert.AlertType type) {
        Alert alert = new Alert(type);
        alert.setTitle(title);
        alert.setHeaderText(null);
        alert.setContentText(message);
        alert.showAndWait();
    }

    private void closeWindow() {
        Stage stage = (Stage) createButton.getScene().getWindow();
        stage.close();
    }

    public void setMainController(MainViewController controller) {
        this.mainController = controller;
    }
}
