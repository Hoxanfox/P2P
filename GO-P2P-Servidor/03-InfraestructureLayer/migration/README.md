# Módulo Migration - Gestión de Esquema de Base de Datos

Este módulo implementa el sistema de migraciones para la aplicación P2P, con los siguientes objetivos:

## Características principales

1. **Ejecución automática de scripts DDL**: Aplicar scripts SQL para crear o alterar tablas, índices y relaciones en la base de datos.

2. **Control de versiones**: Mantener un registro de las migraciones ya aplicadas en la tabla `schema_migrations` para ejecutar solo los cambios pendientes.

3. **Inicialización de datos**: Proporcionar capacidad para insertar datos iniciales, como el usuario administrador.

4. **Reutilización de conexiones**: El módulo no abre nuevas conexiones sino que reutiliza el pool de conexiones existente.

## Componentes

- **Migrator**: Motor principal que ejecuta las migraciones en orden y mantiene el registro de las aplicadas.
- **Migration**: Representación de una migración individual con su versión, descripción y SQL.
- **Scripts de migración**: Archivos SQL integrados en el binario mediante `embed.FS`.

## Dependencias

- **Pool**: Utiliza el módulo de pool para obtener conexiones a la base de datos.
- **Model**: Referencia a entidades del dominio para mantener integridad y coherencia.
