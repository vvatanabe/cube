package worker

import "github.com/go-chi/chi/v5"

type ErrResponse struct {
	HTTPStatusCode int
	Message        string
}

type Api struct {
	Address string
	Port    int
	Worker  *Worker
	Router  *chi.Mux
}
