CREATE DATABASE IF NOT EXISTS devbook;
use devbook;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS seguidores;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
);

CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id) REFERENCES USUARIOS(ID)
    ON DELETE CASCADE,

    seguidor_id int not null,
    FOREIGN KEY (seguidor_id) REFERENCES USUARIOS(ID)
    ON DELETE CASCADE,

    PRIMARY KEY(usuario_id, seguidor_id)
);