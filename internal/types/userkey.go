package types

import "time"

type UserKey struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Fingerprint string
	UserID      string
}
