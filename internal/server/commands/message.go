package commands

import (
	"database/sql"
	"fmt"

	"github.com/charmbracelet/wish"
	"github.com/tobiaskrok/plop/internal/types"
	"github.com/tobiaskrok/plop/internal/ui/term"
	"github.com/tobiaskrok/plop/internal/utils"
)

// createMessage will create a message and a session if its not provided
func (c *Commander) createMessage() (*types.Message, error) {

	var sesh *types.Session
	if id, exists := utils.GetAndParseSessionID(c.sesh.User()); exists {

		s, err := c.db.FindSessionById(c.sesh.Context(), id)
		if err == nil {
			sesh = s
		} else {
			if err != sql.ErrNoRows {
				return nil, err
			} // else seshId is nil so we create it later
		}
	}
	if sesh != nil && sesh.Visibility == types.VisibilityPrivate && c.user.ID != sesh.UserID {
		return nil, utils.ErrUnauthorized
	}

	if sesh == nil {
		// name defaults to generated ID
		s, err := c.createSession(c.sesh.Context(), "")
		sesh = s
		c.log.Info("created new session", "session_id", sesh.ID)

		if err != nil {
			return nil, fmt.Errorf("failed to create a new session: %w", err)
		}
	}
	content, err := utils.ReadInput(c.sesh, utils.MaxContentSize)
	if err != nil {
		return nil, fmt.Errorf("failed to read content: %w", err)
	}

	message, err := c.db.CreateMessage(c.sesh.Context(), content, sesh.ID, c.user.ID)
	if err != nil {
		return nil, err
	}

	wish.Println(c.sesh, "YOOOOOOOOO")

	term.RenderSessionCreated(c.sesh, *sesh)
	c.log.Info("message created", "message_id", message.ID, "session_id", sesh.ID, "len", message.Size)
	return nil, nil
}
