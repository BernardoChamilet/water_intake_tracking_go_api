package routes

import (
	"github.com/go-chi/chi"
)

// Rotear adciona as rotas da api ao roteador
func Rotear() chi.Router {
	r := chi.NewRouter()

	// /login e /logout

	r.Mount("/", SessaoRouter())

	// /usuarios

	r.Mount("/usuarios", UsuariosRouter())

	// /agua

	r.Mount("/agua", AguaRouter())

	return r
}
