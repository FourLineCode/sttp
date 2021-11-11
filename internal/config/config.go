package config

import (
	"flag"

	"github.com/FourLineCode/sttp/pkg/sttp"
)

type Config struct {
	Port uint16
}

type flags struct {
	port uint16
}

func GetConfig() Config {
	f := getFlags()

	return Config{
		Port: f.port,
	}
}

func getFlags() flags {
	defaultPort := flag.Uint64("port", sttp.DefaultPort, "local port to host server (default: 6969)")

	flag.Parse()

	return flags{
		port: uint16(*defaultPort),
	}
}
