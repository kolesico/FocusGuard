package application

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/kolesico/FocusGuard/internal/config"
	"github.com/kolesico/FocusGuard/internal/logger"
	"github.com/kolesico/FocusGuard/internal/controllers"
	"github.com/kolesico/FocusGuard/internal/storage/sqlite"
)

type App struct {
	cfg *config.Config
	storage *sql.DB
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
	db, err := sqlite.NewConnection(ctx, a.cfg)
	if err != nil {
		a.log.Error("BD error", err)
		return err
	}
	a.storage = db

	// Инициализируем работу с методами БД репо
	sqliteRepo := sqlite.NewSqliteRepository(a.storage)

	a.log.Info("Database connect successfully")

	// Инициализация хэндлеров

	eventHandler := controllers.NewEventsHandler(sqliteRepo, *a.log)

}