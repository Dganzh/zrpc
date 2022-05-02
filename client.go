package zrpc

import (
	"context"
	"github.com/Dganzh/zrpc/codec"
	"github.com/Dganzh/zrpc/config"
	pb "github.com/Dganzh/zrpc/core"
	"google.golang.org/grpc"
	"log"
)

type Client struct {
	client pb.RPCClient
	codec  codec.Codec
	cfg    *config.ClientConfig
}

func NewClient(cfg *config.ClientConfig) *Client {
	if cfg == nil {
		cfg = config.GetDefaultClientConfig()
	}
	conn, err := grpc.Dial(cfg.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	var c codec.Codec
	switch cfg.CodecType {
	case pb.CodecType_GOB:
		c = codec.NewGobObject()
	case pb.CodecType_JSON:
		c = codec.NewJsonCodec()
	default:
		log.Fatalf("invalid codec type: %d", cfg.CodecType)
	}
	return &Client{
		client: pb.NewRPCClient(conn),
		codec:  c,
		cfg:    cfg,
	}
}

func (c *Client) Call(handler string, arg interface{}, reply interface{}) bool {
	ctx, cancel := context.WithTimeout(context.Background(), c.cfg.CallTimeout)
	defer cancel()
	data, _ := c.codec.Encode(arg)
	r, err := c.client.Call(ctx, &pb.Request{
		Handler:   handler,
		CodecType: c.codec.GetType(),
		Data:      data,
	})
	if err != nil {
		log.Fatalf("RPC call handler=%s, failed: %v", handler, err)
		return false
	}
	_ = c.codec.Decode(r.GetData(), reply)
	log.Printf("Call %s Result: %+v", handler, reply)
	return true
}
