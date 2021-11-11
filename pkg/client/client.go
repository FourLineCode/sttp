package client

import (
	"net"
	"time"

	"github.com/FourLineCode/sttp/pkg/protocol"
	"github.com/FourLineCode/sttp/pkg/sttp"
)

type Client interface {
	SendMessage(protocol.Url, protocol.Packet) error
}

type client struct{}

func NewClient() Client {
	return client{}
}

func (c client) SendMessage(url protocol.Url, packet protocol.Packet) error {
	conn, err := net.Dial("tcp", sttp.TransformPort(url.Port))
	if err != nil {
		return err
	}
	defer conn.Close()

	packet.LocalAddr = conn.LocalAddr()
	packet.RemoteAddr = conn.RemoteAddr()
	packet.Time = time.Now()
	message, err := protocol.Marshal(packet)
	if err != nil {
		return err
	}

	n, err := conn.Write([]byte(message))
	if err != nil || n == 0 {
		return err
	}

	return nil
}
