package db

import "context"

type DB interface {
	CreateDroplet(ctx context.Context)
}
