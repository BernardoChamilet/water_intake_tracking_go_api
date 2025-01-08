package controllers

import (
	"API/src/auth"
	"API/src/database"
	"API/src/models"
	"API/src/repositories"
	"API/src/responses"
	"API/src/security"
	"encoding/json"
	"io"
	"net/http"
)

// Login executa o login de um usuário
func Login(w http.ResponseWriter, r *http.Request) {
	// Lendo corpo da requisição
	corpoReq, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer r.Body.Close()
	// Passando para struct
	var usuario models.Usuario
	if erro = json.Unmarshal(corpoReq, &usuario); erro != nil {
		responses.RespostaDeErro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = usuario.ValidarLogin(); erro != nil {
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
	// Chamando repositories para buscar senha para comparação
	MatriculaESenha, erro := repositories.BuscarMatriculaESenhaPorEmail(usuario.Email, db)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Verificando se senha está correta
	if erro = security.VerificarSenha(MatriculaESenha.Senha, usuario.Senha); erro != nil {
		responses.RespostaDeErro(w, http.StatusUnauthorized, erro)
		return
	}
	// Gerando token
	token, erro := auth.GerarToken(MatriculaESenha.Matricula)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Colocando token na lista branca
	if erro = repositories.GuardarToken(MatriculaESenha.Matricula, token, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	resposta := models.RespostaLogin{Matricula: MatriculaESenha.Matricula, Token: token}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusOK, resposta)
}

// Logout executa o logout de um usuário logado
func Logout(w http.ResponseWriter, r *http.Request) {
	// Pegando token da requisição
	token, erro := auth.ExtrairToken(r)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnauthorized, erro)
		return
	}
	// Pegando matricula do logado
	matricula_logado, erro := auth.ExtrairMatricula(token)
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusUnauthorized, erro)
		return
	}
	// Abrindo conexão com banco de dados
	db, erro := database.ConectarDB()
	if erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	// Chamando repositorios para retirar token da lista branca
	if erro = repositories.DeletarToken(matricula_logado, token, db); erro != nil {
		responses.RespostaDeErro(w, http.StatusInternalServerError, erro)
		return
	}
	// Enviando resposta de sucesso
	responses.RespostaDeSucesso(w, http.StatusNoContent, nil)
}
