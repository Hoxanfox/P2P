package dto.implementacion.Traffic;

import com.fasterxml.jackson.annotation.JsonProperty;

public class TrafficResponseDto {

    private String command;
    private Object data;

    public TrafficResponseDto() {
    }

    public TrafficResponseDto(String command, Object data) {
        this.command = command;
        this.data = data;
    }

    public String getCommand() {
        return command;
    }

    public void setCommand(String command) {
        this.command = command;
    }

    public Object getData() {
        return data;
    }

    public void setData(Object data) {
        this.data = data;
    }

    @Override
    public String toString() {
        return "TrafficResponseDTO{" +
                "command='" + command + '\'' +
                ", data=" + data +
                '}';
    }
}
