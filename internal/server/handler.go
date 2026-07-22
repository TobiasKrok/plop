package server

import (
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/tobiaskrok/plop/internal/db"
)

type SessionHandler struct {
	db db.DB
}

func (h *SessionHandler) UserAuth() func(next ssh.Handler) ssh.Handler {

	return func(next ssh.Handler) ssh.Handler {

		return func(sesh ssh.Session) {

			fingerprint := go
			next(sesh)
		}
	}
}

func (h *SessionHandler) HandleFunc(_ ssh.Handler) ssh.Handler {
	return func(sesh ssh.Session) {

		userSesh := &UserSession{sesh}

		if userSesh.IsPTY() {
			wish.Println(sesh, "Not implemented yet...!")
		}
	}
} 
