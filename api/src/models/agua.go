package models

import (
	"errors"
	"time"
)

type ConsumoAgua struct {
	UsuarioMatricula int       `json:"usuario_matricula,omitempty"`
	Data             time.Time `json:"data,omitempty"` //yyyy-mm-ddThh:mm:ssZ
	Quantidade       int       `json:"quantidade,omitempty"`
}

// Validar verifica se o campo data está presente e se a quantidade de água e porcentagem da meta foi maior que 0
func (c ConsumoAgua) Validar() error {
	if c.Data.IsZero() {
		return errors.New("data e hora do consumo faltando")
	}
	if c.Quantidade == 0 {
		return errors.New("a quantidade de agua nao podem ser 0")
	}
	return nil
}
