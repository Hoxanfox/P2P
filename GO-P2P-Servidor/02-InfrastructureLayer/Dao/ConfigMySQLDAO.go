package Dao

import (
	"database/sql"
	"errors"
	"model"
)

// ConfigMySQLDAO handles database operations for ConfiguracionServidor entities
type ConfigMySQLDAO struct {
	db *sql.DB
}

// NewConfigMySQLDAO creates a new ConfigMySQLDAO instance
func NewConfigMySQLDAO(db *sql.DB) *ConfigMySQLDAO {
	return &ConfigMySQLDAO{db: db}
}

// Save persists a configuration to the database (create or update)
func (dao *ConfigMySQLDAO) Save(config *model.ConfiguracionServidor) error {
	query := `REPLACE INTO configuracion_servidor 
              (max_conexiones, parametros_mysql, rutas_logs)
              VALUES (?, ?, ?)`

	_, err := dao.db.Exec(
		query,
		config.MaxConexiones(),
		config.ParametrosMySQL(),
		config.RutasLogs(),
	)

	return err
}

// GetConfig retrieves the server configuration from the database
func (dao *ConfigMySQLDAO) GetConfig() (*model.ConfiguracionServidor, error) {
	query := `SELECT max_conexiones, parametros_mysql, rutas_logs 
              FROM configuracion_servidor LIMIT 1`

	row := dao.db.QueryRow(query)
	return dao.scanConfig(row)
}

// Update updates an existing configuration
func (dao *ConfigMySQLDAO) Update(config *model.ConfiguracionServidor) error {
	query := `UPDATE configuracion_servidor SET 
              max_conexiones = ?, parametros_mysql = ?, rutas_logs = ?`

	_, err := dao.db.Exec(
		query,
		config.MaxConexiones(),
		config.ParametrosMySQL(),
		config.RutasLogs(),
	)

	return err
}

// Helper method to scan a row into a configuration
func (dao *ConfigMySQLDAO) scanConfig(row *sql.Row) (*model.ConfiguracionServidor, error) {
	var (
		maxConexiones   int
		parametrosMySQL string
		rutasLogs       string
	)

	if err := row.Scan(&maxConexiones, &parametrosMySQL, &rutasLogs); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no configuration found")
		}
		return nil, err
	}

	config, err := model.NewConfiguracionServidor(
		maxConexiones,
		parametrosMySQL,
		rutasLogs,
	)

	return config, err
}
