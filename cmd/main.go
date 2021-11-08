package main

import (
	"github.com/FourLineCode/sttp/internal/config"
	"github.com/FourLineCode/sttp/internal/server"
)

func main() {
	c := config.GetConfig()
	server.Run(c.Addr, c.Port)
}
