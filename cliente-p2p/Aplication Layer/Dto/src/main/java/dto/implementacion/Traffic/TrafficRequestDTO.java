package dto.implementacion.Traffic;

public class TrafficRequestDTO {

    private String command;

    public TrafficRequestDTO() {
    }

    public TrafficRequestDTO(String command) {
        this.command = command;
    }

    public String getCommand() {
        return command;
    }

    public void setCommand(String command) {
        this.command = command;
    }

    @Override
    public String toString() {
        return "TrafficRequestDTO{" +
                "command='" + command + '\'' +
                '}';
    }
}
