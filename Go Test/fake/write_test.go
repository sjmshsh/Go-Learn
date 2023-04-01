package write_test

import (
	"errors"
	"fmt"
	write "gomock"
	"testing"
)

type fakeM struct {
}

func (f fakeM) SendMail() error {
	fmt.Println("xxx")
	return errors.New("fake err")
}

func TestNewMailClient(t *testing.T) {
	mc := write.NewMailClient(fakeM{})
	err := mc.WriteAndSend()
	if err != nil {
		t.Error(err)
	}
}
