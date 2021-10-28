package routes

import (
	"boiler/api"
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api/v1/", func(r chi.Router) {

		r.Post("/auth/signup", api.Signup)
		r.Post("/auth/login", api.Login)
		r.Post("/auth/logout", api.Logout)
		r.Get("/profile", Authentication(api.GetProfile))

	})
	return r
}
