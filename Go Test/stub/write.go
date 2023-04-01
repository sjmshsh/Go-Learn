package write

import (
	"fmt"
	"time"
)

type Mail interface {
	SendMail(body string) error
}

type MailClient struct {
	m Mail
}

func NewMailClient(m Mail) *MailClient {
	return &MailClient{
		m: m,
	}
}

func (mc *MailClient) WriteAndSend(sender, body string) error {
	singText := GenSign(sender)
	newBody := body + singText

	err := mc.m.SendMail(newBody)

	fmt.Println("write and sendmail")
	return err
}

var GenSign = sign

func sign(sender string) string {
	return fmt.Sprintf("%s --- %s", sender, time.Now().Format(time.RFC1123))
}
