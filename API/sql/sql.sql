CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS usuarios;

DROP TABLE IF EXISTS seguidores;

DROP TABLE IF EXISTS publicaoes;

CREATE TABLE usuarios (
    id int auto_increment PRIMARY KEY,
    nome varchar(50) NOT NULL,
    nick varchar(50) NOT NULL UNIQUE,
    email varchar(60) NOT NULL UNIQUE,
    senha varchar(256) NOT NULL,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE = INNODB;

CREATE TABLE seguidores (
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,


seguidor_id int not null,
FOREIGN KEY (seguidor_id) REFERENCES usuarios (id) ON DELETE CASCADE,

primary key(usuario_id, seguidor_id) ) ENGINE=INNODB;

CREATE TABLE publicacoes (
    id int auto_increment primary key,
    titulo varchar(60) not null,
    conteudo varchar(300) not null,
    autor_id int not null,
    FOREIGN KEY (autor_id) REFERENCES usuarios (id) ON DELETE CASCADE,
    curtidas int default 0,
    criadaEm timestamp default current_timestamp
) ENGINE = INNODB;
