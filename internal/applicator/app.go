package applicator

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/vshigimoto/GoAuthJWTService/internal/config"
	"github.com/vshigimoto/GoAuthJWTService/internal/handlers"
	httpserver "github.com/vshigimoto/GoAuthJWTService/internal/infrastructure/http"
	"github.com/vshigimoto/GoAuthJWTService/internal/infrastructure/mongo"
	"github.com/vshigimoto/GoAuthJWTService/internal/repository"
	"github.com/vshigimoto/GoAuthJWTService/internal/services"
	"go.uber.org/zap"
)

type Applicator struct {
	l   *zap.SugaredLogger
	cfg *config.Config
}

func New(l *zap.SugaredLogger, cfg *config.Config) *Applicator {
	return &Applicator{
		l:   l,
		cfg: cfg,
	}
}

func (a *Applicator) Run() {
	mongoClient, err := mongo.Init(a.cfg.Mongo.URL, a.cfg.Mongo.Username, a.cfg.Mongo.Password)
	if err != nil {
		a.l.Panicf("Failed connection to mongo %v", err)
	}

	repo := repository.New(mongoClient.Database(a.cfg.Mongo.DBName), a.cfg.Mongo.Collection)

	services := services.New(repo)

	handlers := handlers.New(services, a.l)

	server := httpserver.New(a.cfg, handlers.InitRoutes())

	server.SetKeepAlivesEnabled(true)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-shutdown
		a.l.Info("Stop server")
		ctx, cancel := context.WithTimeout(context.Background(), a.cfg.HTTP.ShutdownTimeout)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			a.l.Infof("Failure stop server: %v", err)
		}
	}()
	a.l.Info("Init graceful shutdown completed")
	a.l.Infof("Start server on port: %d", a.cfg.HTTP.Port)
	if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Failure start server: %v", err)
	}
}
