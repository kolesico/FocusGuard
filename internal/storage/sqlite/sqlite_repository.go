package sqlite

import (
	"context"
	"database/sql"

	"github.com/kolesico/FocusGuard/internal/events"
)

type SqlLiteRepo struct {
	db *sql.DB
}

func NewSqliteRepository(db *sql.DB) *SqlLiteRepo {
	return &SqlLiteRepo{db: db}
}

// SaveEvent Сохраняем полученное событие
func (s *SqlLiteRepo) SaveEvent(ctx context.Context, event *events.Events) (int64, error) {
	stmt, err := s.db.Prepare(`
		INSERT INTO events(event, timestamp) VALUES(?, ?)
	`)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(event.Event, event.Timestamp)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
