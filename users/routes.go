package users

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", LoginHandler)
	r.Post("/register", RegistrationHandler)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verify(GetTokenAuth()))
		r.Use(jwtauth.Authenticator)

		r.Get("/user", UserHandler)
	})

	return r
}
