package commands

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/tobiaskrok/plop/internal/db"
	"github.com/tobiaskrok/plop/internal/types"

	"github.com/urfave/cli/v3"
)

func sessionCommand(db db.DB, sesh ssh.Session) *cli.Command {

	return &cli.Command{
		Name:      "session",
		UsageText: "manage sessions",
		Action: func(ctx context.Context, cmd *cli.Command) error {

			logger := log.FromContext(sesh.Context())

			user := sesh.Context().Value(types.UserIDContextKey)

			return nil
		},
	}
}
