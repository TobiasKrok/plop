package types

import "time"

type User struct {
	ID        string
	CreatedAt time.Time
	Name      string
}
