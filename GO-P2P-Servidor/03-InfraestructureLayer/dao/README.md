# Módulo DAO (Data Access Objects)

Este módulo contiene los objetos de acceso a datos para el sistema P2P, implementando el patrón DAO para abstraer y encapsular todo el acceso a la fuente de datos.

## Estructura

El módulo dao es parte de la capa de infraestructura y proporciona una interfaz unificada para interactuar con la base de datos MySQL.

## Dependencias

- **model**: Capa de dominio que contiene las entidades del sistema (`04-DomainLayer/model`)
- **pool**: Gestiona conexiones a la base de datos con manejo eficiente de recursos (`03-InfraestructureLayer/pool`)

## Implementación

Cada entidad en el modelo de dominio tiene su correspondiente implementación DAO:

- `UserMySQLDAO`: Operaciones CRUD para usuarios del servidor
- `ChannelMySQLDAO`: Gestión de canales de comunicación
- `MessageMySQLDAO`: Almacenamiento y recuperación de mensajes
- `ArchivoMySQLDAO`: Manejo de metadatos de archivos compartidos
- `ConfigMySQLDAO`: Configuración del servidor
- `PeerMySQLDAO`: Gestión de nodos pares en la red P2P
- `HeartbeatLogMySQLDAO`: Registro de señales de actividad de los nodos
- `ReplicaEventMySQLDAO`: Eventos de replicación para sincronización
- `RoutedMessageMySQLDAO`: Mensajes enrutados entre nodos
- `LogEntryMySQLDAO`: Registro de eventos del sistema

## Características

- Cada DAO implementa operaciones CRUD estándar
- Utilizan transacciones para operaciones que requieren consistencia
- Manejo adecuado de errores y excepciones
- Compatibilidad con el pool de conexiones para un uso eficiente de recursos

## Uso

```go
import (
    "dao"
    "pool"
)

func main() {
    // Obtener pool de conexiones
    dbPool := pool.GetDBConnectionPool()
    
    // Crear instancia del DAO
    userDAO := dao.NewUserDAO(dbPool)
    
    // Ejecutar operaciones con el DAO
    user, err := userDAO.FindByID(userID)
}
```
