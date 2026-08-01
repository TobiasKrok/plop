package types

import "time"

type Visibility string

const (
	VisibilityPrivate Visibility = "private"
	VisibilityPublic  Visibility = "public"
)

type Session struct {
	ID         string
	Name       string
	UserID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Visibility Visibility
}
