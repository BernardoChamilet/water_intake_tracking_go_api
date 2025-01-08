package repositories

import (
	"API/src/models"
	"database/sql"
	"errors"
	"time"
)

// CriarConsumoAgua insere novo consumo no histórico de água
func CriarConsumoAgua(consumo models.ConsumoAgua, db *sql.DB) error {
	sqlStatement := `INSERT INTO historico_de_agua (usuario_matricula, data_consumo, quantidade) VALUES ($1, $2, $3)`
	_, erro := db.Exec(sqlStatement, consumo.UsuarioMatricula, consumo.Data, consumo.Quantidade)
	if erro != nil {
		return erro
	}
	return nil
}

// BuscarConsumoAgua busca um consumo de água do histórico de água
func BuscarConsumoAgua(matricula int, timestamp time.Time, db *sql.DB) (models.ConsumoAgua, error) {
	sqlStatement := `SELECT usuario_matricula, data_consumo, quantidade FROM historico_de_agua WHERE usuario_matricula=$1 AND data_consumo=$2`
	var consumo models.ConsumoAgua
	if erro := db.QueryRow(sqlStatement, matricula, timestamp).Scan(&consumo.UsuarioMatricula, &consumo.Data, &consumo.Quantidade); erro != nil {
		if erro == sql.ErrNoRows {
			return models.ConsumoAgua{}, errors.New("usuario logado nao consumiu agua nesse timestamp")
		}
		return models.ConsumoAgua{}, erro
	}
	return consumo, nil
}

// AtualizarConsumoAgua atualiza dados de um consumo de água no histórico de água
func AtualizarConsumoAgua(matricula int, timestamp time.Time, consumo models.ConsumoAgua, db *sql.DB) error {
	sqlStatement := `UPDATE historico_de_agua SET data_consumo=$1, quantidade=$2 WHERE usuario_matricula=$4 AND data_consumo=$5`
	result, erro := db.Exec(sqlStatement, consumo.Data, consumo.Quantidade, matricula, timestamp)
	if erro != nil {
		return erro
	}
	// Verifica se alguma linha foi atualizada
	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro // Retorna erro se não foi possível verificar as linhas afetadas
	}
	if rowsAffected == 0 {
		return errors.New("consumo de agua nao encontrado para atualizar dados")
	}
	return nil
}

// DeletarConsumaAgua deleta um consumo de água do histórico de água
func DeletarConsumoAgua(matricula int, timestamp time.Time, db *sql.DB) error {
	sqlStatement := `DELETE FROM historico_de_agua WHERE usuario_matricula=$1 AND data_consumo=$2`
	result, erro := db.Exec(sqlStatement, matricula, timestamp)
	if erro != nil {
		return erro
	}
	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro // Retorna erro se não foi possível verificar as linhas afetadas
	}
	if rowsAffected == 0 {
		return errors.New("usuario logado nao tem nenhum consumo de agua nesse timestamp")
	}
	return nil
}

// BuscarConsumoAguaDia busca todo consumo de água de um dia
func BuscarConsumoAguaDia(matricula int, dia string, db *sql.DB) ([]models.ConsumoAgua, error) {
	sqlStatement := `SELECT usuario_matricula, data_consumo, quantidade FROM historico_de_agua WHERE usuario_matricula = $1 AND data_consumo >= $2 AND data_consumo < $3`
	rows, err := db.Query(sqlStatement, matricula, dia+"T00:00:00Z", dia+"T23:59:59Z")
	if err != nil {
		return []models.ConsumoAgua{}, err
	}
	defer rows.Close()
	var consumosDoDia []models.ConsumoAgua
	// Itera sobre as linhas retornadas
	for rows.Next() {
		var consumo models.ConsumoAgua
		if err := rows.Scan(&consumo.UsuarioMatricula, &consumo.Data, &consumo.Quantidade); err != nil {
			return []models.ConsumoAgua{}, err
		}
		consumosDoDia = append(consumosDoDia, consumo)
	}

	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return []models.ConsumoAgua{}, err
	}
	return consumosDoDia, nil
}

// BuscarConsumoAguaMes busca todo consumo de água de um mês
func BuscarConsumoAguaMes(matricula int, mes time.Time, db *sql.DB) ([]models.ConsumoAgua, error) {
	inicioDoProximoMes := mes.AddDate(0, 1, 0)
	sqlStatement := `SELECT usuario_matricula, data_consumo, quantidade FROM historico_de_agua WHERE usuario_matricula = $1 AND data_consumo >= $2 AND data_consumo < $3`
	rows, err := db.Query(sqlStatement, matricula, mes.Format(time.RFC3339), inicioDoProximoMes.Format(time.RFC3339))
	if err != nil {
		return []models.ConsumoAgua{}, err
	}
	defer rows.Close()
	var consumosDoMes []models.ConsumoAgua
	// Itera sobre as linhas retornadas
	for rows.Next() {
		var consumo models.ConsumoAgua
		if err := rows.Scan(&consumo.UsuarioMatricula, &consumo.Data, &consumo.Quantidade); err != nil {
			return []models.ConsumoAgua{}, err
		}
		consumosDoMes = append(consumosDoMes, consumo)
	}

	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return []models.ConsumoAgua{}, err
	}
	return consumosDoMes, nil
}

// BuscarConsumoAguaSemana busca todo consumo de água de uma semana
func BuscarConsumoAguaSemana(matricula int, inicioSemana time.Time, db *sql.DB) ([]models.ConsumoAgua, error) {
	// Calculando fim da semana
	fimSemana := inicioSemana.AddDate(0, 0, 7) // Adiciona 7 dias

	// Fazendo consulta
	sqlStatement := `SELECT usuario_matricula, data_consumo, quantidade FROM historico_de_agua WHERE usuario_matricula = $1 AND data_consumo >= $2 AND data_consumo < $3`
	rows, err := db.Query(sqlStatement, matricula, inicioSemana.Format(time.RFC3339), fimSemana.Format(time.RFC3339))
	if err != nil {
		return []models.ConsumoAgua{}, err
	}
	defer rows.Close()
	var consumosDaSemana []models.ConsumoAgua
	// Itera sobre as linhas retornadas
	for rows.Next() {
		var consumo models.ConsumoAgua
		if err := rows.Scan(&consumo.UsuarioMatricula, &consumo.Data, &consumo.Quantidade); err != nil {
			return []models.ConsumoAgua{}, err
		}
		consumosDaSemana = append(consumosDaSemana, consumo)
	}

	// Verifica se ocorreu algum erro durante a iteração
	if err = rows.Err(); err != nil {
		return []models.ConsumoAgua{}, err
	}
	return consumosDaSemana, nil
}
