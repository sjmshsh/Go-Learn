package write_test

import (
	"fmt"
	"github.com/prashantv/gostub"
	write "gomock"
	"testing"
)

type fakeM struct {
}

func (f fakeM) SendMail(body string) error {
	fmt.Println("xxx")
	return nil
}

func TestNewMailClient(t *testing.T) {
	//old := write.GenSign
	//// 桩函数 存根函数
	//write.GenSign = func(sender string) string {
	//	return ""
	//}
	//defer func() {
	//	write.GenSign = old
	//}()
	stubs := gostub.Stub(&write.GenSign, func(sender string) string {
		return "xxx"
	})
	defer stubs.Reset()
	mc := write.NewMailClient(fakeM{})
	err := mc.WriteAndSend("1550693033@qq.com", "test content")
	if err != nil {
		t.Error(err)
	}
}
