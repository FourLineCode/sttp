package main

import (
	"github.com/FourLineCode/sttp/internal/chat"
	"github.com/FourLineCode/sttp/internal/config"
	"github.com/FourLineCode/sttp/internal/server"
)

func main() {
	c := config.GetConfig()

	go chat.StartClient(chat.DefaultReader)

	// This keeps all go routines running/waiting
	server.Run(c.Port)
}
