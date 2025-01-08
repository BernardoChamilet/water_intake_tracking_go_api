package repositories

import (
	"API/src/models"
	"database/sql"
	"errors"
)

// CriarUsuario insere um novo usuario no banco de dados
func CriarUsuario(usuario *models.Usuario, db *sql.DB) error {
	sqlStatement := `INSERT INTO usuarios (nome, sobrenome, apelido, celular, email, sexo, data_nascimento, senha) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING matricula`
	if erro := db.QueryRow(sqlStatement, usuario.Nome, usuario.Sobrenome, usuario.Apelido, usuario.Celular, usuario.Email, usuario.Sexo, usuario.DataNascimento, usuario.Senha).Scan(&usuario.Matricula); erro != nil {
		return erro
	}
	return nil
}

// BuscarMatriculaSenhaPorEmail usa um email para buscar matricula e senha de um usuário
func BuscarMatriculaESenhaPorEmail(email string, db *sql.DB) (models.Usuario, error) {
	sqlStatement := `SELECT matricula, senha FROM usuarios WHERE email=$1`
	var usuario models.Usuario
	if erro := db.QueryRow(sqlStatement, email).Scan(&usuario.Matricula, &usuario.Senha); erro != nil {
		if erro == sql.ErrNoRows {
			return models.Usuario{}, errors.New("usuario com esse email nao encontrado")
		}
		return models.Usuario{}, erro
	}
	return usuario, nil
}

// BuscarLogado busca dados exceto a senha de um usuário pela matrícula
func BuscarLogado(matricula int, db *sql.DB) (models.Usuario, error) {
	sqlStatement := `SELECT matricula, nome, sobrenome, apelido, celular, email, sexo, data_nascimento, data_criacao FROM usuarios WHERE matricula=$1`
	var usuario models.Usuario
	if erro := db.QueryRow(sqlStatement, matricula).Scan(&usuario.Matricula, &usuario.Nome, &usuario.Sobrenome, &usuario.Apelido, &usuario.Celular, &usuario.Email, &usuario.Sexo, &usuario.DataNascimento, &usuario.DataCriacao); erro != nil {
		if erro == sql.ErrNoRows {
			return models.Usuario{}, errors.New("matricula nao encontrada")
		}
		return models.Usuario{}, erro
	}
	return usuario, nil
}

// AtualizarConta atualiza nome, sobrenome, apelido, sexo e data de nascimento na tabela usuários
func AtualizarConta(dados models.Usuario, db *sql.DB) error {
	sqlStatement := `UPDATE usuarios SET nome=$1, sobrenome=$2, apelido=$3, sexo=$4, data_nascimento=$5 WHERE matricula=$6`
	result, erro := db.Exec(sqlStatement, dados.Nome, dados.Sobrenome, dados.Apelido, dados.Sexo, dados.DataNascimento, dados.Matricula)
	if erro != nil {
		return erro
	}
	// Verifica se alguma linha foi atualizada
	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro // Retorna erro se não foi possível verificar as linhas afetadas
	}
	if rowsAffected == 0 {
		return errors.New("usuario nao encontrado para atualizar dados")
	}
	return nil
}

// AtualizarCelular atualiza celular na tabela usuários
func AtualizarCelular(dados models.Usuario, db *sql.DB) error {
	sqlStatement := `UPDATE usuarios SET celular=$1 WHERE matricula=$2`
	result, erro := db.Exec(sqlStatement, dados.Celular, dados.Matricula)
	if erro != nil {
		return erro
	}
	// Verifica se alguma linha foi atualizada
	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro // Retorna erro se não foi possível verificar as linhas afetadas
	}
	if rowsAffected == 0 {
		return errors.New("usuario nao encontrado para atualizar dados")
	}
	return nil
}

// AtualizarEmail atualiza email na tabela usuários
func AtualizarEmail(dados models.Usuario, db *sql.DB) error {
	sqlStatement := `UPDATE usuarios SET email=$1 WHERE matricula=$2`
	result, erro := db.Exec(sqlStatement, dados.Email, dados.Matricula)
	if erro != nil {
		return erro
	}
	// Verifica se alguma linha foi atualizada
	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro // Retorna erro se não foi possível verificar as linhas afetadas
	}
	if rowsAffected == 0 {
		return errors.New("usuario nao encontrado para atualizar dados")
	}
	return nil
}

// BuscarSenhaPorMatricula usa matricula para buscar senha de um usuário
func BuscarSenhaPorMatricula(matricula int, db *sql.DB) (string, error) {
	sqlStatement := `SELECT senha FROM usuarios WHERE matricula=$1`
	var senhaSalva string
	if erro := db.QueryRow(sqlStatement, matricula).Scan(&senhaSalva); erro != nil {
		if erro == sql.ErrNoRows {
			return "", errors.New("usuario com essa matricula nao encontrado")
		}
		return "", erro
	}
	return senhaSalva, nil
}

// AtualizarSenha atualiza senha na tabela usuários
func AtualizarSenha(senha string, matricula int, db *sql.DB) error {
	sqlStatement := `UPDATE usuarios SET senha=$1 WHERE matricula=$2`
	result, erro := db.Exec(sqlStatement, senha, matricula)
	if erro != nil {
		return erro
	}
	// Verifica se alguma linha foi atualizada
	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro // Retorna erro se não foi possível verificar as linhas afetadas
	}
	if rowsAffected == 0 {
		return errors.New("usuario nao encontrado para atualizar dados")
	}
	return nil
}
