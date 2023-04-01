package write

import "fmt"

type Mail interface {
	SendMail() error
}

type MailClient struct {
	m Mail
}

func NewMailClient(m Mail) *MailClient {
	return &MailClient{
		m: m,
	}
}

func (mc *MailClient) WriteAndSend() error {
	err := mc.m.SendMail()
	fmt.Println("write and sendmail")
	return err
}
