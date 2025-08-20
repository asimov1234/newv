package command_test

import (
	"context"
	"testing"

	"github.com/asimov1234/newv/app/dispatcher"
	"github.com/asimov1234/newv/app/log"
	. "github.com/asimov1234/newv/app/log/command"
	"github.com/asimov1234/newv/app/proxyman"
	_ "github.com/asimov1234/newv/app/proxyman/inbound"
	_ "github.com/asimov1234/newv/app/proxyman/outbound"
	"github.com/asimov1234/newv/common"
	"github.com/asimov1234/newv/common/serial"
	"github.com/asimov1234/newv/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
