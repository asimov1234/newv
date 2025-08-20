package vmess_test

import (
	"testing"

	"github.com/asimov1234/newv/common"
	"github.com/asimov1234/newv/common/protocol"
	"github.com/asimov1234/newv/common/uuid"
	. "github.com/asimov1234/newv/proxy/vmess"
)

func toAccount(a *Account) protocol.Account {
	account, err := a.AsAccount()
	common.Must(err)
	return account
}

func BenchmarkUserValidator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := NewTimedUserValidator()

		for j := 0; j < 1500; j++ {
			id := uuid.New()
			v.Add(&protocol.MemoryUser{
				Email: "test",
				Account: toAccount(&Account{
					Id: id.String(),
				}),
			})
		}

		common.Close(v)
	}
}
