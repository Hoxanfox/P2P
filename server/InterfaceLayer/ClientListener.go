package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io"
    "net"
    "strings"

    "github.com/google/uuid"
)

// Estructura general del mensaje entrante
type Message struct {
    Command string          `json:"command"`
    Data    json.RawMessage `json:"data"`
}

// Datos esperados para login
type LoginData struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Respuesta genérica del servidor
type GenericResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

// Estructura de usuario con UUID
type User struct {
    ID          string `json:"id"` // UUID
    Nombre      string `json:"nombre"`
    Email       string `json:"email"`
    IsConnected bool   `json:"is_connected"`
}

// Lista simulada de usuarios
var users = []User{
    {
        ID:          uuid.New().String(),
        Nombre:      "juan123",
        Email:       "juan@example.com",
        IsConnected: true,
    },
    {
        ID:          uuid.New().String(),
        Nombre:      "ana456",
        Email:       "ana@example.com",
        IsConnected: false,
    },
    // Puedes agregar más usuarios aquí
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Println("Nueva conexión desde:", conn.RemoteAddr())

    reader := bufio.NewReader(conn)

    for {
        raw, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                fmt.Println("Cliente desconectado:", conn.RemoteAddr())
            } else {
                fmt.Println("Error leyendo del cliente:", err)
            }
            return
        }

        raw = strings.TrimSpace(raw)
        if raw == "" {
            continue
        }

        fmt.Println("Recibido:", raw)

        var msg Message
        err = json.Unmarshal([]byte(raw), &msg)
        if err != nil {
            fmt.Println("Error al parsear JSON:", err)
            continue
        }

        switch msg.Command {
        case "login":
            handleLogin(conn, msg)
        case "list-users":
            handleListUsers(conn)
        default:
            sendResponse(conn, GenericResponse{
                Status:  "error",
                Message: "Comando no soportado",
            })
        }
    }
}

func handleLogin(conn net.Conn, msg Message) {
    var data LoginData
    err := json.Unmarshal(msg.Data, &data)
    if err != nil {
        fmt.Println("Error leyendo datos de login:", err)
        sendResponse(conn, GenericResponse{
            Status:  "error",
            Message: "Error procesando los datos de login",
        })
        return
    }

    if data.Email == "juan.perez@ejeasdasdplo.cm" && data.Password == "123adsasd456" {
        user := map[string]interface{}{
            "id":           uuid.New().String(),
            "nombre":       "juan123",
            "email":        "juan@mail.com",
            "photo":        "base64_encoded_image_string",
            "ip":           "192.168.1.15",
            "created_at":   "2025-04-28T10:30:00",
            "is_connected": true,
        }

        sendResponse(conn, GenericResponse{
            Status:  "success",
            Message: "Inicio de sesión exitoso",
            Data:    user,
        })
    } else {
        sendResponse(conn, GenericResponse{
            Status:  "error",
            Message: "Email o contraseña incorrectos",
        })
    }
}

func handleListUsers(conn net.Conn) {
    if len(users) == 0 {
        sendResponse(conn, GenericResponse{
            Status:  "error",
            Message: "No se pudieron obtener los usuarios registrados",
        })
        return
    }

    sendResponse(conn, GenericResponse{
        Status:  "success",
        Message: "Usuarios registrados obtenidos correctamente",
        Data:    users,
    })
}

func sendResponse(conn net.Conn, content interface{}) {
    bytes, err := json.Marshal(content)
    if err != nil {
        fmt.Println("Error serializando respuesta:", err)
        return
    }

    fmt.Println("Enviando respuesta:", string(bytes))

    if _, err := conn.Write(append(bytes, '\n')); err != nil {
        fmt.Println("Error enviando respuesta al cliente:", err)
    }
}

func main() {
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

        go handleConnection(conn)
    }
}
