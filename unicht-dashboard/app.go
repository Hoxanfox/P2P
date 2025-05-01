package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

// App struct
type App struct {
	ctx context.Context
	logs []string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}


func (a *App) GetActiveUsers() int {
	//TODO: Add logic
	return 42
}

func (a *App) GetGroupCount() int {
	//TODO: Add logic
	return 7
}

func (a *App) GetServerIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Desconocida"
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			ip := ipnet.IP.To4()
			if ip != nil {
				if ip[0] == 192 || ip[0] == 10 {
					return ip.String()
				}
			}
		}
	}
	return "No encontrada"
}

func (a *App) GetOtherServers() []string {
	//Example
	return []string{"192.168.1.2", "192.168.1.3"}
}
func (a *App) GenerateTestLogs() {
	testLogs := []string{
		"Usuario admin ha iniciado sesión",
		"Nuevo grupo 'Desarrollo' creado",
		"Servidor backup iniciado en 192.168.1.3",
		"Actualización de sistema completada",
		"Error de conexión con servidor remoto",
	}
	
	for _, log := range testLogs {
		a.AddLog(log)
		time.Sleep(time.Second) // Simula diferentes timestamps
	}
}
func (a *App) AddLog(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	log := "[" + timestamp + "] " + message
	a.logs = append(a.logs, log)

	// Limita la cantidad de logs guardados (opcional)
	if len(a.logs) > 500 {
		a.logs = a.logs[len(a.logs)-500:]
	}
}

func (a *App) GetLogs() string {
	return stringJoin(a.logs, "\n")
}

func stringJoin(s []string, sep string) string {
	result := ""
	for i, str := range s {
		if i > 0 {
			result += sep
		}
		result += str
	}
	return result
}
