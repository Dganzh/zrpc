package zrpc

import (
	log "github.com/Dganzh/zlog"
	"testing"
)

type HelloService struct {

}


type HelloArgs struct {
	Name string
}

type HelloReply struct {
	OK bool
}


func (h *HelloService) HelloWorld(args *HelloArgs, reply *HelloReply) {
	log.Infow("receive hello world", "name", args.Name)
	reply.OK = true
}



func TestServerAndClient(t *testing.T) {
	addr := "localhost:8989"
	s := NewServer(addr)
	s.Register(&HelloService{})
	go s.Start()

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


