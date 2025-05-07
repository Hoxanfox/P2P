/*--------------------------------------------------------------------
  Migración para insertar datos iniciales
--------------------------------------------------------------------*/

-- Insertar usuario administrador (contraseña: admin123)
-- La contraseña está hasheada con bcrypt
INSERT INTO usuario_servidor (
  id, 
  nombre_usuario, 
  email, 
  contrasena_hash, 
  ip_registrada, 
  fecha_registro
) VALUES (
  'c78f9214-12a1-4aab-8810-0152a8780a7d',  -- UUID fijo para el admin
  'admin',
  'admin@chatserver.com',
  '$2a$10$mK6X3yjyeu4GW6M2M2r3/uWRfP3OJRjb.pBJqZg/G2M9UMBJDFVrq',  -- bcrypt hash de 'admin123'
  '127.0.0.1',
  CURRENT_TIMESTAMP
);

-- Insertar canal general público por defecto
INSERT INTO canal_servidor (
  id,
  nombre,
  descripcion,
  tipo
) VALUES (
  'd2a54c86-faee-44d5-b46c-fec88e12528e',  -- UUID fijo para canal general
  'General',
  'Canal público para todos los usuarios',
  'PUBLICO'
);

-- Agregar admin como miembro del canal general
INSERT INTO canal_miembro (
  usuario_id,
  canal_id,
  rol
) VALUES (
  'c78f9214-12a1-4aab-8810-0152a8780a7d',  -- ID del admin
  'd2a54c86-faee-44d5-b46c-fec88e12528e',  -- ID del canal general
  'ADMIN'
);

-- Insertar configuración inicial del servidor
INSERT INTO configuracion_servidor (
  id,
  max_conexiones,
  parametros_mysql,
  rutas_logs
) VALUES (
  1,  -- ID fijo para configuración
  1000,  -- Máx. conexiones por defecto
  'max_connections=100;wait_timeout=300',  -- Parámetros MySQL
  '/var/log/chat-p2p'  -- Ruta logs por defecto
);
