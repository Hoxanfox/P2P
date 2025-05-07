package model

// NodoEstado indica si un nodo P2P está conectado o desconectado.
type NodoEstado string

const (
    NodoConectado    NodoEstado = "CONECTADO"
    NodoDesconectado NodoEstado = "DESCONECTADO"
)

// Valid comprueba que el NodoEstado sea uno de los valores admitidos.
func (n NodoEstado) Valid() bool {
    switch n {
    case NodoConectado, NodoDesconectado:
        return true
    default:
        return false
    }
}
