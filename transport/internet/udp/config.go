package udp

import (
	"github.com/asimov1234/newv/common"
	"github.com/asimov1234/newv/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
