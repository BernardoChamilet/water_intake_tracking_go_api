package routes

import (
	"API/src/controllers"
	"API/src/middlewares"

	"github.com/go-chi/chi"
)

// AguaRouter retorna roteador de rotas /agua
func AguaRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middlewares.Autenticar)

	r.Post("/", controllers.CriarConsumoAgua)

	r.Get("/{timestamp}", controllers.BuscarConsumoAgua)

	r.Put("/{timestamp}", controllers.AtualizarConsumoAgua)

	r.Delete("/{timestamp}", controllers.DeletarConsumoAgua)

	r.Get("/dia/{dia}", controllers.BuscarConsumoAguaDia)

	r.Get("/mes/{mes}", controllers.BuscarConsumoAguaMes)

	r.Get("/semana/{ano}/{semana}", controllers.BuscarConsumoAguaSemana)

	return r
}
