package server

import (
	"fmt"
	"time"

	"github.com/FourLineCode/sttp/pkg/sttp"
	"github.com/sirupsen/logrus"
)

func Run(port int) {
	server, err := sttp.NewServer(port)
	if err != nil {
		logrus.Panic("Error initializing sttp server ", err.Error())
	}

	server.OnMessage(func(packet sttp.Message) {
		fmt.Printf("%v | Recieved from %v\n", packet.Time.Local().Format(time.RFC822), packet.RemoteAddr.String())
		fmt.Printf("Message: %v\n", packet.Body)
	})

	if err := server.Listen(); err != nil {
		logrus.Panic("Error while listening for connections ", err.Error())
	}
}
