package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"migration"
	"pool"
)

func main() {
	// Configurar flags de línea de comandos
	configFile := flag.String("config", "../../../pool/db_config.yaml", "Ruta al archivo de configuración de la base de datos")
	dryRun := flag.Bool("dry-run", false, "Solo mostrar las migraciones que se aplicarían sin ejecutarlas")
	strict := flag.Bool("strict", true, "Detener en caso de error")
	forceVersion := flag.Int("force-version", 0, "Forzar hasta una versión específica (0 = todas)")
	outOfOrder := flag.Bool("out-of-order", false, "Permitir aplicar migraciones fuera de orden")
	showStatus := flag.Bool("status", false, "Solo mostrar estado actual de migraciones")
	verbose := flag.Bool("verbose", false, "Mostrar logs detallados")
	timeout := flag.Int("timeout", 60, "Tiempo máximo en segundos para ejecutar las migraciones")
	
	flag.Parse()
	
	// Configurar logger
	log := logrus.New()
	if *verbose {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	
	// Crear pool de conexiones
	log.Infof("Conectando a la base de datos con configuración en: %s", *configFile)
	dbPool, err := pool.NewDBConnectionPool(*configFile)
	if err != nil {
		log.WithError(err).Fatal("Error conectando a la base de datos")
	}
	defer dbPool.Close()
	
	// Crear migrador
	migrator := migration.NewMigrator(dbPool).
		WithLogger(log).
		WithDryRun(*dryRun).
		WithStrictMode(*strict).
		WithForceVersion(*forceVersion).
		WithAllowOutOfOrder(*outOfOrder)
	
	// Cargar migraciones
	if err := migrator.LoadEmbeddedMigrations(); err != nil {
		log.WithError(err).Fatal("Error cargando migraciones")
	}
	
	// Si solo queremos ver el estado
	if *showStatus {
		showMigrationStatus(migrator, log)
		return
	}
	
	// Ejecutar migraciones con timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Second)
	defer cancel()
	
	log.Info("Iniciando proceso de migración...")
	if *dryRun {
		log.Warn("MODO DE SIMULACIÓN: No se realizarán cambios reales en la base de datos")
	}
	
	startTime := time.Now()
	if err := migrator.Run(ctx); err != nil {
		log.WithError(err).Fatal("Error ejecutando migraciones")
	}
	
	elapsedTime := time.Since(startTime)
	log.WithField("elapsed_time", elapsedTime.String()).Info("Proceso de migración completado")
	
	// Mostrar estado final
	showMigrationStatus(migrator, log)
}

func showMigrationStatus(migrator *migration.Migrator, log *logrus.Logger) {
	migrations := migrator.Status()
	
	if len(migrations) == 0 {
		fmt.Println("No hay migraciones disponibles")
		return
	}
	
	fmt.Println("\nEstado actual de las migraciones:")
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
		
		// Truncar descripción si es muy larga
		desc := m.Description
		if len(desc) > 40 {
			desc = desc[:37] + "..."
		}
		
		fmt.Printf("| %-8d | %-13s | %-38s | %-8s |\n", 
			m.Version, estado, desc, aplicada)
	}
	
	fmt.Println("+----------+---------------+----------------------------------------+----------+")
	
	// Mostrar versión actual
	currentVersion := migrator.GetCurrentVersion()
	fmt.Printf("\nVersión actual de la base de datos: %d\n\n", currentVersion)
}
