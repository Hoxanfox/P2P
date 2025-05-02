package transport

import (
    "fmt"
    "io"
    "net"
    "bufio"
    "strings"
)

// TCPTransportStrategy implementa la interfaz ITransportStrategy
type TCPTransportStrategy struct {
    Host string
    Port int
}

// Asegura que TCPTransportStrategy implementa ITransportStrategy
var _ ITransportStrategy = (*TCPTransportStrategy)(nil)

// NewTCPTransportStrategy crea una nueva instancia de TCPTransportStrategy
func NewTCPTransportStrategy(host string, port int) *TCPTransportStrategy {
    return &TCPTransportStrategy{
        Host: host,
        Port: port,
    }
}

// SendJson envía un JSON a través de TCP y recibe la respuesta
func (t *TCPTransportStrategy) SendJson(jsonToSend string) string {
    fmt.Println("[DEBUG] Intentando conectar a", t.Host, ":", t.Port)

    // Conectar al servidor
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.Host, t.Port))
    if err != nil {
        fmt.Println("[ERROR] Error durante la conexión:", err)
        return ""
    }
    defer conn.Close()

    fmt.Println("[DEBUG] Conexión establecida.")

    // Enviar el JSON
    fmt.Println("[DEBUG] Enviando JSON:", jsonToSend)
    _, err = fmt.Fprintln(conn, jsonToSend)
    if err != nil {
        fmt.Println("[ERROR] Error enviando el JSON:", err)
        return ""
    }

    // Leer la respuesta
    reader := bufio.NewReader(conn)
    response, err := reader.ReadString('\n')
    if err != nil && err != io.EOF {
        fmt.Println("[ERROR] Error leyendo la respuesta:", err)
        return ""
    }

    // Limpiar el salto de línea
    response = strings.TrimSpace(response)

    fmt.Println("[DEBUG] Respuesta recibida:", response)

    return response
}
