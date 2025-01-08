package models

import (
	"API/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	Matricula      int    `json:"matricula,omitempty"`
	Nome           string `json:"nome,omitempty"`
	Sobrenome      string `json:"sobrenome,omitempty"`
	Apelido        string `json:"apelido,omitempty"`
	Celular        string `json:"celular,omitempty"`
	Email          string `json:"email,omitempty"`
	Sexo           string `json:"sexo,omitempty"`
	DataNascimento string `json:"data_nascimento,omitempty"`
	Senha          string `json:"senha,omitempty"`
	DataCriacao    string `json:"data_criacao,omitempty"`
}

// Validar valida formato e tamanho dos dados, remove espaços em branco e criptografa a senha
func (u *Usuario) Validar() error {
	u.Nome = strings.TrimSpace(u.Nome)
	if len(u.Nome) < 2 {
		return errors.New("nome deve ter pelo menos 2 caracteres")
	}
	u.Sobrenome = strings.TrimSpace(u.Sobrenome)
	if len(u.Sobrenome) < 2 {
		return errors.New("sobrenome deve ter pelo menos 2 caracteres")
	}
	u.Apelido = strings.TrimSpace(u.Apelido)
	if len(u.Apelido) < 2 {
		return errors.New("apelido deve ter pelo menos 2 caracteres")
	}
	u.Sexo = strings.TrimSpace(u.Sexo)
	if len(u.Sexo) != 1 {
		return errors.New("sexo deve ter 1 caracter (M/F)")
	}
	u.Celular = strings.TrimSpace(u.Celular)
	if len(u.Celular) != 11 {
		return errors.New("celular no formato inválido")
	}
	_, erro := time.Parse("2006-01-02", u.DataNascimento)
	if erro != nil {
		return errors.New("data de nascimento inválida, formato esperado: yyyy-mm-dd")
	}
	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("email inváido")
	}
	u.Senha = strings.TrimSpace(u.Senha)
	if len(u.Senha) < 2 {
		return errors.New("senha deve ter pelo menos 2 caracteres")
	}
	senhaHash, erro := security.GerarSenhaComHash(u.Senha)
	if erro != nil {
		return erro
	}
	u.Senha = string(senhaHash)
	return nil
}

// ValidarConta valida formato do nome, sobrenome, apelido, sexo e data de nascimento
func (u *Usuario) ValidarConta() error {
	u.Nome = strings.TrimSpace(u.Nome)
	if len(u.Nome) < 2 {
		return errors.New("nome deve ter pelo menos 2 caracteres")
	}
	u.Sobrenome = strings.TrimSpace(u.Sobrenome)
	if len(u.Sobrenome) < 2 {
		return errors.New("sobrenome deve ter pelo menos 2 caracteres")
	}
	u.Apelido = strings.TrimSpace(u.Apelido)
	if len(u.Apelido) < 2 {
		return errors.New("apelido deve ter pelo menos 2 caracteres")
	}
	u.Sexo = strings.TrimSpace(u.Sexo)
	if len(u.Sexo) != 1 {
		return errors.New("sexo deve ter 1 caracter (M/F)")
	}
	_, erro := time.Parse("2006-01-02", u.DataNascimento)
	if erro != nil {
		return errors.New("data de nascimento invalida, formato esperado: yyyy-mm-dd")
	}
	return nil
}

// ValidarCelular valida celular
func (u *Usuario) ValidarCelular() error {
	u.Celular = strings.TrimSpace(u.Celular)
	if len(u.Celular) != 11 {
		return errors.New("celular no formato invalido")
	}
	return nil
}

// ValidarEmail valida email
func (u *Usuario) ValidarEmail() error {
	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("email invalido")
	}
	return nil
}

// ValidarLogin verifica se dados de login estão presentes
func (u *Usuario) ValidarLogin() error {
	if u.Email == "" || u.Senha == "" {
		return errors.New("campos faltando")
	}
	return nil
}
