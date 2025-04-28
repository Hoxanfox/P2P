package p2p.ui;

import javafx.fxml.FXML;
import javafx.scene.control.Label;
import Interfaces.UIControllerInterface;

public class HelloController implements UIControllerInterface{
    @FXML
    private Label welcomeText;

    @FXML
    @Override
    public void onClickButton() {
        welcomeText.setText("Welcome to JavaFX Application!");
    }
}