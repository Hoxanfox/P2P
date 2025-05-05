package views.autenticacion.mainView;

import controller.implementaciones.ListUsersController;
import dto.implementacion.ListUsers.ListUsersResponseDto;
import dto.implementacion.ListUsers.UsuarioResponseDTO;
import transport.TransportContext;

import javax.swing.*;
import java.awt.*;
import java.util.List;

public class DashboardView extends JFrame {

    private final ListUsersController listUsersController;
    private final JPanel userListPanel;
    private final JLabel statusLabel;
    private final JButton refreshButton;

    public DashboardView(TransportContext transportContext) {
        setTitle("Dashboard de Usuarios");
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setSize(600, 400);
        setLocationRelativeTo(null);

        this.listUsersController = new ListUsersController(transportContext);

        // Panel principal con un layout más dinámico
        userListPanel = new JPanel();
        userListPanel.setLayout(new BoxLayout(userListPanel, BoxLayout.Y_AXIS));
        userListPanel.setBackground(Color.WHITE); // Fondo blanco

        // Etiqueta de estado con estilo
        statusLabel = new JLabel("Cargando usuarios...");
        statusLabel.setFont(new Font("Arial", Font.BOLD, 16));
        statusLabel.setAlignmentX(Component.CENTER_ALIGNMENT);
        userListPanel.add(statusLabel);

        // Botón de refrescar con borde redondeado y color
        refreshButton = new JButton("Refrescar usuarios");
        refreshButton.setFont(new Font("Arial", Font.PLAIN, 14));
        refreshButton.setAlignmentX(Component.CENTER_ALIGNMENT);
        refreshButton.setBackground(new Color(34, 150, 242)); // Azul
        refreshButton.setForeground(Color.WHITE); // Texto blanco
        refreshButton.setFocusPainted(false);
        refreshButton.setBorder(BorderFactory.createLineBorder(new Color(34, 150, 242), 2));
        refreshButton.setCursor(new Cursor(Cursor.HAND_CURSOR));
        refreshButton.addActionListener(e -> cargarUsuarios());

        userListPanel.add(Box.createRigidArea(new Dimension(0, 10)));
        userListPanel.add(refreshButton);

        JScrollPane scrollPane = new JScrollPane(userListPanel);
        add(scrollPane, BorderLayout.CENTER);

        // Cargar usuarios al iniciar
        cargarUsuarios();

        setVisible(true);
    }

    private void cargarUsuarios() {
        SwingWorker<List<UsuarioResponseDTO>, Void> worker = new SwingWorker<>() {
            @Override
            protected List<UsuarioResponseDTO> doInBackground() {
                System.out.println("Obteniendo usuarios desde el backend...");
                ListUsersResponseDto response = listUsersController.obtenerUsuarios();
                return (response != null) ? response.getUsuarios() : null;
            }

            @Override
            protected void done() {
                try {
                    List<UsuarioResponseDTO> usuarios = get();
                    renderizarUsuarios(usuarios);
                } catch (Exception e) {
                    e.printStackTrace();
                    mostrarMensaje("Error al obtener los usuarios.");
                }
            }
        };

        worker.execute();
    }

    private void renderizarUsuarios(List<UsuarioResponseDTO> usuarios) {
        userListPanel.removeAll(); // Limpiar contenido anterior

        if (usuarios == null || usuarios.isEmpty()) {
            mostrarMensaje("No se encontraron usuarios.");
            return;
        }

        for (UsuarioResponseDTO user : usuarios) {
            JPanel userPanel = new JPanel();
            userPanel.setLayout(new BoxLayout(userPanel, BoxLayout.Y_AXIS));
            userPanel.setBackground(new Color(245, 245, 245)); // Fondo gris suave
            userPanel.setBorder(BorderFactory.createEmptyBorder(10, 10, 10, 10));
            userPanel.setMaximumSize(new Dimension(Integer.MAX_VALUE, 100));

            String info = String.format(
                    "<html><b>Nombre:</b> %s<br><b>Email:</b> %s<br><b>Conectado:</b> %s</html>",
                    user.getNombre(), user.getEmail(), user.isIs_connected() ? "Sí" : "No"
            );

            JLabel label = new JLabel(info);
            label.setFont(new Font("Arial", Font.PLAIN, 14));
            label.setForeground(Color.BLACK);

            // Estilo adicional para los usuarios
            userPanel.add(label);

            // Agregar un borde redondeado a cada panel de usuario
            userPanel.setBorder(BorderFactory.createLineBorder(new Color(200, 200, 200), 1, true));

            userListPanel.add(userPanel);
        }

        // Agregar botón al final nuevamente
        userListPanel.add(Box.createRigidArea(new Dimension(0, 10)));
        userListPanel.add(refreshButton);

        userListPanel.revalidate();
        userListPanel.repaint();
    }

    private void mostrarMensaje(String mensaje) {
        userListPanel.removeAll();
        statusLabel.setText(mensaje);
        userListPanel.add(statusLabel);
        userListPanel.add(Box.createRigidArea(new Dimension(0, 10)));
        userListPanel.add(refreshButton);
        userListPanel.revalidate();
        userListPanel.repaint();
    }

    // Método que será llamado por el TrafficController para actualizar la vista
    public void actualizarVistaConUsuarios() {
        cargarUsuarios();  // Actualiza la vista con los nuevos datos de los usuarios
    }
}
