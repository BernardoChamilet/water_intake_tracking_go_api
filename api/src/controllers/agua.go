package controllers

import (
	"API/src/config"
	"API/src/database"
	"API/src/models"
	"API/src/repositories"
	"API/src/responses"
	"API/src/utils"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

// CriarConsumoAgua registra um consumo de água do usuário logado
func CriarConsumoAgua(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando
	var consumo models.ConsumoAgua
	if erro = json.Unmarshal(corpoReq, &consumo); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = consumo.Validar(); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo matricula logado do contexto da requisição
	matriculaLogado := r.Context().Value(config.MatriculaKey).(int)
	consumo.UsuarioMatricula = matriculaLogado
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para inserir dados no banco de dados
	if erro = repositories.CriarConsumoAgua(consumo, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusCreated, consumo)
}

// BuscarConsumoAgua busca dados de um consumo de água do usuário logado
func BuscarConsumoAgua(w http.ResponseWriter, r *http.Request) {
	// Pegando parâremtros da url
	parametro := chi.URLParam(r, "timestamp")
	timestamp, erro := time.Parse(time.RFC3339, parametro)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo matricula logado do contexto da requisição
	matriculaLogado := r.Context().Value(config.MatriculaKey).(int)
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para bucar dados no banco de dados
	consumo, erro := repositories.BuscarConsumoAgua(matriculaLogado, timestamp, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusOK, consumo)
}

// AtualizarConsumoAgua atualiza dados de um consumo de água do usuário logado
func AtualizarConsumoAgua(w http.ResponseWriter, r *http.Request) {
	// Pegando parâremtros da url
	parametro := chi.URLParam(r, "timestamp")
	timestamp, erro := time.Parse(time.RFC3339, parametro)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando
	var consumo models.ConsumoAgua
	if erro = json.Unmarshal(corpoReq, &consumo); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = consumo.Validar(); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo matricula logado do contexto da requisição
	matriculaLogado := r.Context().Value(config.MatriculaKey).(int)
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para atualizar dados adcionais no banco de dados
	if erro = repositories.AtualizarConsumoAgua(matriculaLogado, timestamp, consumo, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}

// DeletarConsumoAgua deleta um consumo de água do usuário logado
func DeletarConsumoAgua(w http.ResponseWriter, r *http.Request) {
	// Pegando parâremtros da url
	parametro := chi.URLParam(r, "timestamp")
	timestamp, erro := time.Parse(time.RFC3339, parametro)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo matricula logado do contexto da requisição
	matriculaLogado := r.Context().Value(config.MatriculaKey).(int)
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para bucar dados no banco de dados
	if erro = repositories.DeletarConsumoAgua(matriculaLogado, timestamp, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}

// BuscarConsumoAguaDia busca todos consumos de água de um dia do usuário logado
func BuscarConsumoAguaDia(w http.ResponseWriter, r *http.Request) {
	// Pegando parâremtros da url
	dia := chi.URLParam(r, "dia")
	_, erro := time.Parse("2006-01-02", dia)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo matricula logado do contexto da requisição
	matriculaLogado := r.Context().Value(config.MatriculaKey).(int)
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para bucar dados no banco de dados
	consumosDoDia, erro := repositories.BuscarConsumoAguaDia(matriculaLogado, dia, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Caso nenhum registro seja encontrado
	if len(consumosDoDia) == 0 {
		responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusOK, consumosDoDia)
}

// BuscarConsumoAguaMes busca todos consumos de água de um mes do usuário logado
func BuscarConsumoAguaMes(w http.ResponseWriter, r *http.Request) {
	// Pegando parâremtros da url
	parametro := chi.URLParam(r, "mes")
	mes, erro := time.Parse("2006-01", parametro)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo matricula logado do contexto da requisição
	matriculaLogado := r.Context().Value(config.MatriculaKey).(int)
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para bucar dados no banco de dados
	consumosDoMes, erro := repositories.BuscarConsumoAguaMes(matriculaLogado, mes, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Caso nenhum registro seja encontrado
	if len(consumosDoMes) == 0 {
		responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusOK, consumosDoMes)
}

// BuscarConsumoAguaSemana busca todos consumos de água de uma semana do usuário logado
func BuscarConsumoAguaSemana(w http.ResponseWriter, r *http.Request) {
	// Pegando parâremtros da url
	anoStr := chi.URLParam(r, "ano")
	// Validando parâmetros
	ano, erro := strconv.Atoi(anoStr)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	semanaStr := chi.URLParam(r, "semana")
	semana, erro := strconv.Atoi(semanaStr)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Calculando o início da semana (segunda-feira)
	inicioSemana, erro := utils.CalcularInicioDaSemana(ano, semana)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo matricula logado do contexto da requisição
	matriculaLogado := r.Context().Value(config.MatriculaKey).(int)
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para bucar dados no banco de dados
	consumosDaSemana, erro := repositories.BuscarConsumoAguaSemana(matriculaLogado, inicioSemana, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Caso nenhum registro seja encontrado
	if len(consumosDaSemana) == 0 {
		responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusOK, consumosDaSemana)
}
