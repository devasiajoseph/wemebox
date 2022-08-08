package mail

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/mailgun/mailgun-go"
)

type Attachment struct {
	Filename string
	Content  []byte
}

type Message struct {
	From        string
	To          string
	Subject     string
	Body        string
	Attachments []Attachment
}

//SendMailGunAttachment Sends mail with attachment from mail gun api
func SendMailGunAttachment(msg Message) error {
	// Your available domain names can be found here:
	// (https://app.mailgun.com/app/domains)
	yourDomain := core.Config.MailDomain // e.g. mg.yourcompany.com

	// You can find the Private API Key in your Account Menu, under "Settings":
	// (https://app.mailgun.com/app/account/security)
	privateAPIKey := core.Config.MailPKey
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(msg.From, msg.Subject, msg.Body, msg.To)
	for _, at := range msg.Attachments {
		message.AddBufferAttachment(at.Filename, at.Content)
	}

	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message	with a 10 second timeout
	resp, id, err := mg.Send(message)

	if err != nil {
		fmt.Println("Error sending email")
		fmt.Println(err)
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}

//SendMailGun Sends mail from mail gun api
func SendMailGun(sender string, recipient string, subject string, body string) error {
	// Your available domain names can be found here:
	// (https://app.mailgun.com/app/domains)
	yourDomain := core.Config.MailDomain // e.g. mg.yourcompany.com

	// You can find the Private API Key in your Account Menu, under "Settings":
	// (https://app.mailgun.com/app/account/security)
	privateAPIKey := core.Config.MailPKey
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message	with a 10 second timeout
	resp, id, err := mg.Send(message)

	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}
