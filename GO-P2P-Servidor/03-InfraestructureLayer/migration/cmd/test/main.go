package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"migration"
	"pool"
)

func main() {
	// Configuración
	configFile := "../../../pool/db_config.yaml"
	logger := setupLogger()

	// Paso 1: Conectar a la base de datos
	logger.Info("Conectando a la base de datos...")
	dbPool, err := pool.NewDBConnectionPool(configFile)
	if err != nil {
		logger.WithError(err).Fatal("Error conectando a la base de datos")
	}
	defer dbPool.Close()

	// Crear un contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Paso 2: Crear migrador y verificar estado inicial
	migrator := migration.NewMigrator(dbPool).
		WithLogger(logger).
		WithStrictMode(true)

	// Cargar las migraciones directamente desde el sistema de archivos
	logger.Info("Cargando migraciones desde archivos...")
	migrationsDir := "../../migrations"
	if err := migrator.LoadMigrationsFromDir(migrationsDir); err != nil {
		logger.WithError(err).Fatal("Error cargando migraciones")
	}

	// Mostrar estado inicial
	logger.Info("Estado inicial de la base de datos:")
	displayStatus(migrator)

	// Paso 3: Verificar si existe alguna tabla de las migraciones
	// y borrarlas si es necesario para la prueba
	if tablesExist(ctx, dbPool) {
		logger.Warn("Se encontraron tablas existentes, limpiando para la prueba...")
		if err := migrator.DropSchemaForTesting(ctx); err != nil {
			logger.WithError(err).Fatal("Error eliminando esquema existente")
		}
	}

	// Paso 4: Ejecutar migraciones
	logger.Info("Ejecutando migraciones...")
	start := time.Now()
	if err := migrator.Run(ctx); err != nil {
		logger.WithError(err).Fatal("Error ejecutando migraciones")
	}
	elapsed := time.Since(start)
	logger.WithField("tiempo", elapsed.String()).Info("Migraciones completadas")

	// Paso 5: Verificar que las tablas y datos se crearon correctamente
	logger.Info("Verificando tablas y datos creados...")
	verifyTablesAndData(ctx, dbPool, logger)
	
	// Paso 6: Mostrar estado final de migraciones
	logger.Info("Estado final de migraciones:")
	displayStatus(migrator)

	// Paso 7: Limpiar - eliminar las tablas creadas
	logger.Info("Limpiando las tablas creadas...")
	if err := migrator.DropSchemaForTesting(ctx); err != nil {
		logger.WithError(err).Fatal("Error eliminando tablas")
	}
	
	logger.Info("Prueba completada con éxito. Todas las tablas han sido eliminadas.")
}

// Configurar el logger
func setupLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return log
}

// Mostrar el estado de las migraciones
func displayStatus(migrator *migration.Migrator) {
	migrations := migrator.Status()
	
	fmt.Println("\nEstado de las migraciones:")
	fmt.Println("+----------+---------------+----------------------------------------+----------+")
	fmt.Println("| Versión  | Estado        | Descripción                            | Aplicada |")
	fmt.Println("+----------+---------------+----------------------------------------+----------+")
	
	for _, m := range migrations {
		estado := "PENDIENTE"
		aplicada := "-"
		if m.Applied {
			estado = "APLICADA"
			aplicada = m.AppliedAt.Format("2006-01-02 15:04:05")
		}
		
		desc := m.Description
		if len(desc) > 40 {
			desc = desc[:37] + "..."
		}
		
		fmt.Printf("| %-8d | %-13s | %-38s | %-8s |\n", 
			m.Version, estado, desc, aplicada)
	}
	
	fmt.Println("+----------+---------------+----------------------------------------+----------+")
	
	currentVersion := migrator.GetCurrentVersion()
	fmt.Printf("Versión actual de la base de datos: %d\n\n", currentVersion)
}

// Verificar si existen tablas de las migraciones
func tablesExist(ctx context.Context, dbPool *pool.DBConnectionPool) bool {
	var tableName string
	// Intentar seleccionar una tabla que sabemos que crearía la migración
	err := dbPool.DB().QueryRowContext(ctx, 
		"SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = 'usuario_servidor' LIMIT 1").
		Scan(&tableName)
	
	return err == nil && tableName != ""
}

// Verificar que las tablas y datos se crearon correctamente
func verifyTablesAndData(ctx context.Context, dbPool *pool.DBConnectionPool, logger *logrus.Logger) {
	// 1. Verificar las tablas principales
	tables := []string{
		"usuario_servidor", 
		"canal_servidor", 
		"canal_miembro", 
		"mensaje_servidor",
		"schema_migrations",
	}
	
	for _, table := range tables {
		if !checkTableExists(ctx, dbPool, table) {
			logger.Fatalf("Error: La tabla %s no fue creada", table)
		}
		logger.Infof("✓ Tabla %s creada correctamente", table)
	}
	
	// 2. Verificar datos de usuario admin
	var count int
	err := dbPool.DB().QueryRowContext(ctx, 
		"SELECT COUNT(*) FROM usuario_servidor WHERE nombre_usuario = 'admin'").Scan(&count)
	
	if err != nil {
		logger.WithError(err).Fatal("Error verificando usuario admin")
	}
	
	if count == 0 {
		logger.Fatal("Error: No se encontró el usuario admin")
	}
	logger.Info("✓ Usuario admin creado correctamente")
	
	// 3. Verificar canales predeterminados
	err = dbPool.DB().QueryRowContext(ctx, 
		"SELECT COUNT(*) FROM canal_servidor WHERE nombre = 'General'").Scan(&count)
	
	if err != nil {
		logger.WithError(err).Fatal("Error verificando canal general")
	}
	
	if count == 0 {
		logger.Fatal("Error: No se encontró el canal general")
	}
	logger.Info("✓ Canal general creado correctamente")
}

// Verificar si una tabla específica existe
func checkTableExists(ctx context.Context, dbPool *pool.DBConnectionPool, tableName string) bool {
	var name string
	err := dbPool.DB().QueryRowContext(ctx, 
		"SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = ?", 
		tableName).Scan(&name)
		
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Printf("Error verificando tabla %s: %v", tableName, err)
		return false
	}
	
	return name != ""
}
