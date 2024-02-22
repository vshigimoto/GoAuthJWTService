package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"

	"github.com/vshigimoto/GoAuthJWTService/internal/applicator"
	"github.com/vshigimoto/GoAuthJWTService/internal/config"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			return
		}
	}(logger)

	l := logger.Sugar()
	l = l.With(zap.String("app", "medods"))

	var cfg config.Config
	err := cleanenv.ReadConfig("config/auth.yaml", &cfg)
	if err != nil {
		l.Fatalf("Failure read config err: %v", err)
		return
	}

	app := applicator.New(l, &cfg)
	app.Run()
}
