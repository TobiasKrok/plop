package utils

import id "github.com/matoous/go-nanoid/v2"

func NewID() string {
	return id.Must(7)
}
