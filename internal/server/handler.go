package server

import (
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/tobiaskrok/plop/internal/db"
	gossh "golang.org/x/crypto/ssh"
)

type SessionHandler struct {
	db db.DB
}

func (h *SessionHandler) UserAuth() func(next ssh.Handler) ssh.Handler {

	return func(next ssh.Handler) ssh.Handler {

		return func(sesh ssh.Session) {

			fingerprint := gossh.FingerprintSHA256(sesh.PublicKey())
			sesh.Context().SetValue(FingerprintContextKey, fingerprint)
			user, err := h.db.FindUserByFingerprint(sesh.Context(), fingerprint)

			if err != nil {
				log.Error("unable to find fingerprint", "err", err)
				wish.Fatalln(sesh, "❌ Unable to authenticate")
				return
			}

			if user == nil {
				wish.Errorln(sesh, "Welcome! Creating your account...")

				user, err := h.db.CreateUser(sesh.Context(), sesh.User(), fingerprint)
				if err != nil {
					log.Error("unable to create user", "err", err)
					wish.Fatalln(sesh, "❌ Unable to create your account, sorry!")
					return
				}
				log.Info("user created", "fingerprint", fingerprint, "name", user.Name)
				wish.Println(sesh, "Welcome! "+user.Name)
			}

			// sesh.Context().SetValue(logger.ContextKey, log)

			log.Info("user authenticated", "fingerprint", fingerprint, "name", user.Name)
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
