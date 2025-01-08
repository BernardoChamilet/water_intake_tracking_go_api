package utils

import (
	"errors"
	"time"
)

// CalcularInicioDaSemana calcula o início de uma semana com base no ano e número da semana (ISO 8601)
func CalcularInicioDaSemana(ano, semana int) (time.Time, error) {
	// Define uma data inicial no ano (1º de janeiro)
	dataInicial := time.Date(ano, 1, 1, 0, 0, 0, 0, time.UTC)

	// Ajusta para a primeira segunda-feira do ano
	for dataInicial.Weekday() != time.Monday {
		dataInicial = dataInicial.AddDate(0, 0, 1)
	}

	// Adiciona semanas para encontrar a semana desejada
	inicioSemana := dataInicial.AddDate(0, 0, (semana-1)*7)

	// Verifica se a semana é válida
	if inicioSemana.Year() != ano {
		return time.Time{}, errors.New("semana com esse numero nao existe nesse ano")
	}
	return inicioSemana, nil
}
