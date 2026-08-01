package utils

import (
	id "github.com/matoous/go-nanoid/v2"
)

func NewID() string {
	return id.Must(7)
}

func NewShortId() string {
	return id.MustGenerate("abcdefghijklmnopqrstuvwxyz123456789", 5)
}
