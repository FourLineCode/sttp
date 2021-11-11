package client

import (
	"net"

	"github.com/FourLineCode/sttp/pkg/sttp"
)

type Client interface {
	SendMessage(uint16, string) error
}

type client struct{}

func NewClient() Client {
	return client{}
}

func (c client) SendMessage(port uint16, message string) error {
	conn, err := net.Dial("tcp", sttp.TransformPort(port))
	if err != nil {
		return err
	}
	defer conn.Close()

	n, err := conn.Write([]byte(message + "\n"))
	if err != nil || n == 0 {
		return err
	}

	return nil
}
