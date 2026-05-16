package main

import (
	"fmt"
	"net/http"

	handler "github.com/nathanaelsanilo/my-app/handler"
	service "github.com/nathanaelsanilo/my-app/service"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!\n%s", readRequest(r))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Resource not found!")
}

func readRequest(r *http.Request) string {
	method := r.Method
	path := r.URL.Path
	name := r.URL.Query().Get("name")

	return fmt.Sprintf("Request---Method::%s---Path::%s---Name::%s", method, path, name)
}

func main() {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	// newer version 1.22+. supported by mux
	// order doesn't matter. more specific win
	http.HandleFunc("GET /users", userHandler.GetListUsersHandler)
	http.HandleFunc("GET /users/{user_id}", userHandler.GetUserByIdHandler)
	http.HandleFunc("POST /users", userHandler.CreateUserHandler)
	http.HandleFunc("PUT /users/{user_id}", userHandler.UpdateUserHandler)
	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":8080", nil)
}
