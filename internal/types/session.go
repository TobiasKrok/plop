package types

import "time"

type Session struct {
	ID         string
	Name       string
	UserID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Visibility string
}
