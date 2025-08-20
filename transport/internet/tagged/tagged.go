package tagged

import (
	"context"

	"github.com/asimov1234/newv/common/net"
	"github.com/asimov1234/newv/features/routing"
)

type DialFunc func(ctx context.Context, dispatcher routing.Dispatcher, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
