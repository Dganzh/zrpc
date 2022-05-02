package zrpc

import (
	"github.com/Dganzh/zrpc/config"
	pb "github.com/Dganzh/zrpc/core"
	"log"
	"testing"
	"time"
)

type HelloService struct{}

type HelloArgs struct {
	Name string
}

type HelloReply struct {
	OK bool
}

func (h *HelloService) HelloWorld(args *HelloArgs, reply *HelloReply) {
	log.Println("Receive hello world", "name", args.Name)
	reply.OK = true
}

func TestServerAndClient(t *testing.T) {
	// start server
	addr := "localhost:8989"
	scfg := &config.ServerConfig{Addr: addr}
	s := NewServer(scfg)
	s.Register(&HelloService{})
	go s.Start()

	time.Sleep(time.Millisecond)

	// create client and send RPC
	ccfg := &config.ClientConfig{Addr: addr, CodecType: pb.CodecType_GOB, CallTimeout: 3 * time.Second}
	client := NewClient(ccfg)
	args := HelloArgs{Name: "dganzh"}
	reply := HelloReply{}
	ok := client.Call("HelloService.HelloWorld", &args, &reply)
	if !ok {
		t.Error("client Call HelloService.HelloWorld failed.")
	}
	if !reply.OK {
		t.Error("client Call HelloService.HelloWorld result incorrect.")
	}
}

func TestCodec(t *testing.T) {
	addr := "localhost:8989"
	scfg := &config.ServerConfig{Addr: addr}
	s := NewServer(scfg)
	s.Register(&HelloService{})
	go s.Start()

	time.Sleep(time.Millisecond)
	ccfg := &config.ClientConfig{Addr: addr, CodecType: pb.CodecType_JSON, CallTimeout: 3 * time.Second}
	client := NewClient(ccfg)
	args := HelloArgs{Name: "dganzh"}
	reply := HelloReply{}
	ok := client.Call("HelloService.HelloWorld", &args, &reply)
	if !ok {
		t.Error("client Call HelloService.HelloWorld failed.")
	}
	if !reply.OK {
		t.Error("client Call HelloService.HelloWorld result incorrect.")
	}

}
