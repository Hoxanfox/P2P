module p2p.ui {
    requires javafx.controls;
    requires javafx.fxml;
    requires javafx.web;

    requires org.controlsfx.controls;
    requires com.dlsc.formsfx;
    requires net.synedra.validatorfx;
    requires org.kordamp.ikonli.javafx;
    requires org.kordamp.bootstrapfx.core;
    requires com.almasb.fxgl.all;

    opens p2p.ui to javafx.fxml;
    opens p2p.ui.Controllers to javafx.fxml;

    exports p2p.ui;
    exports p2p.ui.Controllers;

}