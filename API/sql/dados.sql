insert into usuarios(nome, nick, email, senha) values
("usuario1", "usuario_1", "usuario1@gmail.com", "$2a$10$h11vm/ppaxN1vTkxSromD.dGUP4j5qlSxRQepOHL0IIOapIvA2iXG"),
("usuario2", "usuario_2", "usuario2@gmail.com", "$2a$10$h11vm/ppaxN1vTkxSromD.dGUP4j5qlSxRQepOHL0IIOapIvA2iXG"),
("usuario3", "usuario_3", "usuario3@gmail.com", "$2a$10$h11vm/ppaxN1vTkxSromD.dGUP4j5qlSxRQepOHL0IIOapIvA2iXG");

insert into seguidores(usuario_id, seguidor_id) values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id) values
("Publicação do Usuário 1", "Essa é a publicação do usuário 1!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3!", 3);