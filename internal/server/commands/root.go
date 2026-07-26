package commands

import (
	"context"

	"github.com/charmbracelet/ssh"
	"github.com/tobiaskrok/plop/internal/db"

	cli "github.com/urfave/cli/v3"
)

func RootCommand(db db.DB, sesh ssh.Session) *cli.Command {

	return &cli.Command{
		Name:           "plop",
		ExitErrHandler: func(ctx context.Context, c *cli.Command, err error) {},
		Writer:         sesh,
		Reader:         sesh,
		ErrWriter:      sesh.Stderr(),
		Commands: []*cli.Command{
			sessionCommand(db, sesh),
		},
	}
}
