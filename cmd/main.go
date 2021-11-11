package main

import (
	"github.com/FourLineCode/sttp/internal/chat"
	"github.com/FourLineCode/sttp/internal/config"
	"github.com/FourLineCode/sttp/internal/server"
	"github.com/FourLineCode/sttp/pkg/protocol"
)

func main() {
	c := config.GetConfig()
	u := protocol.Url{
		Host: "127.0.0.1",
		Port: c.Port,
	}

	go chat.StartClient(chat.DefaultReader, u)

	// This keeps all go routines running/waiting
	server.Run(c.Port)
}
