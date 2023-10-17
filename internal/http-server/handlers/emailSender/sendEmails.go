package emailsender

import (
	"crypto/tls"
	"io"
	"mime/multipart"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/jordan-wright/email"
)

const (
	smtpServer = "smtp.gmail.com"
	smtpPort   = "587"
	titleMes   = "Doodocs Backend Challenge"
	bodyMes    = "Hello, have a nice day!!!\nHere your file)"
)

func sendEmails(emails []string, file io.Reader, fileHeader *multipart.FileHeader) error {
	username := os.Getenv("LOGIN")
	password := os.Getenv("PASSWORD")
	auth := smtp.PlainAuth("", username, password, smtpServer)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         smtpServer,
	}

	client, err := smtp.Dial(smtpServer + ":" + smtpPort)
	if err != nil {
		return err
	}
	defer client.Quit()

	err = client.StartTLS(tlsConfig)
	if err != nil {
		return err
	}

	if err := client.Auth(auth); err != nil {
		return err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	for _, emailAddr := range emails {
		e := email.NewEmail()
		e.From = os.Getenv("EMAIL_FROM")
		e.To = []string{emailAddr}
		e.Subject = titleMes
		e.Text = []byte(bodyMes)

		_, fileName := filepath.Split(fileHeader.Filename)
		at, err := e.Attach(file, fileName, "application/octet-stream")
		if err != nil {
			return err
		}
		at.Header = fileHeader.Header
		at.Content = content
		err = e.Send(smtpServer+":"+smtpPort, auth)
		if err != nil {
			return err
		}
	}

	return nil
}
