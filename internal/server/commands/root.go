package commands

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/tobiaskrok/plop/internal/db"
	"github.com/tobiaskrok/plop/internal/types"

	cli "github.com/urfave/cli/v3"
)

type Commander struct {
	db   db.DB
	sesh ssh.Session
	user *types.User
	log  *log.Logger
}

func New(db db.DB, sesh ssh.Session, user *types.User) *Commander {
	return &Commander{
		db:   db,
		sesh: sesh,
		user: user,
		log:  log.FromContext(sesh.Context()).With("user", user.Name),
	}
}

func (c *Commander) Root() *cli.Command {
	return &cli.Command{
		Name:           "plop",
		ExitErrHandler: func(ctx context.Context, cmd *cli.Command, err error) {},
		Writer:         c.sesh,
		Reader:         c.sesh,
		ErrWriter:      c.sesh.Stderr(),
		Commands: []*cli.Command{
			c.sessionCommand(),
		},
	}
}
