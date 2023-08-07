package main

import (
	"fmt"
	"net/http"

	"github.com/Gileno29/gestor-tarefas/config"
	"github.com/Gileno29/gestor-tarefas/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := config.Load()

	if err != nil {
		panic(err)

	}
	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)
	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)
}
