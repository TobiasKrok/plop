package types

import "time"

type User struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Fingerprint string
	Name        string
}
