package pool

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// DBConfig contiene toda la configuración para la conexión a la base de datos
type DBConfig struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
		SSLMode  string `yaml:"sslmode"`
		CertPath string `yaml:"cert_path"`
	} `yaml:"database"`
	Pool struct {
		MaxOpenConnections    int    `yaml:"max_open_connections"`
		MaxIdleConnections    int    `yaml:"max_idle_connections"`
		MaxConnectionLifetime string `yaml:"max_connection_lifetime"`
		MaxConnectionIdleTime string `yaml:"max_connection_idle_time"`
	} `yaml:"pool"`
}

// TLSConfig contiene la configuración para conexiones seguras
type TLSConfig struct {
	CertPath           string
	InsecureSkipVerify bool
}

// DBConnectionPool es un wrapper sobre sql.DB que proporciona métricas
// y gestión centralizada de conexiones a la base de datos MySQL.
type DBConnectionPool struct {
	db             *sql.DB
	config         *DBConfig
	log            *logrus.Logger
	queryCount     int64
	errorCount     int64
	latencySum     int64
	latencyCount   int64
	lastError      error
	lastErrorMutex sync.RWMutex
}

// LoadDBConfig carga la configuración desde un archivo YAML
func LoadDBConfig(configPath string) (*DBConfig, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error leyendo archivo de configuración: %w", err)
	}

	var config DBConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error decodificando configuración YAML: %w", err)
	}

	return &config, nil
}

// NewDBConnectionPool crea un nuevo pool de conexiones a MySQL desde un archivo de configuración
func NewDBConnectionPool(configPath string) (*DBConnectionPool, error) {
	config, err := LoadDBConfig(configPath)
	if err != nil {
		return nil, fmt.Errorf("error cargando configuración: %w", err)
	}

	// Crear logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Registrar TLS si está habilitado
	if config.Database.SSLMode == "REQUIRED" {
		// Resolver ruta del certificado (relativa al archivo de configuración si no es absoluta)
		certPath := config.Database.CertPath
		if !filepath.IsAbs(certPath) {
			certPath = filepath.Join(filepath.Dir(configPath), certPath)
		}

		tlsConfig := TLSConfig{
			CertPath:           certPath,
			InsecureSkipVerify: false,
		}

		if err := registerTLS(tlsConfig); err != nil {
			return nil, fmt.Errorf("error configurando TLS: %w", err)
		}
	}

	// Construir DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
	)

	// Añadir TLS si está habilitado
	if config.Database.SSLMode == "REQUIRED" {
		dsn += "&tls=custom"
	}

	// Abrir la conexión
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error abriendo conexión a MySQL: %w", err)
	}

	// Configurar el pool
	db.SetMaxOpenConns(config.Pool.MaxOpenConnections)
	db.SetMaxIdleConns(config.Pool.MaxIdleConnections)

	// Configurar tiempos máximos
	connLifetime, err := time.ParseDuration(config.Pool.MaxConnectionLifetime)
	if err != nil {
		logger.Warnf("Error en configuración de MaxConnectionLifetime '%s': %v. Usando valor por defecto (1h).",
			config.Pool.MaxConnectionLifetime, err)
		connLifetime = time.Hour
	}
	db.SetConnMaxLifetime(connLifetime)

	idleTime, err := time.ParseDuration(config.Pool.MaxConnectionIdleTime)
	if err != nil {
		logger.Warnf("Error en configuración de MaxConnectionIdleTime '%s': %v. Usando valor por defecto (15m).",
			config.Pool.MaxConnectionIdleTime, err)
		idleTime = 15 * time.Minute
	}
	db.SetConnMaxIdleTime(idleTime)

	// Verificar conexión
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("error verificando conexión a MySQL: %w", err)
	}

	pool := &DBConnectionPool{
		db:     db,
		config: config,
		log:    logger,
	}

	logger.Info("Pool de conexiones a MySQL inicializado correctamente")
	return pool, nil
}

// registerTLS registra una configuración TLS personalizada para MySQL
func registerTLS(config TLSConfig) error {
	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(config.CertPath)
	if err != nil {
		return fmt.Errorf("error leyendo certificado: %w", err)
	}

	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		return fmt.Errorf("error agregando certificado PEM al pool")
	}

	err = mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs:            rootCertPool,
		InsecureSkipVerify: config.InsecureSkipVerify,
	})

	if err != nil {
		return fmt.Errorf("error registrando configuración TLS: %w", err)
	}

	return nil
}

// DB devuelve el objeto *sql.DB subyacente
func (p *DBConnectionPool) DB() *sql.DB {
	return p.db
}

// Stats devuelve las estadísticas nativas del pool de conexiones
func (p *DBConnectionPool) Stats() sql.DBStats {
	return p.db.Stats()
}

// Close cierra el pool de conexiones ordenadamente
func (p *DBConnectionPool) Close() error {
	p.log.Info("Cerrando pool de conexiones a MySQL")
	return p.db.Close()
}

// ExecContext ejecuta una consulta SQL que no devuelve filas, con seguimiento de métricas
func (p *DBConnectionPool) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	// Incrementar contador de consultas
	atomic.AddInt64(&p.queryCount, 1)

	// Medir tiempo de ejecución
	start := time.Now()
	result, err := p.db.ExecContext(ctx, query, args...)
	elapsed := time.Since(start)

	// Actualizar métricas de latencia
	atomic.AddInt64(&p.latencySum, int64(elapsed))
	atomic.AddInt64(&p.latencyCount, 1)

	// Registrar error si ocurre
	if err != nil {
		atomic.AddInt64(&p.errorCount, 1)
		p.lastErrorMutex.Lock()
		p.lastError = err
		p.lastErrorMutex.Unlock()
		p.log.WithFields(logrus.Fields{
			"query":   query,
			"args":    args,
			"elapsed": elapsed,
		}).Error("Error ejecutando consulta SQL:", err)
	}

	return result, err
}

// QueryContext ejecuta una consulta SQL que devuelve filas, con seguimiento de métricas
func (p *DBConnectionPool) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// Incrementar contador de consultas
	atomic.AddInt64(&p.queryCount, 1)

	// Medir tiempo de ejecución
	start := time.Now()
	rows, err := p.db.QueryContext(ctx, query, args...)
	elapsed := time.Since(start)

	// Actualizar métricas de latencia
	atomic.AddInt64(&p.latencySum, int64(elapsed))
	atomic.AddInt64(&p.latencyCount, 1)

	// Registrar error si ocurre
	if err != nil {
		atomic.AddInt64(&p.errorCount, 1)
		p.lastErrorMutex.Lock()
		p.lastError = err
		p.lastErrorMutex.Unlock()
		p.log.WithFields(logrus.Fields{
			"query":   query,
			"args":    args,
			"elapsed": elapsed,
		}).Error("Error ejecutando consulta SQL:", err)
	}

	return rows, err
}

// QueryRowContext ejecuta una consulta SQL que devuelve una sola fila
func (p *DBConnectionPool) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	// Incrementar contador de consultas
	atomic.AddInt64(&p.queryCount, 1)

	// Medir tiempo de ejecución (aunque no podemos saber si hay error hasta Scan())
	start := time.Now()
	row := p.db.QueryRowContext(ctx, query, args...)
	elapsed := time.Since(start)

	// Actualizar métricas de latencia
	atomic.AddInt64(&p.latencySum, int64(elapsed))
	atomic.AddInt64(&p.latencyCount, 1)

	return row
}

// BeginTx inicia una transacción con opciones específicas
func (p *DBConnectionPool) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	// Medir tiempo de ejecución
	start := time.Now()
	tx, err := p.db.BeginTx(ctx, opts)
	elapsed := time.Since(start)

	// Actualizar métricas de latencia
	atomic.AddInt64(&p.latencySum, int64(elapsed))
	atomic.AddInt64(&p.latencyCount, 1)

	// Registrar error si ocurre
	if err != nil {
		atomic.AddInt64(&p.errorCount, 1)
		p.lastErrorMutex.Lock()
		p.lastError = err
		p.lastErrorMutex.Unlock()
		p.log.WithFields(logrus.Fields{
			"elapsed": elapsed,
		}).Error("Error iniciando transacción:", err)
	}

	return tx, err
}

// PingContext verifica que la conexión a la base de datos sigue siendo válida
func (p *DBConnectionPool) PingContext(ctx context.Context) error {
	// Medir tiempo de ejecución
	start := time.Now()
	err := p.db.PingContext(ctx)
	elapsed := time.Since(start)

	// Actualizar métricas de latencia
	atomic.AddInt64(&p.latencySum, int64(elapsed))
	atomic.AddInt64(&p.latencyCount, 1)

	// Registrar error si ocurre
	if err != nil {
		atomic.AddInt64(&p.errorCount, 1)
		p.lastErrorMutex.Lock()
		p.lastError = err
		p.lastErrorMutex.Unlock()
		p.log.WithFields(logrus.Fields{
			"elapsed": elapsed,
		}).Error("Error haciendo ping a la base de datos:", err)
	}

	return err
}

// GetMetrics devuelve métricas del pool para monitorización
func (p *DBConnectionPool) GetMetrics() map[string]interface{} {
	// Obtener stats de sql.DB
	stats := p.db.Stats()

	// Calcular latencia promedio
	var avgLatency time.Duration
	latencySum := atomic.LoadInt64(&p.latencySum)
	latencyCount := atomic.LoadInt64(&p.latencyCount)
	if latencyCount > 0 {
		avgLatency = time.Duration(latencySum / latencyCount)
	}

	// Crear mapa de métricas
	metrics := map[string]interface{}{
		"max_open_conns":     stats.MaxOpenConnections,
		"open_conns":         stats.OpenConnections,
		"in_use":             stats.InUse,
		"idle":               stats.Idle,
		"wait_count":         stats.WaitCount,
		"wait_duration":      stats.WaitDuration,
		"max_idle_closed":    stats.MaxIdleClosed,
		"max_lifetime_closed": stats.MaxLifetimeClosed,
		"query_count":        atomic.LoadInt64(&p.queryCount),
		"error_count":        atomic.LoadInt64(&p.errorCount),
		"avg_latency_ns":     avgLatency,
	}

	// Añadir último error si existe
	p.lastErrorMutex.RLock()
	if p.lastError != nil {
		metrics["last_error"] = p.lastError.Error()
	}
	p.lastErrorMutex.RUnlock()

	return metrics
}
