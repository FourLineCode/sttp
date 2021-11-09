package config

import "flag"

type Config struct {
	Port int
}

type Flags struct {
	Port int
}

func GetConfig(flags Flags) Config {
	return Config{
		Port: flags.Port,
	}
}

func GetFlags() Flags {
	defaultPort := flag.Int("port", 6969, "local port to host server (default: 6969)")

	flag.Parse()

	return Flags{
		Port: *defaultPort,
	}
}
