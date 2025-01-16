INSERT INTO usuarios (nome, nick, email, senha)  
VALUES  
('Ana Costa', 'aninha', 'ana.costa@example.com', '$2a$10$IDPOrmdZpR4cA5/K7hSY9.uKo1JOVCNSGUddnvByJCcEecM5uUOiC'),  
('Carlos Silva', 'carlito', 'carlos.silva@example.com', '$2a$10$IDPOrmdZpR4cA5/K7hSY9.uKo1JOVCNSGUddnvByJCcEecM5uUOiC'),  
('Bruna Oliveira', 'bruolive', 'bruna.oliveira@example.com', '$2a$10$IDPOrmdZpR4cA5/K7hSY9.uKo1JOVCNSGUddnvByJCcEecM5uUOiC'),  
('Daniel Santos', 'dan_santos', 'daniel.santos@example.com', '$2a$10$IDPOrmdZpR4cA5/K7hSY9.uKo1JOVCNSGUddnvByJCcEecM5uUOiC'),  
('Fernanda Lima', 'fefa', 'fernanda.lima@example.com', '$2a$10$IDPOrmdZpR4cA5/K7hSY9.uKo1JOVCNSGUddnvByJCcEecM5uUOiC');

INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES
(1, 5),
(2, 3),
(3, 4),
(4, 2),
(5, 1);

INSERT INTO publicacoes(titulo, conteudo, autor_id)
VALUES
("publicação do usuario 1", "publicação de teste", 1),
("publicação do usuario 1", "publicação de teste", 2),
("publicação do usuario 1", "publicação de teste", 3);