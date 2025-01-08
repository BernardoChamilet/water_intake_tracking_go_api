package models

import (
	"errors"
	"strings"
)

type Senhas struct {
	SenhaAtual string `json:"senha_atual,omitempty"`
	SenhaNova  string `json:"senha_nova,omitempty"`
}

// ValidarSenhas valida tamanho das senhas
func (u *Senhas) ValidarSenhas() error {
	u.SenhaAtual = strings.TrimSpace(u.SenhaAtual)
	if len(u.SenhaAtual) < 2 {
		return errors.New("senha atual incorreta")
	}
	u.SenhaNova = strings.TrimSpace(u.SenhaNova)
	if len(u.SenhaNova) < 2 {
		return errors.New("senha nova deve ter pelo menos 2 caracteres")
	}
	return nil
}
