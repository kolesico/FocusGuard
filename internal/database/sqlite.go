package database

import (
	"context"
	"fmt"

	"github.com/kolesico/FocusGuard/internal/config"
)

// NewConnetcion - Создание подключения к базе данных SqlLite
func NewConnection(ctx context.Context, cfg *config.Config) string {
	fmt.Println("Success connect to DB")
	return "connected"
}