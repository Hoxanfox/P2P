package appmain.ui;

import javafx.application.Application;
import javafx.fxml.FXMLLoader;
import javafx.scene.Scene;
import javafx.stage.Stage;

import java.io.IOException;

import transport.TcpPersistentTransportStrategy;
import transport.TransportContext;

public class UniChatApplication extends Application {

    @Override
    public void start(Stage stage) throws IOException {
        // Cargar la vista del Splash Screen
        FXMLLoader splashLoader = new FXMLLoader(getClass().getResource("/appmain/ui/SplashScreen/splash-screen.fxml"));
        Scene splashScene = new Scene(splashLoader.load(), 800, 600);

        // Mostrar la escena del Splash
        stage.setTitle("Splash Screen!");
        stage.setScene(splashScene);
        stage.show();
    }

    public static void main(String[] args) {
        TcpPersistentTransportStrategy estrategia = new TcpPersistentTransportStrategy("localhost", 9000); // Reemplaza con IP y puerto reales
        TransportContext context = new TransportContext(estrategia);
        launch();
    }
}
