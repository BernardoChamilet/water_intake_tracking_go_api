package repositories

import (
	"database/sql"
	"errors"
)

// GuardarToken coloca o token na lista branca
func GuardarToken(matricula int, token string, db *sql.DB) error {
	sqlStatement := `INSERT INTO lista_branca (usuario_matricula, token) VALUES ($1, $2)`
	_, erro := db.Exec(sqlStatement, matricula, token)
	if erro != nil {
		return erro
	}
	return nil
}

// DeletarToken remove um token da lista branca
func DeletarToken(matricula int, token string, db *sql.DB) error {
	sqlStatement := `DELETE FROM lista_branca WHERE usuario_matricula=$1 and token=$2`
	result, erro := db.Exec(sqlStatement, matricula, token)
	if erro != nil {
		return erro
	}
	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro // Retorna erro se não foi possível verificar as linhas afetadas
	}
	if rowsAffected == 0 {
		return errors.New("nenhum registro encontrado para essa matricula e token")
	}
	return nil
}

// BuscarToken verifica se um token está na lista branca
func BuscarToken(token string, db *sql.DB) (int, error) {
	sqlStatement := `SELECT usuario_matricula FROM lista_branca WHERE token=$1`
	var matricula int
	if erro := db.QueryRow(sqlStatement, token).Scan(&matricula); erro != nil {
		if erro == sql.ErrNoRows {
			return 0, errors.New("token nao consta na lista branca")
		}
		return 0, erro
	}
	return matricula, nil
}
