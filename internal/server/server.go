package server

import (
	"fmt"

	"github.com/FourLineCode/sttp/pkg/sttp"
	"github.com/sirupsen/logrus"
)

func Run(port uint16) {
	server := sttp.NewServer(port)

	server.OnMessage(func(packet sttp.Message) {
		fmt.Printf("%v> %v\n", packet.RemoteAddr.String(), packet.Body)
	})

	if err := server.Listen(); err != nil {
		logrus.Panic("Error while listening for connections ", err.Error())
	}
}
