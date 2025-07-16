insert into usuarios(nome, nick, email, senha) values
("usuario1", "usuario_1", "usuario1@gmail.com", "$2a$10$h11vm/ppaxN1vTkxSromD.dGUP4j5qlSxRQepOHL0IIOapIvA2iXG"),
("usuario2", "usuario_2", "usuario2@gmail.com", "$2a$10$h11vm/ppaxN1vTkxSromD.dGUP4j5qlSxRQepOHL0IIOapIvA2iXG"),
("usuario3", "usuario_3", "usuario3@gmail.com", "$2a$10$h11vm/ppaxN1vTkxSromD.dGUP4j5qlSxRQepOHL0IIOapIvA2iXG")

insert into seguidores(usuario_id, seguidor_id) values
(1, 2),
(3, 1),
(1, 3);
