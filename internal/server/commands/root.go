package commands

import (
	"context"
	"fmt"

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
		Name:        "plop.sh",
		Description: "share messages with agents over ssh",
		UsageText:   "ssh plop.sh",
		ExitErrHandler: func(ctx context.Context, cmd *cli.Command, err error) {
			if err != nil {
				fmt.Fprintln(c.sesh.Stderr(), "error:", err)
			}
		},
		Writer:    c.sesh,
		Reader:    c.sesh,
		ErrWriter: c.sesh.Stderr(),
		Commands: []*cli.Command{
			c.sessionCommand(),
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// emoty means create
			_, err := c.createMessage()
			return err
		},
	}
}
