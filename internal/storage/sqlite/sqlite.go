package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	// "github.com/mattn/go-sqlite3"

	"github.com/kolesico/FocusGuard/internal/config"
)

// NewConnetcion - Создание подключения к базе данных SqlLite
func NewConnection(ctx context.Context, cfg *config.Config) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, fmt.Errorf("Failed open db: %s", err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXIST events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event TEXT NOT NULL,
		timestamp TEXT NOT NULL)
	`)
	if err != nil {
		return nil, fmt.Errorf("Failed prepare table: %s", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("Failed exec create table: %s", err)
	}

	return db, nil
}
