package db

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
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
func (s *Sqlite) CreateSession(ctx context.Context, name string, userID string) (*types.Session, error) {

	tx, err := s.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	id := utils.NewShortId()
	if name == "" {
		name = id
	}
	sesh := &types.Session{
		ID:         id,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Name:       name,
		UserID:     userID,
		Visibility: types.VisibilityPrivate, // by default its private
	}

	const query = `

		INSERT INTO sessions (
	id,
	user_id,
	created_at,
	updated_at,
	name,
	visibility
	) VALUES (?, ?, ?, ?, ?, ?)
		`

	if _, err := tx.ExecContext(ctx, query,
		sesh.ID,
		sesh.UserID,
		sesh.CreatedAt,
		sesh.UpdatedAt,
		sesh.Name,
		sesh.Visibility,
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return sesh, nil
}

func (s *Sqlite) GetMessagesBySessionID(ctx context.Context, sessionId string) ([]types.Message, error) {

	const query = `
		SELECT 
		id,
		created_at,
		updated_at,
		size,
		content,
		session_id,
		user_id
		FROM messages
		WHERE session_id = ?
	 	`
	rows, err := s.QueryContext(ctx, query, sessionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messsages []types.Message
	for rows.Next() {
		m := &types.Message{}
		err := rows.Scan(
			&m.ID,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.Size,
			&m.Content,
			&m.SessionID,
			&m.UserID,
		)
		if err != nil {
			return nil, err
		}
		decompressed, err := utils.Decompress(m.Content)
		if err != nil {
			return nil, fmt.Errorf("failed to decompress message data: %w", err)
		}
		m.Content = decompressed
		messsages = append(messsages, *m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messsages, nil
}

func (s *Sqlite) CreateMessage(ctx context.Context, content []byte, sessionID, userID string) (*types.Message, error) {

	tx, err := s.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	compressed, err := utils.Compress(content)
	if err != nil {
		return nil, err
	}
	message := &types.Message{
		ID:        utils.NewID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    userID,
		SessionID: sessionID,
		Content:   compressed,
		Size:      uint64(len(content)),
	}

	const query = `
		INSERT INTO messages (
	id,
	created_at,
	updated_at,
	user_id,
	session_id,
	content,
	size
	) VALUES (?, ?, ?, ?, ?, ?, ?)
		`

	if _, err := tx.ExecContext(ctx, query,
		message.ID,
		message.CreatedAt,
		message.UpdatedAt,
		message.UserID,
		message.SessionID,
		message.Content,
		message.Size,
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return message, nil
}

func (s *Sqlite) FindSessionById(ctx context.Context, sessionID string) (*types.Session, error) {

	const query = `
		SELECT
		id,
		user_id,
		created_at,
		updated_at,
		name,
		visibility
		FROM sessions
		WHERE  id = ?
		`
	sesh := &types.Session{}
	row := s.QueryRowContext(ctx, query, sessionID)

	err := row.Scan(
		&sesh.ID,
		&sesh.UserID,
		&sesh.CreatedAt,
		&sesh.UpdatedAt,
		&sesh.Name,
		&sesh.Visibility,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return sesh, nil
}
