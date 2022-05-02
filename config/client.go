package config

import (
	pb "github.com/Dganzh/zrpc/core"
	"time"
)

var defaultClientConfig = &ClientConfig{
	Addr:        "localhost:5205",
	CodecType:   pb.CodecType_GOB,
	CallTimeout: 3 * time.Second,
}

type ClientConfig struct {
	Addr        string
	CodecType   pb.CodecType
	CallTimeout time.Duration
}

func GetDefaultClientConfig() *ClientConfig {
	return defaultClientConfig
}
