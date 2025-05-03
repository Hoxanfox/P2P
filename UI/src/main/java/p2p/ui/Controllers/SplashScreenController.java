package p2p.ui.Controllers;

import javafx.animation.*;
import javafx.fxml.FXML;
import javafx.fxml.FXMLLoader;
import javafx.scene.Scene;
import javafx.scene.control.ProgressBar;
import javafx.scene.layout.BorderPane;
import javafx.scene.layout.StackPane;
import javafx.scene.layout.VBox;
import javafx.scene.shape.Circle;
import javafx.scene.text.Text;
import javafx.stage.Stage;
import javafx.util.Duration;

import java.io.IOException;
import java.util.Objects;

public class SplashScreenController {

    @FXML
    private StackPane splashScreenStackPane;

    @FXML
    private VBox splashBox;

    @FXML
    private Text splashText;

    @FXML
    private Text taglineText;

    @FXML
    private ProgressBar loadingBar;

    @FXML
    private Circle circle1;

    @FXML
    private Circle circle2;

    @FXML
    private Circle circle3;

    public void initialize() {
        // Configurar opacidad inicial
        splashScreenStackPane.setOpacity(0.0);

        // Crear animación para aparecer el fondo
        FadeTransition fadeInBackground = new FadeTransition(Duration.seconds(0.8), splashScreenStackPane);
        fadeInBackground.setFromValue(0.0);
        fadeInBackground.setToValue(1.0);

        // Crear animación para el cuadro principal
        splashBox.setScaleX(0.7);
        splashBox.setScaleY(0.7);

        ScaleTransition scaleBox = new ScaleTransition(Duration.seconds(0.6), splashBox);
        scaleBox.setFromX(0.7);
        scaleBox.setFromY(0.7);
        scaleBox.setToX(1.0);
        scaleBox.setToY(1.0);
        scaleBox.setInterpolator(Interpolator.EASE_OUT);

        // Crear animación para el título
        splashText.setOpacity(0);
        FadeTransition fadeInTitle = new FadeTransition(Duration.seconds(0.8), splashText);
        fadeInTitle.setFromValue(0.0);
        fadeInTitle.setToValue(1.0);
        fadeInTitle.setDelay(Duration.seconds(0.3));

        // Crear animación para el tagline
        taglineText.setOpacity(0);
        FadeTransition fadeInTagline = new FadeTransition(Duration.seconds(0.8), taglineText);
        fadeInTagline.setFromValue(0.0);
        fadeInTagline.setToValue(1.0);
        fadeInTagline.setDelay(Duration.seconds(0.6));

        // Animar los círculos
        animateCircles();

        // Animar la barra de progreso
        animateProgressBar();

        // Secuencia para cargar toda la aplicación
        PauseTransition loadingDelay = new PauseTransition(Duration.seconds(3.5));
        loadingDelay.setOnFinished(e -> prepareExitAnimation());

        // Ejecutar todas las animaciones iniciales
        fadeInBackground.play();
        scaleBox.play();
        fadeInTitle.play();
        fadeInTagline.play();
        loadingDelay.play();
    }

    private void animateCircles() {
        // Inicializar círculos
        circle1.setOpacity(0);
        circle2.setOpacity(0);
        circle3.setOpacity(0);

        // Crear animaciones para los círculos
        FadeTransition fadeCircle1 = new FadeTransition(Duration.seconds(0.4), circle1);
        fadeCircle1.setFromValue(0);
        fadeCircle1.setToValue(1);
        fadeCircle1.setDelay(Duration.seconds(0.2));

        FadeTransition fadeCircle2 = new FadeTransition(Duration.seconds(0.4), circle2);
        fadeCircle2.setFromValue(0);
        fadeCircle2.setToValue(1);
        fadeCircle2.setDelay(Duration.seconds(0.4));

        FadeTransition fadeCircle3 = new FadeTransition(Duration.seconds(0.4), circle3);
        fadeCircle3.setFromValue(0);
        fadeCircle3.setToValue(1);
        fadeCircle3.setDelay(Duration.seconds(0.6));

        // Crear animaciones cíclicas de pulso para los círculos
        Timeline pulseAnimation = new Timeline(
                new KeyFrame(Duration.ZERO,
                        new KeyValue(circle1.radiusProperty(), 8),
                        new KeyValue(circle2.radiusProperty(), 8),
                        new KeyValue(circle3.radiusProperty(), 8)
                ),
                new KeyFrame(Duration.seconds(0.5),
                        new KeyValue(circle1.radiusProperty(), 10),
                        new KeyValue(circle2.radiusProperty(), 8),
                        new KeyValue(circle3.radiusProperty(), 10)
                ),
                new KeyFrame(Duration.seconds(1.0),
                        new KeyValue(circle1.radiusProperty(), 8),
                        new KeyValue(circle2.radiusProperty(), 10),
                        new KeyValue(circle3.radiusProperty(), 8)
                ),
                new KeyFrame(Duration.seconds(1.5),
                        new KeyValue(circle1.radiusProperty(), 10),
                        new KeyValue(circle2.radiusProperty(), 8),
                        new KeyValue(circle3.radiusProperty(), 10)
                ),
                new KeyFrame(Duration.seconds(2.0),
                        new KeyValue(circle1.radiusProperty(), 8),
                        new KeyValue(circle2.radiusProperty(), 8),
                        new KeyValue(circle3.radiusProperty(), 8)
                )
        );
        pulseAnimation.setCycleCount(Timeline.INDEFINITE);

        // Ejecutar las animaciones de los círculos
        fadeCircle1.play();
        fadeCircle2.play();
        fadeCircle3.play();
        pulseAnimation.play();
    }

    private void animateProgressBar() {
        // Inicializar barra de progreso
        loadingBar.setProgress(0);

        // Animar la barra con un TimeLine para hacer el efecto más natural
        Timeline progressTimeline = new Timeline();

        KeyValue kv1 = new KeyValue(loadingBar.progressProperty(), 0.3, Interpolator.EASE_BOTH);
        KeyFrame kf1 = new KeyFrame(Duration.seconds(0.8), kv1);

        KeyValue kv2 = new KeyValue(loadingBar.progressProperty(), 0.6, Interpolator.EASE_BOTH);
        KeyFrame kf2 = new KeyFrame(Duration.seconds(1.6), kv2);

        KeyValue kv3 = new KeyValue(loadingBar.progressProperty(), 0.8, Interpolator.EASE_BOTH);
        KeyFrame kf3 = new KeyFrame(Duration.seconds(2.4), kv3);

        KeyValue kv4 = new KeyValue(loadingBar.progressProperty(), 1.0, Interpolator.EASE_OUT);
        KeyFrame kf4 = new KeyFrame(Duration.seconds(3.0), kv4);

        progressTimeline.getKeyFrames().addAll(kf1, kf2, kf3, kf4);
        progressTimeline.play();
    }

    private void prepareExitAnimation() {
        // Crear animación para desaparecer
        FadeTransition fadeOut = new FadeTransition(Duration.seconds(0.8), splashScreenStackPane);
        fadeOut.setFromValue(1.0);
        fadeOut.setToValue(0.0);

        // Ejecutar la transición a la pantalla principal después del fadeOut
        fadeOut.setOnFinished(event -> loadMainScreen());

        // Iniciar fadeOut
        fadeOut.play();
    }

    private void loadMainScreen() {
        try {
            BorderPane mainRoot = FXMLLoader.load(Objects.requireNonNull(getClass().getResource("/p2p/ui/UniChat.fxml")));
            Scene mainScene = new Scene(mainRoot, 800, 600);

            Stage splashStage = (Stage) splashScreenStackPane.getScene().getWindow();
            splashStage.setTitle("UniChat");
            splashStage.setScene(mainScene);
            splashStage.centerOnScreen();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}