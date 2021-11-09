package main

import (
	"github.com/FourLineCode/sttp/internal/config"
	"github.com/FourLineCode/sttp/internal/server"
)

func main() {
	f := config.GetFlags()

	c := config.GetConfig(f)

	server.Run(c.Port)
}
