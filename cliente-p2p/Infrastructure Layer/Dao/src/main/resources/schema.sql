-- Tabla ChatTipo
CREATE TABLE ChatTipo (
    id_chat_tipo TEXT PRIMARY KEY,
    nombre TEXT NOT NULL
);

-- Tabla Chat
CREATE TABLE Chat (
    id_chat TEXT PRIMARY KEY,
    id_chat_tipo TEXT NOT NULL,
    FOREIGN KEY (id_chat_tipo) REFERENCES ChatTipo(id_chat_tipo)
);

-- Tabla Archivos
CREATE TABLE Archivos (
    id_archivo TEXT PRIMARY KEY,
    nombre TEXT NOT NULL,
    binario BLOB
);

-- Tabla UsuariosServidor
CREATE TABLE UsuariosServidor (
    id_usuario_servidor TEXT PRIMARY KEY,
    nombre TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    foto BLOB,
    ip TEXT, -- ISO-8601 format (YYYY-MM-DD)
    estado INTEGER -- 0 or 1
);

-- Tabla MensajeServidor
CREATE TABLE MensajeServidor (
    id_mensaje_servidor TEXT PRIMARY KEY,
    id_chat TEXT NOT NULL,
    id_usuario TEXT NOT NULL,
    contenido TEXT,
    id_archivo TEXT,
    fecha_envio TEXT,
    FOREIGN KEY (id_chat) REFERENCES Chat(id_chat),
    FOREIGN KEY (id_usuario) REFERENCES UsuariosServidor(id_usuario_servidor),
    FOREIGN KEY (id_archivo) REFERENCES Archivos(id_archivo)
);

-- Tabla ChatMiembrosPrivado
CREATE TABLE ChatMiembrosPrivado (
    id_chat_miembros TEXT PRIMARY KEY,
    id_usuario_servidor TEXT NOT NULL,
    id_chat TEXT NOT NULL,
    FOREIGN KEY (id_usuario_servidor) REFERENCES UsuariosServidor(id_usuario_servidor),
    FOREIGN KEY (id_chat) REFERENCES Chat(id_chat)
);

-- Tabla CanalesServidor
CREATE TABLE CanalesServidor (
    id_canal_servidor TEXT PRIMARY KEY,
    nombre TEXT NOT NULL,
    descripcion TEXT
);

-- Tabla ChatMiembrosPublico
CREATE TABLE ChatMiembrosPublico (
    id_chat_miembros TEXT PRIMARY KEY,
    id_canal TEXT NOT NULL,
    id_chat TEXT NOT NULL,
    FOREIGN KEY (id_canal) REFERENCES CanalesServidor(id_canal_servidor),
    FOREIGN KEY (id_chat) REFERENCES Chat(id_chat)
);

-- Tabla CanalMiembros
CREATE TABLE CanalMiembros (
    id_canal_miembro TEXT PRIMARY KEY,
    id_usuario_servidor TEXT NOT NULL,
    id_canal_servidor TEXT NOT NULL,
    FOREIGN KEY (id_usuario_servidor) REFERENCES UsuariosServidor(id_usuario_servidor),
    FOREIGN KEY (id_canal_servidor) REFERENCES CanalesServidor(id_canal_servidor)
);

-- Tabla Invitacion
CREATE TABLE Invitacion (
    id_invitacion TEXT PRIMARY KEY,
    id_usuario TEXT NOT NULL,
    fecha_envio TEXT,
    estado INTEGER,
    FOREIGN KEY (id_usuario) REFERENCES UsuariosServidor(id_usuario_servidor)
);

-- Tabla InvitacionCanal
CREATE TABLE InvitacionCanal (
    id_invitacion_canal TEXT PRIMARY KEY,
    id_canal_servidor TEXT NOT NULL,
    id_invitacion TEXT NOT NULL,
    FOREIGN KEY (id_canal_servidor) REFERENCES CanalesServidor(id_canal_servidor),
    FOREIGN KEY (id_invitacion) REFERENCES Invitacion(id_invitacion)
);

-- Tabla Notificacion
CREATE TABLE Notificacion (
    id_notificacion TEXT PRIMARY KEY,
    id_invitacion TEXT NOT NULL,
    contenido TEXT,
    FOREIGN KEY (id_invitacion) REFERENCES Invitacion(id_invitacion)
);
