-- Tabla ChatTipo
CREATE TABLE ChatTipo (
    id_chat_tipo UUID PRIMARY KEY,
    nombre VARCHAR(255)
);

-- Tabla Chat
CREATE TABLE Chat (
    id_chat UUID PRIMARY KEY,
    id_chat_tipo UUID,
    FOREIGN KEY (id_chat_tipo) REFERENCES ChatTipo(id_chat_tipo)
);

-- Tabla Archivos
CREATE TABLE Archivos (
    id_archivo UUID PRIMARY KEY,
    nombre VARCHAR(255),
    binario BLOB
);

-- Tabla UsuariosServidor
CREATE TABLE UsuariosServidor (
    id_usuario_servidor UUID PRIMARY KEY,
    nombre VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    foto BLOB,
    ip VARCHAR(255),
    fecha_registro DATE,
    estado BOOLEAN
);

-- Tabla MensajeServidor
CREATE TABLE MensajeServidor (
    id_mensaje_servidor UUID PRIMARY KEY,
    id_chat UUID,
    id_usuario UUID,
    contenido TEXT,
    id_archivo UUID,
    fecha_envio DATE,
    FOREIGN KEY (id_chat) REFERENCES Chat(id_chat),
    FOREIGN KEY (id_usuario) REFERENCES UsuariosServidor(id_usuario_servidor),
    FOREIGN KEY (id_archivo) REFERENCES Archivos(id_archivo)
);

-- Tabla ChatMiembrosPrivado
CREATE TABLE ChatMiembrosPrivado (
    id_chat_miembros UUID PRIMARY KEY,
    id_usuario_servidor UUID,
    id_chat UUID,
    FOREIGN KEY (id_usuario_servidor) REFERENCES UsuariosServidor(id_usuario_servidor),
    FOREIGN KEY (id_chat) REFERENCES Chat(id_chat)
);

-- Tabla CanalesServidor
CREATE TABLE CanalesServidor (
    id_canal_servidor UUID PRIMARY KEY,
    nombre VARCHAR(255),
    descripcion VARCHAR(255)
);

-- Tabla ChatMiembrosPublico
CREATE TABLE ChatMiembrosPublico (
    id_chat_miembros UUID PRIMARY KEY,
    id_canal UUID,
    id_chat UUID,
    FOREIGN KEY (id_canal) REFERENCES CanalesServidor(id_canal_servidor),
    FOREIGN KEY (id_chat) REFERENCES Chat(id_chat)
);

-- Tabla CanalMiembros
CREATE TABLE CanalMiembros (
    id_canal_miembro UUID PRIMARY KEY,
    id_usuario_servidor UUID,
    id_canal_servidor UUID,
    FOREIGN KEY (id_usuario_servidor) REFERENCES UsuariosServidor(id_usuario_servidor),
    FOREIGN KEY (id_canal_servidor) REFERENCES CanalesServidor(id_canal_servidor)
);

-- Tabla Invitacion
CREATE TABLE Invitacion (
    id_invitacion UUID PRIMARY KEY,
    id_usuario UUID,
    fecha_envio DATE,
    estado BOOLEAN,
    FOREIGN KEY (id_usuario) REFERENCES UsuariosServidor(id_usuario_servidor)
);

-- Tabla InvitacionCanal
CREATE TABLE InvitacionCanal (
    id_invitacion_canal UUID PRIMARY KEY,
    id_canal_servidor UUID,
    id_invitacion UUID,
    FOREIGN KEY (id_canal_servidor) REFERENCES CanalesServidor(id_canal_servidor),
    FOREIGN KEY (id_invitacion) REFERENCES Invitacion(id_invitacion)
);

-- Tabla Notificacion
CREATE TABLE Notificacion (
    id_notificacion UUID PRIMARY KEY,
    id_invitacion UUID,
    contenido VARCHAR(255),
    FOREIGN KEY (id_invitacion) REFERENCES Invitacion(id_invitacion)
);
