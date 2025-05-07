package model

import "testing"

func TestNewConfiguracionServidor_Success(t *testing.T) {
    cfg, err := NewConfiguracionServidor(
        100,
        "user=root password=xyz dbname=servidor sslmode=require",
        "/var/logs/servidor",
    )
    if err != nil {
        t.Fatalf("esperaba sin error, obtuvo %v", err)
    }
    if cfg.MaxConexiones() != 100 {
        t.Errorf("MaxConexiones: esperado %d, obtuvo %d", 100, cfg.MaxConexiones())
    }
    if cfg.ParametrosMySQL() != "user=root password=xyz dbname=servidor sslmode=require" {
        t.Errorf("ParametrosMySQL: esperado %q, obtuvo %q",
            "user=root password=xyz dbname=servidor sslmode=require", cfg.ParametrosMySQL())
    }
    if cfg.RutasLogs() != "/var/logs/servidor" {
        t.Errorf("RutasLogs: esperado %q, obtuvo %q", "/var/logs/servidor", cfg.RutasLogs())
    }
}

func TestNewConfiguracionServidor_Errors(t *testing.T) {
    cases := []struct {
        name             string
        maxConexiones    int
        parametrosMySQL  string
        rutasLogs        string
        wantErr          error
    }{
        {"MaxConexiones cero", 0, "p", "r", ErrMaxConexionesNoPositivo},
        {"MaxConexiones negativa", -5, "p", "r", ErrMaxConexionesNoPositivo},
        {"ParametrosMySQL vacío", 10, "", "r", ErrParametrosMySQLVacio},
        {"RutasLogs vacío", 10, "p", "", ErrRutasLogsVacio},
    }
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            _, err := NewConfiguracionServidor(c.maxConexiones, c.parametrosMySQL, c.rutasLogs)
            if err != c.wantErr {
                t.Errorf("%s: esperado error %v, obtuvo %v", c.name, c.wantErr, err)
            }
        })
    }
}
