package yaml

import (
	"context"
	"io"

	"github.com/asimov1234/newv/common"
	"github.com/asimov1234/newv/common/cmdarg"
	"github.com/asimov1234/newv/common/errors"
	"github.com/asimov1234/newv/core"
	"github.com/asimov1234/newv/infra/conf"
	"github.com/asimov1234/newv/infra/conf/serial"
	"github.com/asimov1234/newv/main/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "YAML",
		Extension: []string{"yaml", "yml"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				cf := &conf.Config{}
				for i, arg := range v {
					errors.LogInfo(context.Background(), "Reading config: ", arg)
					r, err := confloader.LoadConfig(arg)
					if err != nil {
						return nil, errors.New("failed to read config: ", arg).Base(err)
					}
					c, err := serial.DecodeYAMLConfig(r)
					if err != nil {
						return nil, errors.New("failed to decode config: ", arg).Base(err)
					}
					if i == 0 {
						// This ensure even if the muti-json parser do not support a setting,
						// It is still respected automatically for the first configure file
						*cf = *c
						continue
					}
					cf.Override(c, arg)
				}
				return cf.Build()
			case io.Reader:
				return serial.LoadYAMLConfig(v)
			default:
				return nil, errors.New("unknown type")
			}
		},
	}))
}
