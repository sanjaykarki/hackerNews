package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (a *application) routes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	if a.debug {
		r.Use(middleware.Logger)
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(a.appName))
	})
	r.Get("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Comments"))
	})

	return r
}
