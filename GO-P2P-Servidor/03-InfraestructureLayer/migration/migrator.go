package migration

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"pool"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

// Migration representa una migración de base de datos con versión y contenido SQL
type Migration struct {
	Version     int
	Description string
	SQL         string
	Applied     bool
	AppliedAt   time.Time
}

// Migrator es el componente central que gestiona la aplicación de migraciones
type Migrator struct {
	dbPool        *pool.DBConnectionPool
	migrations    []*Migration
	log           *logrus.Logger
	tableName     string
	dryRun        bool
	strictMode    bool
	forceVersion  int
	allowOutOfOrder bool
}

// NewMigrator crea una nueva instancia del migrador
func NewMigrator(dbPool *pool.DBConnectionPool) *Migrator {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return &Migrator{
		dbPool:        dbPool,
		migrations:    make([]*Migration, 0),
		log:           logger,
		tableName:     "schema_migrations",
		strictMode:    true,
	}
}

// WithLogger configura un logger personalizado
func (m *Migrator) WithLogger(log *logrus.Logger) *Migrator {
	m.log = log
	return m
}

// WithDryRun configura el modo de simulación sin aplicar cambios
func (m *Migrator) WithDryRun(dryRun bool) *Migrator {
	m.dryRun = dryRun
	return m
}

// WithStrictMode configura si se debe fallar en caso de error
func (m *Migrator) WithStrictMode(strict bool) *Migrator {
	m.strictMode = strict
	return m
}

// WithForceVersion fuerza una versión específica (para rollbacks)
func (m *Migrator) WithForceVersion(version int) *Migrator {
	m.forceVersion = version
	return m
}

// WithAllowOutOfOrder permite aplicar migraciones en orden diferente al de versión
func (m *Migrator) WithAllowOutOfOrder(allow bool) *Migrator {
	m.allowOutOfOrder = allow
	return m
}

// LoadEmbeddedMigrations carga las migraciones desde los archivos SQL embebidos
func (m *Migrator) LoadEmbeddedMigrations() error {
	entries, err := fs.ReadDir(migrationsFS, "migrations")
	if err != nil {
		return fmt.Errorf("error leyendo directorio de migraciones: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
			file := entry.Name()
			version, description, err := parseFilename(file)
			if err != nil {
				m.log.WithError(err).Warnf("Omitiendo archivo inválido: %s", file)
				continue
			}

			content, err := fs.ReadFile(migrationsFS, filepath.Join("migrations", file))
			if err != nil {
				return fmt.Errorf("error leyendo archivo de migración %s: %w", file, err)
			}

			m.migrations = append(m.migrations, &Migration{
				Version:     version,
				Description: description,
				SQL:         string(content),
				Applied:     false,
			})
		}
	}

	// Ordenar por versión
	sort.Slice(m.migrations, func(i, j int) bool {
		return m.migrations[i].Version < m.migrations[j].Version
	})

	return nil
}

// LoadMigrationsFromDir carga las migraciones de un directorio del sistema de archivos
// Este método es útil para pruebas y entornos de desarrollo
func (m *Migrator) LoadMigrationsFromDir(dirPath string) error {
	// Verificar que el directorio existe
	info, err := os.Stat(dirPath)
	if err != nil {
		return fmt.Errorf("error accediendo al directorio %s: %w", dirPath, err)
	}

	if !info.IsDir() {
		return fmt.Errorf("%s no es un directorio", dirPath)
	}

	// Leer entradas del directorio
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("error leyendo directorio %s: %w", dirPath, err)
	}

	// Si no hay archivos SQL, reportar error
	sqlFound := false
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
			sqlFound = true
			break
		}
	}

	if !sqlFound {
		return fmt.Errorf("no se encontraron archivos SQL en %s", dirPath)
	}

	m.log.Infof("Cargando migraciones desde directorio: %s", dirPath)

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}

		filename := entry.Name()

		// Parsear nombre para obtener versión y descripción
		version, description, err := parseFilename(filename)
		if err != nil {
			m.log.WithError(err).Warnf("Ignorando archivo con formato inválido: %s", filename)
			continue
		}

		// Leer contenido del archivo
		filePath := filepath.Join(dirPath, filename)
		sqlContent, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("error leyendo archivo de migración %s: %w", filename, err)
		}

		// Crear estructura de migración
		migration := &Migration{
			Version:     version,
			Description: description,
			SQL:         string(sqlContent),
			Applied:     false,
		}

		m.migrations = append(m.migrations, migration)
		m.log.Debugf("Cargada migración: %d %s", version, description)
	}

	// Ordenar migraciones por versión
	sort.Slice(m.migrations, func(i, j int) bool {
		return m.migrations[i].Version < m.migrations[j].Version
	})

	m.log.Infof("Cargadas %d migraciones desde directorio", len(m.migrations))
	return nil
}

// LoadExternalMigrations carga migraciones desde archivos en un directorio externo
func (m *Migrator) LoadExternalMigrations(dirPath string) error {
	entries, err := filepath.Glob(filepath.Join(dirPath, "*.sql"))
	if err != nil {
		return fmt.Errorf("error buscando archivos SQL en %s: %w", dirPath, err)
	}

	for _, entry := range entries {
		file := filepath.Base(entry)
		version, description, err := parseFilename(file)
		if err != nil {
			m.log.WithError(err).Warnf("Omitiendo archivo inválido: %s", file)
			continue
		}

		content, err := fs.ReadFile(migrationsFS, filepath.Join("migrations", file))
		if err != nil {
			return fmt.Errorf("error leyendo archivo de migración %s: %w", file, err)
		}

		m.migrations = append(m.migrations, &Migration{
			Version:     version,
			Description: description,
			SQL:         string(content),
			Applied:     false,
		})
	}

	// Ordenar por versión
	sort.Slice(m.migrations, func(i, j int) bool {
		return m.migrations[i].Version < m.migrations[j].Version
	})

	return nil
}

// CreateMigrationsTable crea la tabla de control de migraciones si no existe
func (m *Migrator) CreateMigrationsTable(ctx context.Context) error {
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			version     BIGINT       NOT NULL PRIMARY KEY,
			description VARCHAR(255) NOT NULL,
			applied_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`, m.tableName)

	if m.dryRun {
		m.log.Info("Modo simulación: SQL que se ejecutaría para crear tabla de migraciones:")
		m.log.Info(query)
		return nil
	}

	_, err := m.dbPool.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error creando tabla de control de migraciones: %w", err)
	}

	m.log.Infof("Tabla de migraciones '%s' verificada/creada", m.tableName)
	return nil
}

// LoadAppliedMigrations carga el estado de migraciones aplicadas desde la BD
func (m *Migrator) LoadAppliedMigrations(ctx context.Context) error {
	query := fmt.Sprintf("SELECT version, description, applied_at FROM %s", m.tableName)
	
	rows, err := m.dbPool.QueryContext(ctx, query)
	if err != nil {
		// Si el error es que la tabla no existe, intentamos crearla
		if strings.Contains(err.Error(), "doesn't exist") {
			m.log.Info("Tabla de migraciones no encontrada, creándola...")
			if err := m.CreateMigrationsTable(ctx); err != nil {
				return err
			}
			return nil // No hay migraciones aplicadas aún
		}
		return fmt.Errorf("error consultando migraciones aplicadas: %w", err)
	}
	defer rows.Close()

	// Crear un mapa de versiones aplicadas
	applied := make(map[int]time.Time)
	for rows.Next() {
		var version int
		var description string
		var appliedAt time.Time
		if err := rows.Scan(&version, &description, &appliedAt); err != nil {
			return fmt.Errorf("error leyendo migración aplicada: %w", err)
		}
		applied[version] = appliedAt
	}

	// Actualizar el estado de las migraciones cargadas
	for _, migration := range m.migrations {
		if appliedAt, exists := applied[migration.Version]; exists {
			migration.Applied = true
			migration.AppliedAt = appliedAt
		}
	}

	m.log.Infof("Migraciones previas cargadas: %d aplicadas", len(applied))
	return nil
}

// Run ejecuta las migraciones pendientes
func (m *Migrator) Run(ctx context.Context) error {
	// Verificar que tenemos migraciones
	if len(m.migrations) == 0 {
		m.log.Warn("No hay migraciones para aplicar")
		return nil
	}

	// Primero crear tabla de migraciones si no existe
	if err := m.CreateMigrationsTable(ctx); err != nil {
		return err
	}

	// Cargar estado actual de las migraciones
	if err := m.LoadAppliedMigrations(ctx); err != nil {
		return err
	}

	// Contadores para log
	totalCount := len(m.migrations)
	appliedCount := 0
	pendingCount := 0
	newAppliedCount := 0
	
	for _, migration := range m.migrations {
		if migration.Applied {
			appliedCount++
			continue
		}
		
		pendingCount++
		
		// Si es modo forzado, verificar versión
		if m.forceVersion > 0 && migration.Version > m.forceVersion {
			m.log.Infof("Omitiendo migración %d (%s) por encima de versión forzada %d", 
				migration.Version, migration.Description, m.forceVersion)
			continue
		}
		
		// Aplicar la migración
		if err := m.applyMigration(ctx, migration); err != nil {
			if m.strictMode {
				return fmt.Errorf("error aplicando migración %d: %w", migration.Version, err)
			}
			m.log.WithError(err).Errorf("Error aplicando migración %d, continuando por modo no estricto", migration.Version)
			continue
		}
		
		newAppliedCount++
	}
	
	m.log.WithFields(logrus.Fields{
		"total": totalCount,
		"previously_applied": appliedCount,
		"pending": pendingCount,
		"newly_applied": newAppliedCount,
	}).Info("Proceso de migración completado")
	
	return nil
}

// applyMigration aplica una migración específica
func (m *Migrator) applyMigration(ctx context.Context, migration *Migration) error {
	m.log.Infof("Aplicando migración %d: %s", migration.Version, migration.Description)
	
	// Modo simulación, solo mostrar SQL
	if m.dryRun {
		m.log.Info("Modo simulación: SQL que se ejecutaría:")
		m.log.Info(migration.SQL)
		return nil
	}
	
	// Iniciar transacción
	tx, err := m.dbPool.DB().BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error iniciando transacción: %w", err)
	}
	
	// Asegurar que hacemos rollback en caso de error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	
	// Ejecutar el SQL de la migración dividiendo por instrucciones
	// y eliminando comentarios para no confundir los delimitadores
	lines := strings.Split(migration.SQL, "\n")
	var cleanedLines []string
	
	inCommentBlock := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Manejo de bloques de comentarios
		if strings.HasPrefix(line, "/*") {
			inCommentBlock = true
			continue
		}
		
		if inCommentBlock {
			if strings.Contains(line, "*/") {
				inCommentBlock = false
			}
			continue
		}
		
		// Ignorar líneas de comentario
		if strings.HasPrefix(line, "--") {
			continue
		}
		
		// Eliminar comentarios al final de línea
		commentIdx := strings.Index(line, "--")
		if commentIdx >= 0 {
			line = line[:commentIdx]
		}
		
		// Añadir la línea si no está vacía
		if line != "" {
			cleanedLines = append(cleanedLines, line)
		}
	}
	
	// Unir todas las líneas en un solo string
	cleanedSQL := strings.Join(cleanedLines, " ")
	
	// Separar correctamente las instrucciones SQL respetando las comillas
	statements := m.splitSQLStatements(cleanedSQL)
	
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		
		m.log.Debugf("Ejecutando instrucción:\n%s", stmt)
		
		if _, err = tx.ExecContext(ctx, stmt); err != nil {
			return fmt.Errorf("error ejecutando SQL: %w\n\nEn instrucción:\n%s", err, stmt)
		}
	}
	
	// Registrar la migración como aplicada
	insertQuery := fmt.Sprintf("INSERT INTO %s (version, description) VALUES (?, ?)", m.tableName)
	if _, err = tx.ExecContext(ctx, insertQuery, migration.Version, migration.Description); err != nil {
		return fmt.Errorf("error registrando migración como aplicada: %w", err)
	}
	
	// Confirmar transacción
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error confirmando transacción: %w", err)
	}
	
	// Actualizar estado de la migración
	migration.Applied = true
	migration.AppliedAt = time.Now()
	
	m.log.Infof("Migración %d aplicada exitosamente", migration.Version)
	return nil
}

// rollbackMigration revierte una migración específica (si está implementado el down)
func (m *Migrator) rollbackMigration(ctx context.Context, migration *Migration) error {
	// Implementación de rollback para futuras versiones
	return fmt.Errorf("rollback no implementado")
}

// Status devuelve el estado actual de las migraciones
func (m *Migrator) Status() []*Migration {
	return m.migrations
}

// parseFilename extrae versión y descripción del nombre del archivo
func parseFilename(filename string) (int, string, error) {
	// Verificar extensión
	if !strings.HasSuffix(filename, ".sql") {
		return 0, "", fmt.Errorf("archivo %s no tiene extensión .sql", filename)
	}
	
	// Extraer parte sin extensión
	basename := strings.TrimSuffix(filename, ".sql")
	
	// Separar por primer underscore
	parts := strings.SplitN(basename, "_", 2)
	if len(parts) != 2 {
		return 0, "", fmt.Errorf("formato de nombre inválido: %s", filename)
	}
	
	// Verificar que la descripción no esté vacía
	if parts[1] == "" {
		return 0, "", fmt.Errorf("descripción vacía en nombre %s", filename)
	}
	
	// Convertir versión a entero
	version, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", fmt.Errorf("versión inválida en nombre %s: %w", filename, err)
	}
	
	// Normalizar descripción reemplazando underscores con espacios
	description := strings.ReplaceAll(parts[1], "_", " ")
	
	return version, description, nil
}

// GetCurrentVersion devuelve la versión más alta aplicada
func (m *Migrator) GetCurrentVersion() int {
	currentVersion := 0
	for _, migration := range m.migrations {
		if migration.Applied && migration.Version > currentVersion {
			currentVersion = migration.Version
		}
	}
	return currentVersion
}

// GetMigrationByVersion busca una migración por su versión
func (m *Migrator) GetMigrationByVersion(version int) *Migration {
	for _, migration := range m.migrations {
		if migration.Version == version {
			return migration
		}
	}
	return nil
}

// splitSQLStatements divide un string de SQL en instrucciones individuales
// respetando las comillas para no dividir incorrectamente por punto y coma
// dentro de cadenas de texto
func (m *Migrator) splitSQLStatements(sql string) []string {
	var statements []string
	var currentStmt strings.Builder
	
	// Estados para el parsing
	inQuote := false
	lastChar := ' '
	
	for _, char := range sql {
		// Manejar comillas (detectar cuando estamos dentro de una cadena)
		if char == '\'' && lastChar != '\\' {
			inQuote = !inQuote
		}
		
		// Si encontramos un punto y coma fuera de comillas, es un separador de instrucciones
		if char == ';' && !inQuote {
			currentStmt.WriteRune(char) // Incluir el punto y coma en la instrucción
			stmtText := strings.TrimSpace(currentStmt.String())
			
			if stmtText != "" && stmtText != ";" {
				statements = append(statements, stmtText)
			}
			
			currentStmt.Reset()
		} else {
			currentStmt.WriteRune(char)
		}
		
		lastChar = char
	}
	
	// No olvidar la última instrucción si no termina con punto y coma
	finalStmt := strings.TrimSpace(currentStmt.String())
	if finalStmt != "" {
		// Asegurarse de que termine con punto y coma
		if !strings.HasSuffix(finalStmt, ";") {
			finalStmt += ";"
		}
		statements = append(statements, finalStmt)
	}
	
	return statements
}

// DropSchemaForTesting elimina todas las tablas (solo para pruebas)
func (m *Migrator) DropSchemaForTesting(ctx context.Context) error {
	// Este método solo debe usarse en entornos de prueba
	// Verificar si estamos conectados a MySQL
	var dbName string
	if err := m.dbPool.DB().QueryRowContext(ctx, "SELECT DATABASE()").Scan(&dbName); err != nil {
		return fmt.Errorf("error verificando base de datos: %w", err)
	}
	
	// Solo procedemos si obtuvimos un nombre de base de datos
	if dbName == "" {
		return fmt.Errorf("método solo soportado para MySQL con base de datos seleccionada")
	}
	
	// Obtener lista de tablas
	query := `
		SELECT table_name 
		FROM information_schema.tables 
		WHERE table_schema = DATABASE()
	`
	
	rows, err := m.dbPool.QueryContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error consultando tablas: %w", err)
	}
	defer rows.Close()
	
	// Recopilar nombres de tablas
	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return fmt.Errorf("error leyendo nombre de tabla: %w", err)
		}
		tables = append(tables, tableName)
	}
	
	if len(tables) == 0 {
		return nil // No hay tablas para eliminar
	}
	
	// Desactivar verificación de foreign keys
	if _, err := m.dbPool.ExecContext(ctx, "SET FOREIGN_KEY_CHECKS = 0"); err != nil {
		return fmt.Errorf("error desactivando foreign key checks: %w", err)
	}
	
	// Eliminar cada tabla
	for _, table := range tables {
		dropQuery := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", table)
		if _, err := m.dbPool.ExecContext(ctx, dropQuery); err != nil {
			m.log.WithError(err).Errorf("Error eliminando tabla %s", table)
		}
	}
	
	// Reactivar verificación de foreign keys
	if _, err := m.dbPool.ExecContext(ctx, "SET FOREIGN_KEY_CHECKS = 1"); err != nil {
		return fmt.Errorf("error reactivando foreign key checks: %w", err)
	}
	
	m.log.Infof("Esquema eliminado: %d tablas eliminadas", len(tables))
	return nil
}
