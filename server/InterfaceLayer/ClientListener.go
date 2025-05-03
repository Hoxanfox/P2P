package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Message struct {
	Command string          `json:"command"`
	Data    json.RawMessage `json:"data"`
}

// Estructuras de datos
type User struct {
	ID     string
	Email  string
	Nombre string
}

type File struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

type Member struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
}

type Chat struct {
	ID         int      `json:"id"`
	Tipo       string   `json:"tipo"`
	Miembros   []Member `json:"miembros"`
	TipoChatID string      `json:"tipoChatId"`
}

type MessageData struct {
	Mensaje struct {
		Remitente struct {
			ID     string `json:"id"`
			Nombre string `json:"correo"` // se llama 'correo' en el JSON de entrada
		} `json:"remitente"`
		Destinatario struct {
			ID     string `json:"id"`
			Nombre string `json:"correo"` // se llama 'correo' en el JSON de entrada
		} `json:"destinatario"`
		Contenido  string `json:"contenido"`
		FechaEnvio string `json:"fechaEnvio"`
		Archivo    *File  `json:"archivo"`
	} `json:"mensaje"`
}

type MessageResponse struct {
	ID           string     `json:"id"`
	Remitente    Member  `json:"remitente"`
	Destinatario Member  `json:"destinatario"`
	Contenido    string  `json:"contenido"`
	FechaEnvio   string  `json:"fechaEnvio"`
	Archivo      *File   `json:"archivo"`
	Chat         Chat    `json:"chat"`
}

type GenericResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Usuarios disponibles
var users = []User{
	{
		ID:     "210c3aea-d243-4b6c-8456-7bb67ff5306e",
		Email:  "ana.torres@correo.com",
		Nombre: "Ana Torres",
	},
	{
		ID:     "0db20b34-00a6-48d0-8ebb-49de460a99a4",
		Email:  "luis.mendoza@correo.com",
		Nombre: "Luis Mendoza",
	},
}

// Envío de respuesta
func sendResponse(conn net.Conn, response GenericResponse) {
	respBytes, _ := json.Marshal(response)
	conn.Write(append(respBytes, '\n'))
}

// Ruta: send-message-user
func handleSendMessageUser(conn net.Conn, msg Message) {
	var request MessageData
	if err := json.Unmarshal(msg.Data, &request); err != nil {
		sendResponse(conn, GenericResponse{"error", "Datos inválidos del mensaje", nil})
		return
	}

	remitenteEmail := request.Mensaje.Remitente.Nombre
	destinatarioEmail := request.Mensaje.Destinatario.Nombre

	var remitente *User
	var destinatario *User

	for i := range users {
		if users[i].Email == remitenteEmail {
			remitente = &users[i]
		}
		if users[i].Email == destinatarioEmail {
			destinatario = &users[i]
		}
	}

	if remitente == nil || destinatario == nil {
		sendResponse(conn, GenericResponse{"error", "No se pudo enviar el mensaje. Verifique los datos del usuario o del chat.", nil})
		return
	}

	response := MessageResponse{
		ID: "210c3aea-d243-4b6c-8456-7bb67ff5306e",
		Remitente: Member{
			ID:     remitente.ID,
			Nombre: remitente.Nombre,
		},
		Destinatario: Member{
			ID:     destinatario.ID,
			Nombre: destinatario.Nombre,
		},
		Contenido:  request.Mensaje.Contenido,
		FechaEnvio: request.Mensaje.FechaEnvio,
		Archivo:    request.Mensaje.Archivo,
		Chat: Chat{
			ID:   7,
			Tipo: "privado",
			Miembros: []Member{
				{ID: remitente.ID, Nombre: remitente.Nombre},
				{ID: destinatario.ID, Nombre: destinatario.Nombre},
			},
			TipoChatID: "210c3aea-d243-4b6c-8456-7bb67ff5306e",
		},
	}

	sendResponse(conn, GenericResponse{
		Status:  "success",
		Message: "Mensaje enviado correctamente",
		Data:    response,
	})
}

// Manejo de comandos
func handleConnection(conn net.Conn) {
	defer conn.Close()

	var buffer = make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("[ERROR] No se pudo leer la conexión:", err)
		return
	}

	var msg Message
	if err := json.Unmarshal(buffer[:n], &msg); err != nil {
		sendResponse(conn, GenericResponse{"error", "Formato de mensaje inválido", nil})
		return
	}

	if msg.Command == "send-message-user" {
		handleSendMessageUser(conn, msg)
	} else {
		sendResponse(conn, GenericResponse{"error", "Comando no reconocido", nil})
	}
}

// Servidor principal
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
