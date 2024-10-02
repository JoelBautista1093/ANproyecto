create database an_proyect;
use an_proyect;

create table USERS (
    id int not null auto_increment,
    email varchar(255) not null,
    name varchar(255) not null,
    password varchar(255) not null,
    primary key(id)
);

CREATE TABLE ACTIVOS (
    id SERIAL PRIMARY KEY,
    codigo_activo_br int,
    codigo_activo_qr int,
    created_by int not null,
    tipo VARCHAR(50) NOT NULL,        -- Tipo de activo (impresora, computadora, etc.)
    marca VARCHAR(50) NOT NULL,
    modelo VARCHAR(50) NOT NULL,
    numero_serie VARCHAR(50) NOT NULL UNIQUE,
    estado VARCHAR(20) NOT NULL,      -- Estado (nuevo, usado, reparado, etc.)
    fecha_adquisicion DATE NOT NULL,
    primary key (id),
    foreign key (created_by) REFERENCES USERS(id)
);

CREATE TABLE ACTIVOS_RECIBIDOS (
    id SERIAL PRIMARY KEY,
    codigo_activo_br int,
    codigo_activo_qr int,
    activo_id INT REFERENCES Activos(id),
    fecha_recibido DATE NOT NULL,
    recibido_por VARCHAR(100) NOT NULL,
    observaciones TEXT,
    primary key (id)
);

CREATE TABLE ACTIVOS_ENTREGADOS (
    id SERIAL PRIMARY KEY,
    codigo_activo_br int,
    codigo_activo_qr int,
    activo_id INT REFERENCES Activos(id),
    fecha_entregado DATE NOT NULL,
    entregado_a VARCHAR(100) NOT NULL,
    observaciones TEXT,
    primary key (id)
);

CREATE TABLE ACTUALIZACIONES (
    id SERIAL PRIMARY KEY,
    codigo_activo_br int,
    codigo_activo_qr int,
    activo_id INT REFERENCES Activos(id),
    fecha_actualizacion DATE NOT NULL,
    descripcion TEXT NOT NULL,
    estado_anterior VARCHAR(20),
    estado_nuevo VARCHAR(20),
    primary key (id)
);

create table ROLES (
    id int not null auto_increment,
    name varchar(255) not null,
    primary key(id)
);

create table USER_ROLES(
    id int not null auto_increment,
    user_id int not null,
    role_id int not null,
    primary key (id),
    foreign key (user_id) REFERENCES USERS(id)
    foreign key (role_id) REFERENCES ROLES(id)
);

