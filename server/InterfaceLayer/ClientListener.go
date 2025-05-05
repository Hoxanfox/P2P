package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
)

// Estructura general del mensaje recibido
type Message struct {
	Command string          `json:"command"`
	Data    json.RawMessage `json:"data"`
}

// Estructura del usuario
type User struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Nombre      string `json:"nombre"`
	Password    string `json:"password"`
	Foto        string `json:"foto"`
	IP          string `json:"ip"`
	CreatedAt   string `json:"created_at"`
	IsConnected bool   `json:"is_connected"`
}

// Solicitudes
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nombre   string `json:"nombre"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Estructura de respuesta
type GenericResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Usuarios simulados (iniciales)
var users = []User{
	{
		ID:          uuid.New().String(),
		Email:       "juan@example.com",
		Nombre:      "juan123",
		Password:    "1234",
		Foto:        "No disponible",
		IP:          "127.0.0.1",
		CreatedAt:   "2025-04-28T10:30:00",
		IsConnected: true,
	},
	{
		ID:          uuid.New().String(),
		Email:       "ana@example.com",
		Nombre:      "ana456",
		Password:    "5678",
		Foto:        "No disponible",
		IP:          "127.0.0.2",
		CreatedAt:   "2025-04-28T11:00:00",
		IsConnected: false,
	},
}

// Enviar respuestas al cliente
func sendResponse(conn net.Conn, response GenericResponse) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println("[ERROR] Marshal response:", err)
		return
	}
	conn.Write(append(respBytes, '\n'))
	fmt.Printf("[DEBUG] Enviado respuesta: %v\n", response) // Debug: respuesta enviada
}

// Manejador de login
func handleLogin(conn net.Conn, msg Message) {
	var request LoginRequest
	if err := json.Unmarshal(msg.Data, &request); err != nil {
		fmt.Println("[DEBUG] Error al deserializar mensaje de login:", err)
		sendResponse(conn, GenericResponse{"error", "Datos inválidos de login", nil})
		return
	}

	var user *User
	for i := range users {
		if users[i].Email == request.Email && users[i].Password == request.Password {
			user = &users[i]
			break
		}
	}

	if user == nil {
		fmt.Println("[DEBUG] Usuario no encontrado o credenciales incorrectas")
		sendResponse(conn, GenericResponse{"error", "Email o contraseña incorrectos", nil})
		return
	}

	fmt.Printf("[DEBUG] Usuario encontrado: %v\n", user) // Debug: usuario encontrado
	sendResponse(conn, GenericResponse{
		Status:  "success",
		Message: "Inicio de sesión exitoso",
		Data: map[string]interface{}{
			"id":           user.ID,
			"nombre":       user.Nombre,
			"email":        user.Email,
			"photo":        user.Foto,
			"ip":           user.IP,
			"created_at":   user.CreatedAt,
			"is_connected": user.IsConnected,
		},
	})
}

// Manejador de registro
func handleRegister(conn net.Conn, msg Message) {
	var request RegisterRequest
	if err := json.Unmarshal(msg.Data, &request); err != nil {
		fmt.Println("[DEBUG] Error al deserializar mensaje de registro:", err)
		sendResponse(conn, GenericResponse{"error", "Datos inválidos de registro", nil})
		return
	}

	// Verifica email duplicado
	for _, user := range users {
		if user.Email == request.Email {
			fmt.Println("[DEBUG] Email duplicado detectado:", request.Email)
			sendResponse(conn, GenericResponse{"error", "El email ya está registrado", nil})
			return
		}
	}

	// Crear nuevo usuario
	newUser := User{
		ID:          uuid.New().String(),
		Email:       request.Email,
		Nombre:      request.Nombre,
		Password:    request.Password,
		Foto:        "No disponible",
		IP:          "127.0.0.1", // podrías obtenerla desde conn.RemoteAddr()
		CreatedAt:   time.Now().Format(time.RFC3339),
		IsConnected: false,
	}
	users = append(users, newUser)

	fmt.Printf("[DEBUG] Usuario registrado: %v\n", newUser) // Debug: nuevo usuario registrado

	sendResponse(conn, GenericResponse{
		Status:  "success",
		Message: "Registro exitoso",
		Data: map[string]interface{}{
			"id":           newUser.ID,
			"nombre":       newUser.Nombre,
			"email":        newUser.Email,
			"photo":        newUser.Foto,
			"ip":           newUser.IP,
			"created_at":   newUser.CreatedAt,
			"is_connected": newUser.IsConnected,
		},
	})
}

func handleListUsers(conn net.Conn) {
    // Si no hay usuarios, enviar error
	if len(users) == 0 {
		sendResponse(conn, GenericResponse{
			Status:  "error",
			Message: "No se pudieron obtener los usuarios registrados",
			Data:    nil,
		})
		return
	}

    // Si hay usuarios, construir la lista de usuarios
	var userList []map[string]interface{}
	for _, user := range users {
		userList = append(userList, map[string]interface{}{
			"id":           user.ID,
			"nombre":       user.Nombre,
			"email":        user.Email,
			"is_connected": user.IsConnected,
		})
	}

    // Enviar la respuesta con la lista de usuarios
	sendResponse(conn, GenericResponse{
		Status:  "success",
		Message: "Usuarios registrados obtenidos correctamente",
		Data:    userList,
	})
}


// Manejador de conexión
func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("[DEBUG] Nueva conexión aceptada desde:", conn.RemoteAddr())

	// Loop para seguir esperando comandos hasta que la conexión se cierre
	for {
		buffer := make([]byte, 4096)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("[ERROR] No se pudo leer la conexión:", err)
			return
		}

		var msg Message
		if err := json.Unmarshal(buffer[:n], &msg); err != nil {
			fmt.Println("[DEBUG] Error al deserializar mensaje:", err)
			sendResponse(conn, GenericResponse{"error", "Formato de mensaje inválido", nil})
			return
		}

		fmt.Printf("[DEBUG] Mensaje recibido: %v\n", msg) // Debug: mensaje recibido

		switch msg.Command {
		case "login":
			handleLogin(conn, msg)
		case "register":
			handleRegister(conn, msg)
		case "list-users":
			handleListUsers(conn)
		default:
			fmt.Println("[DEBUG] Comando no reconocido:", msg.Command)
			sendResponse(conn, GenericResponse{"error", "Comando no reconocido", nil})
		}
	}
}

// Función principal del servidor
func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("[INFO] Servidor TCP escuchando en el puerto 9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[ERROR] Error al aceptar conexión:", err)
			continue
		}
		go handleConnection(conn)
	}
}
