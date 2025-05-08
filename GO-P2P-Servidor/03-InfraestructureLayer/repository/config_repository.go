package repository

import (
	"context"
	"dao"
	"model"
)

// ConfigRepository implementa la interfaz IConfigRepository del dominio
// utilizando el ConfigMySQLDAO para las operaciones de base de datos
type ConfigRepository struct {
	configDAO *dao.ConfigMySQLDAO
}

// NewConfigRepository crea una nueva instancia de ConfigRepository
func NewConfigRepository(configDAO *dao.ConfigMySQLDAO) *ConfigRepository {
	return &ConfigRepository{
		configDAO: configDAO,
	}
}

// Get recupera la configuración del servidor de la base de datos
func (r *ConfigRepository) Get(ctx context.Context) (*model.ConfiguracionServidor, error) {
	// Utiliza el DAO para obtener la configuración
	return r.configDAO.GetConfig()
}

// Update actualiza la configuración del servidor en la base de datos
func (r *ConfigRepository) Update(ctx context.Context, cfg *model.ConfiguracionServidor) error {
	// Utiliza el DAO para actualizar la configuración
	return r.configDAO.Update(cfg)
}
