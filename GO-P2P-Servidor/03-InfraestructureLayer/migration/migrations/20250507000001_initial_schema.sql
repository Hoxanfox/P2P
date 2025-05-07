/*--------------------------------------------------------------------
  Migración inicial para crear el esquema de la base de datos
--------------------------------------------------------------------*/

/*--------------------------------------------------------------------
  ENUM simulados con CHECK  (MySQL 8.0 admite CHECK CONSTRAINT)
--------------------------------------------------------------------*/
-- CanalTipo : PUBLICO | PRIVADO
-- EventoTipo : LOGIN | MENSAJE | ARCHIVO | CANAL
-- NodoEstado : CONECTADO | DESCONECTADO
-- EstadoInvitacion : PENDIENTE | ACEPTADA | RECHAZADA

/*--------------------------------------------------------------------
  Tabla  usuario_servidor
--------------------------------------------------------------------*/
CREATE TABLE IF NOT EXISTS usuario_servidor (
  id                CHAR(36)      PRIMARY KEY,
  nombre_usuario    VARCHAR(255)  NOT NULL UNIQUE,
  email             VARCHAR(255)  NOT NULL UNIQUE,
  contrasena_hash   VARCHAR(255)  NOT NULL,
  foto_url          VARCHAR(512),
  ip_registrada     VARCHAR(45)   NOT NULL,
  fecha_registro    TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  is_connected      BOOLEAN       NOT NULL DEFAULT FALSE
);

/*--------------------------------------------------------------------
  Tabla canal_servidor
--------------------------------------------------------------------*/
CREATE TABLE IF NOT EXISTS canal_servidor (
  id           CHAR(36)     PRIMARY KEY,
  nombre       VARCHAR(255) NOT NULL UNIQUE,
  descripcion  TEXT,
  tipo         VARCHAR(10)  NOT NULL,
  CHECK (tipo IN ('PUBLICO','PRIVADO'))
);

/*--------------------------------------------------------------------
  Tabla canal_miembro  (relación N-N usuario-canal con rol)
--------------------------------------------------------------------*/
CREATE TABLE IF NOT EXISTS canal_miembro (
  usuario_id  CHAR(36) NOT NULL,
  canal_id    CHAR(36) NOT NULL,
  rol         VARCHAR(50) NOT NULL,
  PRIMARY KEY (usuario_id , canal_id),
  FOREIGN KEY (usuario_id) REFERENCES usuario_servidor(id) ON DELETE CASCADE,
  FOREIGN KEY (canal_id)   REFERENCES canal_servidor(id)   ON DELETE CASCADE
);

/*--------------------------------------------------------------------
  Tabla invitacion_canal
--------------------------------------------------------------------*/
CREATE TABLE IF NOT EXISTS invitacion_canal (
  id              CHAR(36)    PRIMARY KEY,
  canal_id        CHAR(36)    NOT NULL,
  destinatario_id CHAR(36)    NOT NULL,
  estado          VARCHAR(10) NOT NULL,
  fecha_envio     TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CHECK (estado IN ('PENDIENTE','ACEPTADA','RECHAZADA')),
  FOREIGN KEY (canal_id)        REFERENCES canal_servidor(id)   ON DELETE CASCADE,
  FOREIGN KEY (destinatario_id) REFERENCES usuario_servidor(id) ON DELETE CASCADE
);

/*--------------------------------------------------------------------
  Tabla notificacion
--------------------------------------------------------------------*/
CREATE TABLE IF NOT EXISTS notificacion (
  id             CHAR(36)   PRIMARY KEY,
  usuario_id     CHAR(36)   NOT NULL,
  contenido      TEXT       NOT NULL,
  fecha          TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  leido          BOOLEAN    NOT NULL DEFAULT FALSE,
  invitacion_id  CHAR(36),
  FOREIGN KEY (usuario_id)    REFERENCES usuario_servidor(id) ON DELETE CASCADE,
  FOREIGN KEY (invitacion_id) REFERENCES invitacion_canal(id) ON DELETE SET NULL
);

/*--------------------------------------------------------------------
  Tabla chat_privado
--------------------------------------------------------------------*/
CREATE TABLE IF NOT EXISTS chat_privado (
  id CHAR(36) PRIMARY KEY
);

-- Miembros del chat 1-a-1 (permite extender a >2 si hiciera falta)
CREATE TABLE IF NOT EXISTS chat_privado_usuario (
  chat_id    CHAR(36) NOT NULL,
  usuario_id CHAR(36) NOT NULL,
  PRIMARY KEY (chat_id , usuario_id),
  FOREIGN KEY (chat_id)    REFERENCES chat_privado(id)       ON DELETE CASCADE,
  FOREIGN KEY (usuario_id) REFERENCES usuario_servidor(id)   ON DELETE CASCADE
);

/*--------------------------------------------------------------------
  Tabla mensaje_servidor
--------------------------------------------------------------------*/
CREATE TABLE IF NOT EXISTS mensaje_servidor (
  id                 CHAR(36)   PRIMARY KEY,
  remitente_id       CHAR(36)   NOT NULL,
  destino_usuario_id CHAR(36),           -- NULL si es mensaje de canal
  canal_id           CHAR(36),           -- NULL si es mensaje directo
  chat_privado_id    CHAR(36),           -- NULL para canal
  contenido          TEXT       NOT NULL,
  timestamp          TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  archivo_id         CHAR(36),
  -- Reglas de validación básicas
  CHECK (
        (destino_usuario_id IS NULL AND canal_id IS NOT NULL)
     OR (destino_usuario_id IS NOT NULL AND canal_id IS NULL)
  ),
  FOREIGN KEY (remitente_id)       REFERENCES usuario_servidor(id) ON DELETE CASCADE,
  FOREIGN KEY (destino_usuario_id) REFERENCES usuario_servidor(id) ON DELETE CASCADE,
  FOREIGN KEY (canal_id)           REFERENCES canal_servidor(id)   ON DELETE CASCADE,
  FOREIGN KEY (chat_privado_id)    REFERENCES chat_privado(id)     ON DELETE CASCADE
);

/*--------------------------------------------------------------------
  Tabla archivo_metadata
--------------------------------------------------------------------*/
CREATE TABLE IF NOT EXISTS archivo_metadata (
  id              CHAR(36)   PRIMARY KEY,
  nombre_original VARCHAR(255) NOT NULL,
  tamano_bytes    BIGINT       NOT NULL,
  ruta_almacen    VARCHAR(1024) NOT NULL,
  subido_por      CHAR(36)     NOT NULL,
  fecha_subida    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (subido_por) REFERENCES usuario_servidor(id) ON DELETE CASCADE
);

-- Relación 1-1 (mensaje-archivo) se maneja con UNIQUE
ALTER TABLE mensaje_servidor
  ADD CONSTRAINT fk_msg_archivo
  FOREIGN KEY (archivo_id) REFERENCES archivo_metadata(id) ON DELETE SET NULL,
  ADD UNIQUE (archivo_id);

/*--------------------------------------------------------------------
  Tabla log_entry
--------------------------------------------------------------------*/
CREATE TABLE IF NOT EXISTS log_entry (
  id          CHAR(36)   PRIMARY KEY,
  tipo_evento VARCHAR(10) NOT NULL,
  detalle     TEXT       NOT NULL,
  timestamp   TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  usuario_id  CHAR(36),
  CHECK (tipo_evento IN ('LOGIN','MENSAJE','ARCHIVO','CANAL')),
  FOREIGN KEY (usuario_id) REFERENCES usuario_servidor(id) ON DELETE SET NULL
);

/*--------------------------------------------------------------------
  Tabla configuracion_servidor  (1 sola fila típica)
--------------------------------------------------------------------*/
CREATE TABLE IF NOT EXISTS configuracion_servidor (
  id               INT          PRIMARY KEY DEFAULT 1,
  max_conexiones   INT          NOT NULL,
  parametros_mysql TEXT         NOT NULL,
  rutas_logs       TEXT         NOT NULL
);

/*--------------------------------------------------------------------
  Tablas P2P
--------------------------------------------------------------------*/
CREATE TABLE IF NOT EXISTS peer (
  id_nodo   CHAR(36)   PRIMARY KEY,
  direccion VARCHAR(255) NOT NULL,
  estado    VARCHAR(15) NOT NULL,
  CHECK (estado IN ('CONECTADO','DESCONECTADO'))
);

CREATE TABLE IF NOT EXISTS heartbeat_log (
  id          CHAR(36)  PRIMARY KEY,
  nodo_id     CHAR(36)  NOT NULL,
  enviado_at  TIMESTAMP NOT NULL,
  recibido_at TIMESTAMP NOT NULL,
  FOREIGN KEY (nodo_id) REFERENCES peer(id_nodo) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS replica_event (
  id            CHAR(36)  PRIMARY KEY,
  entidad_tipo  VARCHAR(30) NOT NULL,
  entidad_id    CHAR(36)    NOT NULL,
  evento_at     TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  origen_nodo_id CHAR(36)   NOT NULL,
  FOREIGN KEY (origen_nodo_id) REFERENCES peer(id_nodo) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS routed_message (
  mensaje_id      CHAR(36)  NOT NULL,
  nodo_destino_id CHAR(36)  NOT NULL,
  enruta_at       TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (mensaje_id , nodo_destino_id),
  FOREIGN KEY (mensaje_id)      REFERENCES mensaje_servidor(id) ON DELETE CASCADE,
  FOREIGN KEY (nodo_destino_id) REFERENCES peer(id_nodo)        ON DELETE CASCADE
);

/*--------------------------------------------------------------------
  Índices de apoyo
--------------------------------------------------------------------*/
CREATE INDEX idx_usuario_email          ON usuario_servidor(email);
CREATE INDEX idx_msg_remitente          ON mensaje_servidor(remitente_id);
CREATE INDEX idx_msg_canal              ON mensaje_servidor(canal_id);
CREATE INDEX idx_msg_chat               ON mensaje_servidor(chat_privado_id);
CREATE INDEX idx_notif_usuario          ON notificacion(usuario_id);
CREATE INDEX idx_invitacion_destinatario ON invitacion_canal(destinatario_id);
CREATE INDEX idx_archivo_subido_por     ON archivo_metadata(subido_por);
CREATE INDEX idx_peer_estado            ON peer(estado);
