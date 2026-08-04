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

	noti.Messagef(` Continue sending to this session via: 
ssh s:%s@plop.sh 

	`, session.ID)

	noti.Render(sesh)
}
