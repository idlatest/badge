package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
	"github.com/idlatest/badge/users"
)

var Upgrader websocket.Upgrader

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)

	Cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Content-Type", "X-Auth-Token", "*"},
		Debug:            false,
	})

	r.Use(Cors.Handler)

	r.Mount("/auth", users.Routes())

	return r

}
func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := Upgrader.Upgrade(w, r, nil)

		go func(conn *websocket.Conn) {
			for {
				messageType, msg, _ := conn.ReadMessage()

				conn.WriteMessage(messageType, msg)
			}
		}(conn)
	})

	http.ListenAndServe(":3000", nil)
}
