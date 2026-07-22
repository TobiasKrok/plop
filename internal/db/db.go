package db

import (
	"context"

	"github.com/tobiaskrok/plop/internal/types"
)

type DB interface {
	Migrate(ctx context.Context) error
	CreateSession(ctx context.Context, session *types.Session) error
	GetMessagesBySession(ctx context.Context, sessionId string) ([]types.Message, error)
	CreateMessage(ctx context.Context, message *types.Message) error
}
