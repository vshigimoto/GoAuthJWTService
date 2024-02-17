package server

import (
	"context"
	"github.com/vshigimoto/GoAuthJWTService/internal/config"
	"net/http"
)

type Server struct {
	http *http.Server
}

func New(cfg *config.Config, router http.Handler) *Server {
	return &Server{
		http: &http.Server{
			Addr:           cfg.Http.Address,
			Handler:        router,
			ReadTimeout:    cfg.Http.ReadTimeout,
			WriteTimeout:   cfg.Http.WriteTimeout,
			MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		},
	}
}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}

func (s *Server) SetKeepAlivesEnabled(v bool) {
	s.http.SetKeepAlivesEnabled(v)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
