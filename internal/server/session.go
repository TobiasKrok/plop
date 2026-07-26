package server

import (
	"github.com/charmbracelet/ssh"
)

type UserSession struct {
	ssh.Session
}

// remote interactive sessions
func (s *UserSession) IsPTY() bool {

	_, _, isPty := s.Pty()
	return isPty

}
