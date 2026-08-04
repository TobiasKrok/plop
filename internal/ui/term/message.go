package term

import (
	"github.com/charmbracelet/ssh"
	"github.com/tobiaskrok/plop/internal/types"
	"github.com/tobiaskrok/plop/internal/ui/styles"
)

func RenderMessageCreated(sesh ssh.Session, session types.Session, message types.Message) {

	noti := Notification{
		Title: "␥ Message published",
		Color: styles.Colors.Success,
	}

	noti.Messagef(` Message has been published to session '%s'

To view this message:
ssh s:%@plop.sh view %s

	`, session.ID, message.ID)

	noti.Render(sesh)
}
