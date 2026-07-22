package db

import (
	"context"
	"database/sql"
	"embed"

	_ "github.com/mattn/go-sqlite3"
	goose "github.com/pressly/goose/v3"
	"github.com/tobiaskrok/plop/internal/types"
)

//go:embed migrations/*.sql
var migrations embed.FS

type Sqlite struct {
	*sql.DB
}

func NewSqlite(dsn string) (DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	return &Sqlite{db}, nil
}

func (s *Sqlite) Migrate(ctx context.Context) error {
	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}
	return goose.UpContext(ctx, s.DB, "migrations")
}

func (s *Sqlite) CreateSession(ctx context.Context, session *types.Session) error {

	return nil
}

func (s *Sqlite) GetMessagesBySession(ctx context.Context, sessionId string) ([]types.Message, error) {
	return nil, nil
}

func (s *Sqlite) CreateMessage(ctx context.Context, message *types.Message) error {
	return nil
}
