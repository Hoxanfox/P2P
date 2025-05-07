package model

// CanalTipo define si un canal es público o privado, para validarlo en dominio y persistencia.
type CanalTipo string

const (
	CanalPublico CanalTipo = "PUBLICO"
	CanalPrivado CanalTipo = "PRIVADO"
)

// Valid comprueba que el CanalTipo sea uno de los valores esperados.
func (c CanalTipo) Valid() bool {
	return c == CanalPublico || c == CanalPrivado
}
