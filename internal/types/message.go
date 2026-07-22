package types

type Message struct {
	ID        string
	Name      string
	Size      uint64
	Content   []byte
	UserID    string
	SessionID string
}

//TODO compress
