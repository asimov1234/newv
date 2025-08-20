package all

import (
	"github.com/asimov/newv/main/commands/all/api"
	"github.com/asimov/newv/main/commands/all/convert"
	"github.com/asimov/newv/main/commands/all/tls"
	"github.com/asimov/newv/main/commands/base"
)

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		convert.CmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
	)
}
