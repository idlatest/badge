package users

import (
	"github.com/go-chi/chi"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", LoginHandler)
	r.Post("/register", RegistrationHandler)

	return r
}
