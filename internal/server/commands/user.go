package commands

import "fmt"

func (c *Commander) setUserCurrentSession(sessionId string) error {

	session, err := c.db.FindSessionById(c.sesh.Context(), sessionId)

	if err != nil {
		return err
	}
	if session == nil {
		return fmt.Errorf("session '%s' does not exist", sessionId)
	}
	c.user.CurrentSession = sessionId
	err = c.db.UpdateUser(c.sesh.Context(), c.user)
	if err != nil {
		return err
	}
	return nil
}
