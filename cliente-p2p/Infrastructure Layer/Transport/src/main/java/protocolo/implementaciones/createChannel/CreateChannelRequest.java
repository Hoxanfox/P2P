package protocolo.implementaciones.createChannel;

import dto.implementacion.CreateChannel.CreateChannelRequestDto;
import dto.implementacion.CreateChannel.Channel.InvitacionDto;
import dto.implementacion.CreateChannel.Channel.MiembroCanalDto;
import protocolo.interfaces.RequestRoute;

public class CreateChannelRequest implements RequestRoute {

    private final CreateChannelRequestDto data;

    public CreateChannelRequest(CreateChannelRequestDto data) {
        this.data = data;
    }

    @Override
    public String toJson() {
        StringBuilder sb = new StringBuilder();
        sb.append("{\"command\":\"create-channel\",\"data\":{");

        // Nombre y descripción
        sb.append("\"nombre\":\"").append(escape(data.getNombre())).append("\",");
        sb.append("\"descripcion\":\"").append(escape(data.getDescripcion())).append("\",");

        // Miembros
        sb.append("\"miembros\":[");
        for (int i = 0; i < data.getMiembros().size(); i++) {
            MiembroCanalDto m = data.getMiembros().get(i);
            sb.append(String.format("{\"id\":\"%s\",\"email\":\"%s\"}",
                    m.getId(), escape(m.getEmail())));
            if (i < data.getMiembros().size() - 1) sb.append(",");
        }
        sb.append("],");

        // Invitaciones
        sb.append("\"invitaciones\":[");
        for (int i = 0; i < data.getInvitaciones().size(); i++) {
            InvitacionDto inv = data.getInvitaciones().get(i);
            sb.append("{\"destinatario\":");
            sb.append(String.format("{\"id\":\"%s\",\"email\":\"%s\"},",
                    inv.getDestinatario().getId(),
                    escape(inv.getDestinatario().getEmail())));
            sb.append(String.format("\"fechaEnvio\":\"%s\",", escape(inv.getFechaEnvio())));
            sb.append("\"estado\":").append(
                    inv.getEstado() == null ? "null" : "\"" + escape(inv.getEstado()) + "\""
            );
            sb.append("}");
            if (i < data.getInvitaciones().size() - 1) sb.append(",");
        }
        sb.append("],");

        // Chat tipo
        sb.append("\"chat\":{");
        sb.append("\"tipo\":\"").append(escape(data.getTipo())).append("\"");
        sb.append("}");

        sb.append("}}");
        return sb.toString();
    }

    private String escape(String input) {
        return input == null ? "" : input.replace("\"", "\\\"");
    }
}
