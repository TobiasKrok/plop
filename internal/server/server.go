package server

import (
	"net"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
	"github.com/tobiaskrok/plop/internal/db"
)

type Server struct {
	*ssh.Server
}

func New(db db.DB) (*Server, error) {

	sessionHandler := &SessionHandler{
		db: db,
	}
	sshServer, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort("localhost", "2222")),
		wish.WithHostKeyPath("host_key"),
		wish.WithPublicKeyAuth(func(_ ssh.Context, _ ssh.PublicKey) bool {
			return true
		}),
		wish.WithMiddleware(
			sessionHandler.HandleFunc,
			logging.Middleware(),
		),
	)

	if err != nil {
		return nil, err
	}

	return &Server{sshServer}, nil
}

