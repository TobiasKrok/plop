package db

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"time"

	_ "github.com/mattn/go-sqlite3"
	goose "github.com/pressly/goose/v3"
	"github.com/tobiaskrok/plop/internal/types"
	"github.com/tobiaskrok/plop/internal/utils"
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

func (s *Sqlite) CreateUser(ctx context.Context, name, fingerprint string) (*types.User, error) {
	tx, err := s.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	u := &types.User{
		ID:          utils.NewID(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Name:        name,
		Fingerprint: fingerprint,
	}

	const query = `

		INSERT INTO users (
	id,
	created_at,
	updated_at,
	name,
	fingerprint
	) VALUES (?, ?, ?, ?, ?)
		`

	if _, err := tx.ExecContext(ctx, query,
		u.ID,
		u.CreatedAt,
		u.UpdatedAt,
		u.Name,
		u.Fingerprint,
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
func (s *Sqlite) FindUserByFingerprint(ctx context.Context, fingerprint string) (*types.User, error) {

	const query = `
		SELECT
		id,
		created_at,
		updated_at,
		name,
		fingerprint
		FROM users
		WHERE fingerprint = ?
		`
	user := &types.User{}
	row := s.QueryRowContext(ctx, query, fingerprint)

	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Name,
		&user.Fingerprint,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// TODO: max sessions
func (s *Sqlite) CreateSession(ctx context.Context, name string) (*types.Session, error) {

	tx, err := s.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	u := &types.Session{
		ID:        utils.NewID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}

	const query = `

		INSERT INTO users (
	id,
	created_at,
	updated_at,
	name,
	fingerprint
	) VALUES (?, ?, ?, ?, ?)
		`

	if _, err := tx.ExecContext(ctx, query,
		u.ID,
		u.CreatedAt,
		u.UpdatedAt,
		u.Name,
		u.Fingerprint,
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
	return nil
}

func (s *Sqlite) GetMessagesBySession(ctx context.Context, sessionId string) ([]types.Message, error) {
	return nil, nil
}

func (s *Sqlite) CreateMessage(ctx context.Context, message *types.Message) error {
	return nil
}
