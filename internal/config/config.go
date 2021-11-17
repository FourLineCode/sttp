package config

import (
	"flag"
	"net"

	"github.com/FourLineCode/sttp/pkg/sttp"
	"github.com/sirupsen/logrus"
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
	port := sttp.DefaultPort

retry:
	ln, err := net.Listen("tcp", sttp.TransformPort(port))
	if err != nil {
		port++
		goto retry
	}
	ln.Close()
	logrus.Warnf("Couldn't connect to default port. Falling back to %v", port)

	defaultPort := flag.Uint64("port", uint64(port), "local port to host server (default: 6969)")

	flag.Parse()

	return flags{
		port: uint16(*defaultPort),
	}
}
