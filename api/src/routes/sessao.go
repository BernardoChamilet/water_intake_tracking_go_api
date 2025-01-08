package routes

import (
	"API/src/controllers"
	"API/src/middlewares"

	"github.com/go-chi/chi"
)

// SessaoRouter retorna o roteador de rotas /login e /logout
func SessaoRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", controllers.Login)

	r.With(middlewares.Autenticar).Delete("/logout", controllers.Logout)

	return r
}
