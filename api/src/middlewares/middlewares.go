package middlewares

import (
	"API/src/auth"
	"API/src/config"
	"API/src/database"
	"API/src/repositories"
	"API/src/responses"
	"context"
	"net/http"
)

// Autenticar verifica se exsite um token no cabeçalho da req e se ele é válido
func Autenticar(proximaFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//vendo se o token é válido
		token, erro := auth.ExtrairToken(r)
		if erro != nil {
			responses.RespostaDeErro(w, http.StatusUnauthorized, erro)
			return
		}
		if erro = auth.ValidarToken(token); erro != nil {
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
		// Vendo se token consta na lista branca
		_, erro = repositories.BuscarToken(token, db)
		if erro != nil {
			responses.RespostaDeErro(w, http.StatusUnauthorized, erro)
			return
		}
		// Buscando matrícula do usuário logado do token
		matriculaLogado, erro := auth.ExtrairMatricula(token)
		if erro != nil {
			responses.RespostaDeErro(w, http.StatusUnauthorized, erro)
			return
		}
		// Salvando matricula no contexto da requisição
		ctx := context.WithValue(r.Context(), config.MatriculaKey, matriculaLogado)
		proximaFunc.ServeHTTP(w, r.WithContext(ctx))
	})
}
