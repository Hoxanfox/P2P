package appmain.ui;

import appmain.ui.Controllers.GroupViewController;
import javafx.fxml.FXML;
import javafx.fxml.FXMLLoader;
import javafx.scene.Scene;
import javafx.scene.control.*;
import javafx.scene.layout.*;
import javafx.scene.shape.Circle;
import javafx.stage.Modality;
import javafx.stage.Stage;
import javafx.stage.StageStyle;
import appmain.ui.Controllers.LoginViewController;
import appmain.ui.Controllers.ProfileViewController;
import appmain.ui.Controllers.RegisterViewController;

import java.io.IOException;
import java.util.Objects;

import controller.implementaciones.RegisterController;

public class MainViewController {

    @FXML private BorderPane mainContainer;
    @FXML private VBox sidebarContainer;
    @FXML private TabPane contactsTabPane;
    @FXML private VBox contactsContainer;
    @FXML private VBox groupsContainer;
    @FXML private StackPane contentArea;
    @FXML private VBox welcomeScreen;
    @FXML private BorderPane chatScreen;
    @FXML private VBox loginScreen;
    @FXML private Button loginButton;
    @FXML private Button registerButton;
    @FXML private Button settingsButton;
    @FXML private Button profileButton;
    @FXML private Label usernameLabel;
    @FXML private Circle profilePic;
    @FXML private TextField searchField;
    @FXML private Button addGroupButton;
    @FXML private TextField messageField;
    @FXML private Button sendButton;
    @FXML private Button attachButton;

    private RegisterViewController registerViewController;

    // Variables para los datos del perfil y estado de autenticación
    private boolean isLogin = false; // Cambia a true cuando el usuario está logueado
    private String username = "";
    private String email = "";
    private String fullName = "";
    private String status = "Disponible";

    @FXML
    public void initialize() {
        registerViewController = new RegisterViewController();

        // Configurar la interfaz según el estado de autenticación
        updateAuthenticationState();

        // Configurar listeners para los botones de autenticación
        setupAuthButtons();

        // Configurar listener para el botón de perfil
        setupProfileButton();

        // Cargar contactos y grupos de ejemplo
        loadSampleContacts();
        loadSampleGroups();

        // Configurar el campo de mensaje y el botón de enviar
        setupChatControls();
    }


    private void updateAuthenticationState() {
        if (!isLogin) {
            // Mostrar pantalla de login y ocultar otras vistas
            loginScreen.setVisible(true);
            welcomeScreen.setVisible(false);
            chatScreen.setVisible(false);
            sidebarContainer.setVisible(false);

            // Ocultar controles de la barra superior
            profileButton.setVisible(false);
            settingsButton.setVisible(false);
        } else {
            // Ocultar pantalla de login y mostrar interfaz principal
            loginScreen.setVisible(false);
            welcomeScreen.setVisible(true);
            sidebarContainer.setVisible(true);

            // Mostrar controles de la barra superior
            profileButton.setVisible(true);
            settingsButton.setVisible(true);

            // Actualizar información del usuario
            usernameLabel.setText(username.isEmpty() ? "Usuario" : username);
        }
    }

    private void setupAuthButtons() {
        loginButton.setOnAction(event -> openLoginForm());
        registerButton.setOnAction(event -> openRegisterForm());
        addGroupButton.setOnAction(event -> openNewGroupForm());
        settingsButton.setOnAction(event -> openSettingsView());
        attachButton.setOnAction(event -> openSubmitFile() );
    }

    private void setupProfileButton() {
        profileButton.setOnAction(event -> openProfileForm());
    }

    private void setupChatControls() {
        sendButton.setOnAction(event -> {
            String message = messageField.getText().trim();
            if (!message.isEmpty()) {
                sendMessage(message);
                messageField.clear();
            }
        });

        messageField.setOnAction(event -> {
            String message = messageField.getText().trim();
            if (!message.isEmpty()) {
                sendMessage(message);
                messageField.clear();
            }
        });
    }

    public void openLoginForm() {
        try {
            FXMLLoader loader = new FXMLLoader(getClass().getResource("/appmain/ui/Registration/login-view.fxml"));
            Scene loginScene = new Scene(loader.load(), 400, 500);

            LoginViewController controller = loader.getController();
            controller.setMainController(this);

            Stage loginStage = new Stage();
            loginStage.initStyle(StageStyle.UNDECORATED);
            loginStage.initModality(Modality.APPLICATION_MODAL);
            loginStage.setScene(loginScene);
            loginStage.showAndWait();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public void openRegisterForm() {
        try {
            FXMLLoader loader = new FXMLLoader(getClass().getResource("/appmain/ui/Registration/register-view.fxml"));
            Scene registerScene = new Scene(loader.load(), 450, 650);

            RegisterViewController controller = loader.getController();
            controller.setMainController(this);

            Stage registerStage = new Stage();
            registerStage.initStyle(StageStyle.UNDECORATED);
            registerStage.initModality(Modality.APPLICATION_MODAL);
            registerStage.setScene(registerScene);
            registerStage.showAndWait();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private void openProfileForm() {
        try {
            FXMLLoader loader = new FXMLLoader(getClass().getResource("/appmain/ui/Profile/profile-view.fxml"));
            Scene profileScene = new Scene(loader.load(), 500, 700);

            ProfileViewController controller = loader.getController();
            controller.setUserData(username, email, fullName, status);
            controller.setMainController(this);

            Stage profileStage = new Stage();
            profileStage.initStyle(StageStyle.UNDECORATED);
            profileStage.initModality(Modality.APPLICATION_MODAL);
            profileStage.setScene(profileScene);
            profileStage.showAndWait();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private void openNewGroupForm() {
        try {
            FXMLLoader loader = new FXMLLoader(getClass().getResource("/appmain/ui/Registration/group-view.fxml"));
            Scene groupScene = new Scene(loader.load(), 400, 500);

            GroupViewController controller = loader.getController();
            controller.setMainController(this);

            Stage groupStage= new Stage();
            groupStage.initStyle(StageStyle.UNDECORATED);
            groupStage.initModality(Modality.APPLICATION_MODAL);
            groupStage.setScene(groupScene);
            groupStage.showAndWait();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private void openSettingsView() {
        Alert alert = new Alert(Alert.AlertType.INFORMATION);
        alert.setTitle("Configuración del usuario");
        alert.setHeaderText(null);
        alert.setContentText("Funcionalidad no implementada todavía.");
        alert.showAndWait();
    }

    private void openSubmitFile() {
        Alert alert = new Alert(Alert.AlertType.INFORMATION);
        alert.setTitle("Adjuntar documento");
        alert.setHeaderText(null);
        alert.setContentText("Funcionalidad no implementada todavía.");
        alert.showAndWait();
    }

    // Método llamado desde LoginViewController cuando el login es exitoso
    public void onLoginSuccess(String username, String email) {
        this.username = username;
        this.email = email;
        this.isLogin = true;
        updateAuthenticationState();
    }

    // Método llamado desde ProfileViewController cuando se actualizan los datos del perfil
    public void updateProfileInfo(String username, String email, String fullName, String status) {
        this.username = username;
        this.email = email;
        this.fullName = fullName;
        this.status = status;
        usernameLabel.setText(username);
    }

    //TODO: Implemeentar logica para traer los contactos reales
    private void loadSampleContacts() {
        // Este método cargará contactos de ejemplo en el sidebar
        // En una implementación real, estos vendrían de una base de datos o servicio
        String[] sampleNames = {"Ana García", "Carlos López", "María Rodríguez", "Juan Pérez", "Laura Torres"};

        for (String name : sampleNames) {
            HBox contactItem = createContactItem(name, "En línea", false);
            contactsContainer.getChildren().add(contactItem);
        }
    }

    //TODO: Implementar logica para traer los ejemplso reales
    private void loadSampleGroups() {
        // Este método cargará grupos de ejemplo en el sidebar
        // En una implementación real, estos vendrían de una base de datos o servicio
        String[] sampleGroups = {"Proyecto Final", "Estudio Matemáticas", "Club de Debate", "Taller de Programación"};

        for (String name : sampleGroups) {
            HBox groupItem = createContactItem(name, "3 miembros", true);
            groupsContainer.getChildren().add(groupItem);
        }
    }

    private HBox createContactItem(String name, String status, boolean isGroup) {
        HBox itemContainer = new HBox();
        itemContainer.getStyleClass().add("contact-item");

        Circle profileCircle = new Circle(20);
        profileCircle.getStyleClass().add("contact-pic");

        VBox infoContainer = new VBox();
        infoContainer.getStyleClass().add("contact-info");

        Label nameLabel = new Label(name);
        nameLabel.getStyleClass().add("contact-name");

        Label statusLabel = new Label(status);
        statusLabel.getStyleClass().add("contact-status");

        infoContainer.getChildren().addAll(nameLabel, statusLabel);
        itemContainer.getChildren().addAll(profileCircle, infoContainer);

        // Agregar evento de clic para abrir el chat
        itemContainer.setOnMouseClicked(event -> openChat(name, status, isGroup));

        return itemContainer;
    }

    private void openChat(String name, String status, boolean isGroup) {
        // Cambiar a la vista de chat y actualizar la información
        welcomeScreen.setVisible(false);
        chatScreen.setVisible(true);

        // Actualizar la información del encabezado del chat
        Label chatNameLabel = (Label) chatScreen.lookup("#chatNameLabel");
        Label chatStatusLabel = (Label) chatScreen.lookup("#chatStatusLabel");

        if (chatNameLabel != null) chatNameLabel.setText(name);
        if (chatStatusLabel != null) chatStatusLabel.setText(isGroup ? status : "En línea");
    }

    private void sendMessage(String message) {
        // Este método añadiría el mensaje a la conversación actual
        // En una implementación real, enviaría el mensaje a través de un servicio
        VBox messagesContainer = (VBox) chatScreen.lookup("#messagesContainer");

        HBox messageBox = new HBox();
        messageBox.getStyleClass().addAll("message-bubble", "sent-message");

        Label messageText = new Label(message);
        messageText.setWrapText(true);

        messageBox.getChildren().add(messageText);
        messagesContainer.getChildren().add(messageBox);
    }
}