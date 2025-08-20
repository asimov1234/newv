package udp

import (
	"github.com/asimov/newv/common/buf"
	"github.com/asimov/newv/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
