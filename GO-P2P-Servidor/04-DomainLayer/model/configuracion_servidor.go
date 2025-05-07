package model

import "errors"

// Errores de validación para ConfiguracionServidor
var (
    ErrMaxConexionesNoPositivo = errors.New("maxConexiones debe ser mayor que cero")
    ErrParametrosMySQLVacio    = errors.New("parametrosMySQL no puede estar vacío")
    ErrRutasLogsVacio          = errors.New("rutasLogs no puede estar vacío")
)

// ConfiguracionServidor agrupa los parámetros de configuración del servidor.
type ConfiguracionServidor struct {
    maxConexiones   int
    parametrosMySQL string
    rutasLogs       string
}

// NewConfiguracionServidor construye una ConfiguracionServidor válida.
func NewConfiguracionServidor(
    maxConexiones int,
    parametrosMySQL, rutasLogs string,
) (*ConfiguracionServidor, error) {
    if maxConexiones <= 0 {
        return nil, ErrMaxConexionesNoPositivo
    }
    if parametrosMySQL == "" {
        return nil, ErrParametrosMySQLVacio
    }
    if rutasLogs == "" {
        return nil, ErrRutasLogsVacio
    }
    return &ConfiguracionServidor{
        maxConexiones:   maxConexiones,
        parametrosMySQL: parametrosMySQL,
        rutasLogs:       rutasLogs,
    }, nil
}

// Getters
func (c *ConfiguracionServidor) MaxConexiones() int    { return c.maxConexiones }
func (c *ConfiguracionServidor) ParametrosMySQL() string { return c.parametrosMySQL }
func (c *ConfiguracionServidor) RutasLogs() string       { return c.rutasLogs }
