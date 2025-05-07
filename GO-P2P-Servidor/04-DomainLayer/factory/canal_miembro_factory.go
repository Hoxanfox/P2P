package factory

import (
	"github.com/google/uuid"
	"model"
)

// ICanalMiembroFactory define la creación de CanalMiembro.
type ICanalMiembroFactory interface {
	// Create genera un CanalMiembro con el rol especificado.
	Create(canalID, usuarioID uuid.UUID, rol string) (*model.CanalMiembro, error)
	
	// CreateOwner genera un CanalMiembro con el rol "owner".
	CreateOwner(canalID, usuarioID uuid.UUID) (*model.CanalMiembro, error)
	
	// CreateMember genera un CanalMiembro con el rol "member".
	CreateMember(canalID, usuarioID uuid.UUID) (*model.CanalMiembro, error)
}

// canalMiembroFactory es la implementación de ICanalMiembroFactory.
type canalMiembroFactory struct{}

// NewCanalMiembroFactory devuelve una instancia de ICanalMiembroFactory.
func NewCanalMiembroFactory() ICanalMiembroFactory {
	return &canalMiembroFactory{}
}

// Create implementa ICanalMiembroFactory: crea un CanalMiembro con el rol especificado.
func (f *canalMiembroFactory) Create(canalID, usuarioID uuid.UUID, rol string) (*model.CanalMiembro, error) {
	return model.NewCanalMiembro(canalID, usuarioID, rol)
}

// CreateOwner implementa ICanalMiembroFactory: crea un CanalMiembro con rol "owner".
func (f *canalMiembroFactory) CreateOwner(canalID, usuarioID uuid.UUID) (*model.CanalMiembro, error) {
	return f.Create(canalID, usuarioID, "owner")
}

// CreateMember implementa ICanalMiembroFactory: crea un CanalMiembro con rol "member".
func (f *canalMiembroFactory) CreateMember(canalID, usuarioID uuid.UUID) (*model.CanalMiembro, error) {
	return f.Create(canalID, usuarioID, "member")
}
