# Módulo Pool - Gestión de Recursos Compartidos

Este módulo implementa pools de recursos para la aplicación P2P, con los siguientes objetivos:

## Características principales

1. **Gestión centralizada de recursos de E/S**: Mantiene abiertos y reutiliza objetos costosos (sockets TCP y handles de base de datos) durante toda la vida del proceso.

2. **Aislamiento**: Los demás paquetes (DAO, Transport, ReplicaManager, HeartbeatService...) no saben abrir ni cerrar conexiones; únicamente "piden" instancias al pool y las devuelven.

3. **Configuración única**: Lee parámetros (máx. conexiones, time-outs, certificados TLS, back-off, etc.) desde configuración del servidor y/o archivos YAML. Un solo lugar que cambia si se mueve la BD o se ajustan límites concurrentes.

4. **Telemetría**: Cada pool expone contadores/medidas (InUse, Idle, Errors, BytesSent, etc.) que los servicios y el Dashboard usan para monitorización en tiempo real.


