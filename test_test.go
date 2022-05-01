package zrpc

import (
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
	s := NewServer(addr)
	s.Register(&HelloService{})
	go s.Start()

	time.Sleep(time.Millisecond)

	// create client and send RPC
	client := NewClient(addr)
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
