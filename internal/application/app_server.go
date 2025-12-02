package application

import (
	"context"
	"log/slog"

	"github.com/kolesico/FocusGuard/internal/config"
	"github.com/kolesico/FocusGuard/internal/database"
	"github.com/kolesico/FocusGuard/internal/logger"
)

type App struct {
	cfg *config.Config
	storage string
	log *slog.Logger
}

func NewApp(configPath string) (*App, error) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	log := logger.InitLogger(cfg.LogLevel)

	log.Info("Success init application")

	return &App{cfg: cfg, log: log}, nil
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Формируем подключение к БД
	a.storage = database.NewConnection(ctx, a.cfg)

	a.log.Info("Database connect successfully")

	// Инициализация хэндлеров

	eventHandler := controllers.NewEventHandler()

}