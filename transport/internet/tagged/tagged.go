package tagged

import (
	"context"

	"github.com/asimov/newv/common/net"
	"github.com/asimov/newv/features/routing"
)

type DialFunc func(ctx context.Context, dispatcher routing.Dispatcher, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
