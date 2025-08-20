package tcp

import (
	"github.com/asimov/newv/common"
	"github.com/asimov/newv/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
