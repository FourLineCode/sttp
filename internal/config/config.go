package config

type Config struct {
	Addr string
	Port string
}

func GetConfig() Config {
	return Config{
		Addr: "tcp",
		Port: ":8080",
	}
}
