package mail

import (
	"testing"

	"github.com/devasiajoseph/wemebox/core"
)

func TestMailgun(t *testing.T) {
	core.SetupConfig()
	m := Mail{Sender: "icmr@nic.in", Recipient: "devasiajoseph@gmail.com", Subject: "Testing mail gun", Body: "This is a tes email"}
	err := m.SendMailGun()
	if err != nil {
		t.Errorf("Error sending mailgun")
	}
}
