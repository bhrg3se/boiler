package routes

import (
	"boiler/api/auth"
	"boiler/api/chat"
	"boiler/api/ping"
	"boiler/api/users"
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api/v1/", func(r chi.Router) {
		r.Get("/ping", Authentication(ping.Ping))
		r.Get("/premium", Authentication(PremiumUser(users.PremiumContent)))

		r.Get("/auth/login", Authentication(auth.Login))

		r.Get("/users", Authentication(PremiumUser(users.PremiumContent)))

		r.Get("/ws/", AuthenticationWS(chat.Join))

	})
	return r
}
