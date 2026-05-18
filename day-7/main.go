package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/nathanaelsanilo/my-app/handler"
	"github.com/nathanaelsanilo/my-app/service"
)

func getListUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "log::getListUsers::ok")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// use chi to get url param
	// without chi we need to use `r.PathValue("user_id")`
	userId := chi.URLParam(r, "user_id")
	fmt.Fprintf(w, "log::getUser::ok::userId::%s", userId)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "log::createUser::ok")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "log::updateUser::ok")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "log::deleteUser::ok")
}

func registerMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}

func userMiddleware(h http.Handler) http.Handler {
	return nil
}

func main() {
	r := chi.NewRouter()
	registerMiddleware(r)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello from Chi!")
	// })

	// r.Get("/users/{user_id}", func(w http.ResponseWriter, r *http.Request) {
	// 	userId := chi.URLParam(r, "user_id") // extract url param
	// 	fmt.Fprintf(w, "User id::%s", userId)
	// })

	// multiple params
	r.Get("/users/{user_id}/friends/{friend_id}/posts", func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "user_id")
		friendId := chi.URLParam(r, "friend_id")
		fmt.Fprintf(w, "userId::%s::friendId::%s", userId, friendId)
	})

	// apply route group
	r.Route("/users", func(r chi.Router) {
		// apply custom middleware
		// only applies under /users path
		r.Use(userMiddleware)

		r.Get("/", getListUsers)
		r.Get("/{user_id}", getUser)
		r.Post("/", createUser)
		r.Put("/{user_id}", updateUser)
		r.Delete("/{user_id}", deleteUser)
	})

	svc := service.NewNoteService()
	hndlr := handler.NewNoteHandler(svc)
	r.Route("/notes", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)

		r.Get("/", hndlr.GetList)
		r.Post("/", hndlr.Create)
		r.Get("/{id}", hndlr.GetDetail)
		r.Put("/{id}", hndlr.Update)
		r.Delete("/{id}", hndlr.Delete)
	})

	// notice r being passed
	// so we dont have to define http.HandleFunc
	// chi takeover
	http.ListenAndServe(":8080", r)
}
