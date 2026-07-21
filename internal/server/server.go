package server

import (
	"fmt"
	"net"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
)

type Server struct {
	*ssh.Server
}

func New() (*Server, error) {
	sshServer, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort("localhost", "2222")),
		wish.WithHostKeyPath("host_key"),
		wish.WithPublicKeyAuth(func(_ ssh.Context, _ ssh.PublicKey) bool {
			return true
		}),
		wish.WithMiddleware(
			func(next ssh.Handler) ssh.Handler {
				return func(sess ssh.Session) {
					wish.Println(sess, fmt.Sprintf("Hello %s", sess.User()))
					next(sess)
				}
			},
			logging.Middleware(),
		),
	)

	if err != nil {
		return nil, err
	}

	return &Server{sshServer}, nil
}

