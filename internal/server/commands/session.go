package commands

import (
	"context"

	"github.com/urfave/cli/v3"
)

func (c *Commander) sessionCommand() *cli.Command {
	return &cli.Command{
		Name:      "session",
		UsageText: "manage sessions",
		Commands: []*cli.Command{
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "create a new session",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name",
						Usage: "meaningful name for the session",
					},
				},
				Action: func(_ context.Context, cmd *cli.Command) error {
					c.log.Info(cmd.String("name"))
					// s, err := c.db.CreateSession(c.sesh.Context())
					return nil
				},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			c.log.Info("session command invoked", "userID", c.user.ID)
			return nil
		},
	}
}
