package mockdemo

import (
	"github.com/golang/mock/gomock"
	"github.com/prashantv/gostub"
	"testing"
)

func TestMC_WriteAndSend(t *testing.T) {
	stubs := gostub.Stub(&getSign, func() string {
		return "xxx"
	})
	defer stubs.Reset()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockMail := NewMockMail(mockCtrl)
	mockMail.EXPECT().sendMail("x1", "x1", "x1", "x1").Return(nil).Times(1)
	mc := NewMC(mockMail)
	mc.WriteAndSend("x1", "x1", "x1", "x1")
}
