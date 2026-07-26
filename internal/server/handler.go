package server

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/tobiaskrok/plop/internal/db"
	"github.com/tobiaskrok/plop/internal/server/commands"
	"github.com/tobiaskrok/plop/internal/types"
	gossh "golang.org/x/crypto/ssh"
)

type SessionHandler struct {
	db db.DB
}

// authenticate looks up the user by their public key fingerprint,
// creating an account on first connection.
func (h *SessionHandler) authenticate(sesh ssh.Session) (*types.User, error) {
	fingerprint := gossh.FingerprintSHA256(sesh.PublicKey())

	user, err := h.db.FindUserByFingerprint(sesh.Context(), fingerprint)
	if err != nil {
		return nil, fmt.Errorf("find user by fingerprint: %w", err)
	}

	if user == nil {
		user, err = h.db.CreateUser(sesh.Context(), sesh.User(), fingerprint)
		if err != nil {
			return nil, fmt.Errorf("create user: %w", err)
		}
		log.Info("user created", "fingerprint", fingerprint, "name", user.Name)
	}

	log.Info("user authenticated", "fingerprint", fingerprint, "name", user.Name)
	return user, nil
}

func (h *SessionHandler) HandleFunc(next ssh.Handler) ssh.Handler {
	return func(sesh ssh.Session) {

		user, err := h.authenticate(sesh)
		if err != nil {
			log.Error("authentication failed", "err", err)
			wish.Fatalln(sesh, "❌ Unable to authenticate")
			return
		}

		if len(sesh.Command()) > 0 {
			cmdr := commands.New(h.db, sesh, user)
			args := append([]string{"plop"}, sesh.Command()...)
			if err := cmdr.Root().Run(sesh.Context(), args); err != nil {
				log.Error("command failed", "err", err)
			}
			return
		}

		userSesh := &UserSession{sesh}
		if userSesh.IsPTY() {
			wish.Println(sesh, "Not implemented yet...!")
		}
	}
}

// // copied from snips.sh
// func readFile(sesh *UserSession, maxSize uint64) ([]byte, error) {
// 	content := make([]byte, 0)
// 	size := uint64(0)
// 	for {
// 		buf := make([]byte, UploadBufferSize)
// 		n, err := sesh.Read(buf)
// 		isEOF := errors.Is(err, io.EOF)
// 		if err != nil && !isEOF {
// 			return nil, err
// 		}
//
// 		size += uint64(n)
// 		content = append(content, buf[:n]...)
//
// 		if size > maxSize {
// 			return nil, ErrFileTooLarge
// 		}
//
// 		if isEOF {
// 			if size == 0 {
// 				return nil, ErrEmptyContent
// 			}
// 			return content, nil
// 		}
// 	}
// }
