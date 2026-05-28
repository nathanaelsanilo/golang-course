package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	userStore := NewUserStore()
	userService := NewUserService(userStore)
	userHandler := NewUserHandler(userService)

	r.Route("/users", func(r chi.Router) {
		r.Get("/{user_id}", userHandler.GetById)
		r.Post("/", userHandler.Create)
	})

	const port = ":8080"
	http.ListenAndServe(port, r)
}
