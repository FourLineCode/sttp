package config

import "flag"

type Config struct {
	Port int
}

type flags struct {
	port int
}

func GetConfig() Config {
	f := getFlags()

	return Config{
		Port: f.port,
	}
}

func getFlags() flags {
	defaultPort := flag.Int("port", 6969, "local port to host server (default: 6969)")

	flag.Parse()

	return flags{
		port: *defaultPort,
	}
}
