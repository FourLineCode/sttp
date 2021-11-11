package server

import (
	"fmt"

	"github.com/FourLineCode/sttp/internal/chat"
	"github.com/FourLineCode/sttp/pkg/sttp"
	"github.com/sirupsen/logrus"
)

func Run(port uint16) {
	server := sttp.NewServer(port)

	server.OnMessage(func(packet sttp.Message) {
		// fmt.Printf("%v | Recieved from %v\n", packet.Time.Local().Format(time.RFC822), packet.RemoteAddr.String())
		fmt.Printf("%v> %v\n", packet.RemoteAddr.String(), packet.Body)
	})

	chatClient := chat.NewChatClient()
	go func() {
		if err := chatClient.Start(); err != nil {
			logrus.Panic("Error starting chat client ", err.Error())
		}
	}()

	if err := server.Listen(); err != nil {
		logrus.Panic("Error while listening for connections ", err.Error())
	}
}
