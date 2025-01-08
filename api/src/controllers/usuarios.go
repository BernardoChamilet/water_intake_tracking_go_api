package controllers

import (
	"API/src/config"
	"API/src/database"
	"API/src/models"
	"API/src/repositories"
	"API/src/responses"
	"API/src/security"
	"encoding/json"
	"io"
	"net/http"
)

// CriarUsuario cria um novo usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando
	var usuario models.Usuario
	if erro = json.Unmarshal(corpoReq, &usuario); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Aqui campos também são formatos e senhas criptogradas
	if erro = usuario.Validar(); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para inserir dados no banco de dados
	if erro = repositories.CriarUsuario(&usuario, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusCreated, usuario)
}

// BuscarLogado busca dados de um usuário logado
func BuscarLogado(w http.ResponseWriter, r *http.Request) {
	// Extraindo matricula logado do contexto da requisição
	matriculaLogado := r.Context().Value(config.MatriculaKey).(int)
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para buscar dados do usuário logado
	dados, erro := repositories.BuscarLogado(matriculaLogado, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta
	responses.RespostaDeSucesso(w, http.StatusOK, dados)
}

// AtualizarConta atualiza nome, sobrenome, apelido, sexo e data de nascimento de um usuário
func AtualizarConta(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando dados
	var dadosDaConta models.Usuario
	if erro = json.Unmarshal(corpoReq, &dadosDaConta); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = dadosDaConta.ValidarConta(); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo matricula logado do contexto da requisição
	matriculaLogado := r.Context().Value(config.MatriculaKey).(int)
	dadosDaConta.Matricula = matriculaLogado
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para atualizar dados no banco de dados
	if erro = repositories.AtualizarConta(dadosDaConta, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}

// AtualizarCelular atualiza celular de um usuário
func AtualizarCelular(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando dados
	var celular models.Usuario
	if erro = json.Unmarshal(corpoReq, &celular); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = celular.ValidarCelular(); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo matricula logado do contexto da requisição
	matriculaLogado := r.Context().Value(config.MatriculaKey).(int)
	celular.Matricula = matriculaLogado
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para atualizar dados no banco de dados
	if erro = repositories.AtualizarCelular(celular, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}

// AtualizarEmail atualiza email de um usuário
func AtualizarEmail(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando dados
	var email models.Usuario
	if erro = json.Unmarshal(corpoReq, &email); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = email.ValidarEmail(); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	// Extraindo matricula logado do contexto da requisição
	matriculaLogado := r.Context().Value(config.MatriculaKey).(int)
	email.Matricula = matriculaLogado
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositories para atualizar dados no banco de dados
	if erro = repositories.AtualizarEmail(email, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}

// AtualizarSenha atualiza senha de um usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct e validando dados
	var senhas models.Senhas
	if erro = json.Unmarshal(corpoReq, &senhas); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = senhas.ValidarSenhas(); erro != nil {
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
	// Chamando repositories para buscar senha no banco de dados
	senhaSalva, erro := repositories.BuscarSenhaPorMatricula(matriculaLogado, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Vendo se senha salva é igual a recebida
	if erro = security.VerificarSenha(senhaSalva, senhas.SenhaAtual); erro != nil {
		responses.RespostaDeErro(w, http.StatusUnauthorized, erro)
		return
	}
	// Criptografando senha nova para guardar no banco
	senhaNovaHash, erro := security.GerarSenhaComHash(senhas.SenhaNova)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Chamando repositories para atualizar senha no banco
	if erro = repositories.AtualizarSenha(string(senhaNovaHash), matriculaLogado, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}
