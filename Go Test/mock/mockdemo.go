//go:generate mockgen -source=./mockdemo.go -destination=./mock_mockdemo.go -package=mockdemo Mail

package mockdemo

import "time"

type Mail interface {
	sendMail(subject string, sender string, dst string, body string) error
}

type MC struct {
	m Mail
}

func NewMC(m Mail) *MC {
	return &MC{m: m}
}

func sign() string {
	return time.Now().Format(time.RFC1123)
}

var getSign = sign

func (c *MC) WriteAndSend(subject string, sender string, dst string, body string) error {
	s := getSign()
	body = body + s
	err := c.m.sendMail(subject, sender, dst, body)
	if err != nil {
		return err
	}
	return nil
}
