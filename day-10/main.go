package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nathanaelsanilo/my-app/handler"
	"github.com/nathanaelsanilo/my-app/services"
	"github.com/nathanaelsanilo/my-app/stores"
)

func main() {

	r := chi.NewRouter()

	userStore := stores.NewUserStore()
	userService := services.NewUserService(userStore)
	userHandler := handler.NewUserHandler(userService)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetAll)
		r.Get("/{user_id}", userHandler.GetDetail)
		r.Post("/", userHandler.Create)
		r.Put("/{user_id}", userHandler.Update)
		r.Delete("/{user_id}", userHandler.Delete)
	})

	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
}
