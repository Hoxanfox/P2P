package main

import (
	"fmt"
	"interfacesLayer"
	"transport"
)

func main() {
	// Crear una estrategia TCP apuntando al servidor
	tcpStrategy := transport.NewTcpTransportStrategy("localhost", 9000)

	// Crear un contexto de transporte usando la estrategia TCP
	context := transport.NewTransportContext(tcpStrategy)

	// JSON de prueba
	jsonToSend := `{"command":"say_hello","data":"Hola desde Golang"}`

	// Ejecutar el envío
	response := context.ExecuteSend(jsonToSend)

	// Mostrar la respuesta
	fmt.Println("Respuesta del servidor:", response)
}
