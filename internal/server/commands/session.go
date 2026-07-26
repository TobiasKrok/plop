package commands

import (
	"context"

	"github.com/urfave/cli/v3"
)

func (c *Commander) sessionCommand() *cli.Command {
	return &cli.Command{
		Name:      "session",
		UsageText: "manage sessions",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			c.log.Info("session command invoked", "userID", c.user.ID)
			return nil
		},
	}
}
