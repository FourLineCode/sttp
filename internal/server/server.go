package server

import (
	"fmt"

	"github.com/FourLineCode/sttp/pkg/logger"
	"github.com/FourLineCode/sttp/pkg/sttp"
)

func Run(port uint16) {
	server := sttp.NewServer(port)

	server.OnMessage(func(packet sttp.Message) {
		suffix := fmt.Sprintf("[%v] %v:%v:", packet.Time.Format("Jan 02 | 15:04"), packet.Host, packet.Port)
		logger.Custom(suffix, logger.Green, packet.Body)
	})

	if err := server.Listen(); err != nil {
		logger.Panic("Error while listening for connections %v", err.Error())
	}
}
