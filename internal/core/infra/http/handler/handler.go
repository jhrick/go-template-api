package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	*chi.Mux
}

func New() *Handler {
	mux := chi.NewMux()
	mux.Use(middleware.Logger)

	return &Handler{
		mux,
	}
}
