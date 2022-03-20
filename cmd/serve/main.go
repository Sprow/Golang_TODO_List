package main

import (
	"fmt"
	"github.com/Sprow/todo/cmd/serve/handler"
	"github.com/Sprow/todo/internal/session"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	s := session.NewManagerSession()
	h := handler.NewHandler(s)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	h.Register(r)
	err := http.ListenAndServe("0.0.0.0:8081", r)
	if err != nil {
		fmt.Println(err)
	}
}
