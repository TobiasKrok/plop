package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
	"github.com/tobiaskrok/plop/internal/types"
	"github.com/tobiaskrok/plop/internal/ui/term"
	"github.com/tobiaskrok/plop/internal/utils"
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
				Action: func(ctx context.Context, cmd *cli.Command) error {
					s, err := c.createSession(ctx, cmd.String("name"))
					if err != nil {
						c.log.Error("failed to create session", "err", err)
						return err
					}
					fmt.Println(c.sesh.Subsystem())
					term.RenderSessionCreated(c.sesh, *s)
					c.log.Info("created session", "session_id", s.ID, "session_name", s.Name)
					return nil
				},
			},
			{
				Name:      "show",
				Usage:     "show a session's messages",
				ArgsUsage: "<session-id>",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					sessionID := cmd.Args().First()
					if sessionID == "" {
						return errors.New("session id required")
					}
					messages, err := c.showSession(ctx, sessionID)
					if err != nil {
						return err
					}
					for _, m := range messages {
						fmt.Fprintln(c.sesh, string(m.Content))
					}
					return nil
				},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			//TODO: print help here
			return nil
		},
	}
}

// TODO: check visibility
func (c *Commander) showSession(ctx context.Context, sessionID string) ([]types.Message, error) {

	sesh, err := c.db.FindSessionById(ctx, sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("session %s does not exist", sessionID)
		}
		return nil, err
	}

	if sesh.Visibility == types.VisibilityPrivate && c.user.ID != sesh.UserID {
		return nil, utils.ErrUnauthorized
	}

	messages, err := c.db.GetMessagesBySessionID(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	c.log.Debug("fetched messages from session", "session_id", sesh.ID, "user_id", sesh.UserID, "total_messages", len(messages))
	return messages, nil
}

func (c *Commander) createSession(ctx context.Context, name string) (*types.Session, error) {
	for range 5 {
		s, err := c.db.CreateSession(ctx, name, c.user.ID)
		if err == nil {
			return s, nil
		}
		sqliteErr, ok := errors.AsType[sqlite3.Error](err)
		if !ok || sqliteErr.ExtendedCode != sqlite3.ErrConstraintUnique {
			return nil, err
		}
		c.log.Warn("generated ID that already existed while creating a session, will retry")
	}
	return nil, errors.New("failed to generate a unique session id")
}
