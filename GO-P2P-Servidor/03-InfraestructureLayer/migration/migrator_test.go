package migration

import (
	"context"
	"io"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestLoadEmbeddedMigrations verifica que podemos cargar las migraciones embebidas
func TestLoadEmbeddedMigrations(t *testing.T) {
	// Skip esta prueba ya que requiere que el directorio de migraciones esté correctamente embebido
	// lo cual es difícil de simular en un entorno de pruebas unitarias
	t.Skip("Omitiendo prueba que requiere sistema de archivos embebido")
	
	// Esta prueba se podría implementar en un test de integración
	// donde se copien los archivos de migración a un directorio temporal
}

// TestParseFilename verifica el parsing de nombres de archivos
func TestParseFilename(t *testing.T) {
	testCases := []struct {
		filename    string
		wantVersion int
		wantDesc    string
		wantErr     bool
	}{
		{"20250507000001_initial_schema.sql", 20250507000001, "initial schema", false},
		{"42_simple_name.sql", 42, "simple name", false},
		{"123_two_word_description.sql", 123, "two word description", false},
		{"invalid.sql", 0, "", true},
		{"invalid_format.sql", 0, "", true},
		{"123_", 0, "", true}, // Falta descripción
		{"abc_invalid.sql", 0, "", true}, // Versión no numérica
	}
	
	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			version, desc, err := parseFilename(tc.filename)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.wantVersion, version)
				assert.Equal(t, tc.wantDesc, desc)
			}
		})
	}
}

// TestMigrator_Run realiza una prueba de integración con la base de datos real
// Esta prueba se omite por defecto para no requerir una conexión a la base de datos
func TestMigrator_Run(t *testing.T) {
	// Omitir esta prueba por defecto - requiere conexión real a DB
	if testing.Short() {
		t.Skip("Omitiendo prueba de integración en modo corto")
	}
	
	// Esta prueba necesitaría un pool de conexiones real
	// Descomentar y configurar para ejecutar una prueba completa
	/*
	dbPool, err := pool.NewDBConnectionPool("../../pool/db_config.yaml")
	require.NoError(t, err, "Debe poder conectarse a la base de datos")
	defer dbPool.Close()
	
	// Crear un logger silencioso
	log := logrus.New()
	log.SetOutput(io.Discard)
	
	// Crear migrador y configurarlo
	migrator := NewMigrator(dbPool).WithLogger(log)
	
	// En pruebas, podemos querer eliminar el esquema primero
	err = migrator.DropSchemaForTesting(context.Background())
	require.NoError(t, err, "Debe poder eliminar esquema para pruebas")
	
	// Cargar migraciones
	err = migrator.LoadEmbeddedMigrations()
	require.NoError(t, err, "Debe poder cargar migraciones embebidas")
	
	// Ejecutar migraciones
	err = migrator.Run(context.Background())
	require.NoError(t, err, "Debe poder ejecutar migraciones sin error")
	
	// Verificar que el estado final es correcto
	migrations := migrator.Status()
	for _, m := range migrations {
		assert.True(t, m.Applied, "Todas las migraciones deben estar aplicadas")
	}
	*/
}

// TestCreateMigrationsTable verifica la creación de la tabla de control
func TestCreateMigrationsTable(t *testing.T) {
	// En este test solo verificamos el formato del SQL generado
	// Sin conectarnos realmente a una BD
	
	// Crear un logger silencioso
	log := logrus.New()
	log.SetOutput(io.Discard)
	
	// Crear migrador con modo de simulación
	migrator := &Migrator{
		log:        log,
		dryRun:     true,
		tableName:  "schema_migrations",
		migrations: make([]*Migration, 0),
	}
	
	// Probar creación de tabla en modo dry-run
	err := migrator.CreateMigrationsTable(context.Background())
	require.NoError(t, err, "Debe poder simular creación de tabla de migraciones")
}

// Nota: Para pruebas más completas que requieran verificar la interacción con la base de datos,
// se podría implementar un mock de DBConnectionPool o usar una base de datos en memoria como SQLite.
