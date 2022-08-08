/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>, June 2018
 */

package mail

import "log"

//Mail structure
type Mail struct {
	Sender    string
	Recipient string
	Subject   string
	Body      string
}

//SendMailGun Sends mail from mail gun api
func (m *Mail) SendMailGun() error {
	err := SendMailGun(m.Sender, m.Recipient, m.Subject, m.Body)
	if err != nil {
		log.Println("Error in mailgun api")
	}
	return err
}
