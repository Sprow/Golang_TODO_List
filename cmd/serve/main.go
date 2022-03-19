package main

import (
	"fmt"
	"github.com/Sprow/todo/cmd/serve/handler"
	"github.com/Sprow/todo/internal/todo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	todoListManager := todo.NewManager(todo.List{
		Title: "todo list 1",
		Items: []todo.Item{{
			Done: false,
			Text: "learn English",
		}, {
			Done: true,
			Text: " learn French",
		}},
	})
	h := handler.NewHandler(todoListManager)

	h.Register(r)

	err := http.ListenAndServe("0.0.0.0:8081", r)
	if err != nil {
		fmt.Println(err)
	}
}
