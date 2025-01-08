CREATE TABLE IF NOT EXISTS usuarios (
    matricula SERIAL PRIMARY KEY,
    nome VARCHAR(30) NOT NULL,
    sobrenome VARCHAR(50) NOT NULL,
    apelido VARCHAR(30) NOT NULL,
    celular CHAR(11) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    sexo CHAR(1) NOT NULL,
    data_nascimento DATE NOT NULL,
    senha VARCHAR(128) NOT NULL,
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE IF NOT EXISTS lista_branca (
    usuario_matricula INT NOT NULL,
    token VARCHAR(255) NOT NULL,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (usuario_matricula, token),
    FOREIGN KEY (usuario_matricula) REFERENCES usuarios(matricula) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS historico_de_agua (
    usuario_matricula INT NOT NULL,
    data_consumo TIMESTAMP NOT NULL,
    quantidade INT NOT NULL,
    PRIMARY KEY (usuario_matricula, data_consumo),
    FOREIGN KEY (usuario_matricula) REFERENCES usuarios(matricula) ON DELETE CASCADE
);