package dao

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"pool"
)

// setupTestDB configura una conexión a la base de datos para pruebas.
// Utiliza la configuración real del sistema.
func setupTestDB(t *testing.T) *pool.DBConnectionPool {
	// En este punto verificamos si realmente queremos conectarnos a la BD
	// Para evitar efectos secundarios en entornos no preparados, vamos a
	// saltarnos las pruebas si no podemos conectar correctamente
	//t.Skip("Prueba saltada: Requiere una conexión a base de datos real configurada correctamente")
	
	// El código siguiente sólo se ejecutaría si descomentáramos la línea anterior
	// y tuviéramos configurada correctamente la base de datos
	
	// Ruta al archivo de configuración de la base de datos
	configPath := filepath.Join("..", "pool", "db_config.yaml")
	
	// Inicializar el pool de conexiones
	dbPool, err := pool.NewDBConnectionPool(configPath)
	if err != nil {
		t.Fatalf("Error inicializando pool de conexiones: %v", err)
	}
	
	// Verificar que la conexión funciona usando un contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	err = dbPool.PingContext(ctx)
	if err != nil {
		t.Fatalf("Error conectando a la base de datos: %v", err)
	}

	// Inicializar las tablas necesarias para las pruebas
	err = createTestTables(dbPool, ctx)
	if err != nil {
		t.Fatalf("Error creando tablas para pruebas: %v", err)
	}
	
	return dbPool
}

// cleanupTestDB realiza limpieza después de las pruebas
func cleanupTestDB(t *testing.T, dbPool *pool.DBConnectionPool) {
	// Nota: Si quisiéramos eliminar las tablas de prueba, descomentar este código:
	/*
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dropTestTables(dbPool, ctx)
	*/

	// Por ahora, no eliminamos las tablas para no interferir con otras pruebas

	// Cerrar la conexión
	if dbPool != nil {
		if err := dbPool.Close(); err != nil {
			t.Logf("Error cerrando pool de conexiones: %v", err)
		}
	}
}

// createTestTables crea las tablas necesarias para las pruebas
func createTestTables(dbPool *pool.DBConnectionPool, ctx context.Context) error {
	// Primero eliminar las tablas en orden inverso para manejar las dependencias
	dropTables := []string{
		// 1. Primero eliminar tablas dependientes
		`DROP TABLE IF EXISTS chat_privado_usuarios;`,
		`DROP TABLE IF EXISTS invitaciones_canal;`, 
		`DROP TABLE IF EXISTS canal_miembros;`,
		`DROP TABLE IF EXISTS notificaciones;`,
		// 2. Luego eliminar tablas principales
		`DROP TABLE IF EXISTS chats_privados;`,
		`DROP TABLE IF EXISTS canales;`,
	}

	// Ejecutar cada sentencia DROP TABLE
	for _, query := range dropTables {
		_, err := dbPool.ExecContext(ctx, query)
		if err != nil {
			// No consideramos un error si la tabla no existe
			return fmt.Errorf("error eliminando tabla: %w", err)
		}
	}

	// Lista de sentencias SQL para crear las tablas requeridas en orden correcto
	createTables := []string{
		// 1. Primero crear tablas principales
		// Tabla de canales
		`CREATE TABLE IF NOT EXISTS canales (
			id VARCHAR(36) PRIMARY KEY,
			nombre VARCHAR(100) NOT NULL,
			descripcion TEXT,
			tipo VARCHAR(20) NOT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,

		// Tabla de chats privados
		`CREATE TABLE IF NOT EXISTS chats_privados (
			id VARCHAR(36) PRIMARY KEY
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,

		// 2. Luego crear tablas dependientes
		// Tabla de miembros de canales
		`CREATE TABLE IF NOT EXISTS canal_miembros (
			canal_id VARCHAR(36) NOT NULL,
			usuario_id VARCHAR(36) NOT NULL,
			rol VARCHAR(20) NOT NULL,
			PRIMARY KEY (canal_id, usuario_id),
			FOREIGN KEY (canal_id) REFERENCES canales(id) ON DELETE CASCADE
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,

		// Tabla de invitaciones a canales
		`CREATE TABLE IF NOT EXISTS invitaciones_canal (
			id VARCHAR(36) PRIMARY KEY,
			canal_id VARCHAR(36) NOT NULL,
			destinatario_id VARCHAR(36) NOT NULL,
			estado VARCHAR(20) NOT NULL,
			fecha_envio DATETIME NOT NULL,
			FOREIGN KEY (canal_id) REFERENCES canales(id) ON DELETE CASCADE
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,

		// Tabla de notificaciones
		`CREATE TABLE IF NOT EXISTS notificaciones (
			id VARCHAR(36) PRIMARY KEY,
			usuario_id VARCHAR(36) NOT NULL,
			contenido TEXT NOT NULL,
			fecha DATETIME NOT NULL,
			leido BOOLEAN NOT NULL DEFAULT false,
			invitacion_id VARCHAR(36)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,

		// Tabla de usuarios de chats privados
		`CREATE TABLE IF NOT EXISTS chat_privado_usuarios (
			chat_privado_id VARCHAR(36) NOT NULL,
			usuario_id VARCHAR(36) NOT NULL,
			PRIMARY KEY (chat_privado_id, usuario_id),
			FOREIGN KEY (chat_privado_id) REFERENCES chats_privados(id) ON DELETE CASCADE
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
	}

	// Ejecutar cada sentencia CREATE TABLE
	for _, query := range createTables {
		_, err := dbPool.ExecContext(ctx, query)
		if err != nil {
			return fmt.Errorf("error creando tabla: %w", err)
		}
	}

	return nil
}

// dropTestTables elimina las tablas creadas para las pruebas
func dropTestTables(dbPool *pool.DBConnectionPool, ctx context.Context) error {
	// Lista de tablas a eliminar
	dropTables := []string{
		"DROP TABLE IF EXISTS notificaciones;",
	}

	// Ejecutar cada sentencia DROP TABLE
	for _, query := range dropTables {
		_, err := dbPool.ExecContext(ctx, query)
		if err != nil {
			return fmt.Errorf("error eliminando tabla: %w", err)
		}
	}

	return nil
}
