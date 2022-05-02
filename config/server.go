package config

var defaultServerConfig = &ServerConfig{
	Addr: "localhost:5205",
}

type ServerConfig struct {
	Addr string
}

func GetDefaultServerConfig() *ServerConfig {
	return defaultServerConfig
}
