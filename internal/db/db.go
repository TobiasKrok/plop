package db

import (
	"context"

	"github.com/tobiaskrok/plop/internal/types"
)

type DB interface {
	Migrate(ctx context.Context) error
	FindUserByFingerprint(ctx context.Context, fingerprint string) (*types.User, error)
	CreateUser(ctx context.Context, name, fingerprint string) (*types.User, error)
	CreateSession(ctx context.Context, session *types.Session) error
	GetMessagesBySession(ctx context.Context, sessionId string) ([]types.Message, error)
	CreateMessage(ctx context.Context, message *types.Message) error
}
