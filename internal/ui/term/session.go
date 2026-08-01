package term

import (
	"github.com/charmbracelet/ssh"
	"github.com/tobiaskrok/plop/internal/types"
	"github.com/tobiaskrok/plop/internal/ui/styles"
)

func RenderSessionCreated(sesh ssh.Session, session types.Session) {

	noti := Notification{
		Title: "␥ Session created",
		Color: styles.Colors.Success,
	}

	noti.Messagef(`Continue this session: 

		%$

	`)

	noti.Render(sesh)
}
