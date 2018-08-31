package users

import (
	"fmt"

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

		fmt.Println(jwtauth.TokenCtxKey)
		fmt.Println(jwtauth.ErrorCtxKey)

		r.Get("/user", UserHandler)
	})

	return r
}
