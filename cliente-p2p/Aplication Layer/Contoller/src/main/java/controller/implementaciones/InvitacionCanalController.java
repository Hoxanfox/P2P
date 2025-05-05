package controller.implementaciones;

import dto.implementacion.InvitacionCanal.InvitacionCanalRequestDto;
import facade.implementaciones.InvitacionCanalFacade;
import transport.TransportContext;
import transport.TcpPersistentTransportStrategy;
import transport.interfaces.ITransportStrategy;

public class InvitacionCanalController {

    private final InvitacionCanalFacade invitacionCanalFacade = new InvitacionCanalFacade();
    private final TransportContext context;
    private final InvitacionCanalRequestDto requestDto;

    // ✅ Ahora recibe el contexto y el request en el constructor
    public InvitacionCanalController(TransportContext context, InvitacionCanalRequestDto requestDto) {
        this.context = context;
        this.requestDto = requestDto;
    }

    public void enviarInvitacionAUsuario() {
        System.out.println("[InvitacionCanalController] === INICIO ENVÍO DE INVITACIÓN ===");

        try {
            invitacionCanalFacade.ejecutarFlujo(requestDto, context);
            System.out.println("[InvitacionCanalController] === INVITACIÓN ENVIADA CORRECTAMENTE ===");
        } catch (Exception e) {
            System.err.println("[InvitacionCanalController] Error al enviar la invitación: " + e.getMessage());
        } finally {
            // ✅ Cierra la estrategia si corresponde
            ITransportStrategy strategy = context.getStrategy();
            if (strategy instanceof TcpPersistentTransportStrategy tcpStrategy) {
                tcpStrategy.close();
                System.out.println("[InvitacionCanalController] Conexión cerrada.");
            }
        }

        System.out.println("[InvitacionCanalController] === FIN ENVÍO DE INVITACIÓN ===");
    }
}
