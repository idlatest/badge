package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/idlatest/badge/users"
)

func main() {
	r := chi.NewRouter()

	r.Mount("/auth", users.Routes())

	http.ListenAndServe(":3000", r)
}
