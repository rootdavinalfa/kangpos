/*
 * Copyright (c) 2021.
 * Davin Alfarizky Putra Basudewa
 * KangPos API
 * Mail Delivery via API
 */

package common

import (
	"fmt"
	"net/smtp"
)

type ClientConfig struct {
	Address  string
	Host     string
	Username string
	Password string
	Port     string
}

type MailTemplate struct {
	From    string
	To      string
	Cc      string
	Subject string
}

type Mailer struct {
	Configuration *ClientConfig
	Auth          smtp.Auth
}

func (m *Mailer) InitializeMail(clientConfig *ClientConfig) {
	m.Configuration = clientConfig
	m.Configuration.Address = fmt.Sprintf("%s:%s", clientConfig.Host, clientConfig.Port)

	// Authentication
	m.Auth = smtp.PlainAuth("", clientConfig.Username, clientConfig.Password, clientConfig.Host)
}

func (m *Mailer) BuildEmailData(template *MailTemplate, body string) []byte {
	headerMessage := ""
	header := make(map[string]string)
	header["From"] = template.From
	header["To"] = template.To
	header["Cc"] = template.Cc
	header["Subject"] = template.Subject
	for key, value := range header {
		headerMessage += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	final := fmt.Sprintf("%s\r\n%s", headerMessage, body)
	return []byte(final)
}

func (m *Mailer) SendMail(to string, body []byte) error {

	err := smtp.SendMail(
		m.Configuration.Address,
		m.Auth,
		m.Configuration.Username,
		[]string{to},
		[]byte(body),
	)

	if err != nil {
		return err
	}
	return nil
}
