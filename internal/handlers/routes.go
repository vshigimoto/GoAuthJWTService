package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

const (
	access  = "/auth/token"
	refresh = "/auth/token/refresh"
)

func (h *Handlers) InitRoutes() http.Handler {
	router := chi.NewRouter()
	router.Get(access, h.GetTokens)
	router.Post(refresh, h.RefreshToken)
	//TODO: add swagger endpoint

	h.l.Infof("GET %s endpoint", access)
	h.l.Infof("POST %s endpoint", refresh)

	h.l.Info("Init Routes completed")
	return router
}
