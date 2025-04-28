package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io"
    "net"
    "strings"
)

// Definimos el mensaje general que se intercambia
type Message struct {
    Command string          `json:"command"`
    Data    json.RawMessage `json:"data"`
}

// Estructura de datos que el cliente manda al registrarse
type RegisterData struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Photo    string `json:"photo"`
    IP       string `json:"ip"`
}

// Estructura de respuesta al registrar
type RegisterResponse struct {
    Status   string `json:"status"`
    Message  string `json:"message"`
    UserID   int    `json:"userId"`
    Username string `json:"username"`
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Println("Nueva conexión:", conn.RemoteAddr())

    reader := bufio.NewReader(conn)

    for {
        // Leer datos hasta encontrar un salto de línea
        data, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                fmt.Println("Cliente cerró la conexión")
            } else {
                fmt.Println("Error leyendo:", err)
            }
            break
        }

        // Limpiar espacios en blanco
        data = strings.TrimSpace(data)
        fmt.Println("Recibido:", data)

        var msg Message
        err = json.Unmarshal([]byte(data), &msg)
        if err != nil {
            fmt.Println("Error deserializando JSON:", err)
            continue
        }

        // Ejecutar lógica basada en el comando recibido
        switch msg.Command {
        case "say_hello":
            // El cliente envía un saludo
            var saludo string
            err := json.Unmarshal(msg.Data, &saludo)
            if err != nil {
                fmt.Println("Error leyendo saludo:", err)
                continue
            }
            fmt.Printf("Comando recibido: %s, Datos: %s\n", msg.Command, saludo)

            responseMessage := "¡Hola! Recibí tu mensaje: " + saludo
            responseData, _ := json.Marshal(responseMessage)

            response := Message{
                Command: "response",
                Data:    responseData,
            }
            respBytes, _ := json.Marshal(response)
            conn.Write(append(respBytes, '\n'))

        case "register":
            // El cliente envía datos de registro
            var registerData RegisterData
            err := json.Unmarshal(msg.Data, &registerData)
            if err != nil {
                fmt.Println("Error leyendo datos de registro:", err)
                continue
            }
            fmt.Printf("Comando recibido: %s, Datos: %+v\n", msg.Command, registerData)

            // Aquí podrías guardar a la base de datos y generar un ID real
            fakeUserID := 123 // Simulación

            responseContent := RegisterResponse{
                Status:   "success",
                Message:  "Usuario registrado exitosamente",
                UserID:   fakeUserID,
                Username: registerData.Username,
            }

            responseData, _ := json.Marshal(responseContent)

            response := Message{
                Command: "register_response",
                Data:    responseData,
            }
            respBytes, _ := json.Marshal(response)
            conn.Write(append(respBytes, '\n'))

        default:
            fmt.Println("Comando no reconocido:", msg.Command)

            errorMessage := "Comando no reconocido: " + msg.Command
            responseData, _ := json.Marshal(errorMessage)

            response := Message{
                Command: "error",
                Data:    responseData,
            }
            respBytes, _ := json.Marshal(response)
            conn.Write(append(respBytes, '\n'))
        }
    }
}

func main() {
    // Escuchar en el puerto 9000
    listener, err := net.Listen("tcp", ":9000")
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    fmt.Println("Servidor TCP escuchando en el puerto 9000...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error aceptando conexión:", err)
            continue
        }
        go handleConnection(conn) // Manejar cada conexión concurrentemente
    }
}
