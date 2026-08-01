package types

import "time"

type Message struct {
	ID        string
	Size      uint64
	Content   []byte
	UserID    string
	SessionID string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//TODO compress
