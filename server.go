package zrpc

import (
	"context"
	"errors"
	"github.com/Dganzh/zrpc/codec"
	"github.com/Dganzh/zrpc/config"
	pb "github.com/Dganzh/zrpc/core"
	"google.golang.org/grpc"
	"log"
	"net"
	"reflect"
	"strings"
)

// Handler Func(ctx context.Context, Args ArgsType, ReplyType)
type Handler struct {
	Func      reflect.Value
	NParams   int
	NReturn   int
	ArgsType  reflect.Type // maybe nil
	ReplyType reflect.Type // maybe nil
}

func newHandler(method reflect.Method) *Handler {
	h := &Handler{}
	h.Func = method.Func
	h.NParams = method.Type.NumIn()
	h.NReturn = method.Type.NumOut()
	if h.NParams == 3 {
		h.ArgsType = method.Func.Type().In(2)
	}
	if h.NReturn == 2 {
		h.ReplyType = method.Func.Type().Out(0)
	}
	return h
}

type Service struct {
	name     string
	handlers map[string]*Handler
}

type Server struct {
	pb.UnimplementedRPCServer
	services map[string]*Service
	codecMap map[pb.CodecType]codec.Codec
	cfg      *config.ServerConfig
}

func (s *Server) Register(service interface{}) {
	svcType := reflect.TypeOf(service)
	svcValue := reflect.ValueOf(service)
	handlers := make(map[string]*Handler, svcType.NumMethod())
	for m := 0; m < svcType.NumMethod(); m++ {
		name := svcType.Method(m).Name
		log.Printf("register name %s", name)
		method := svcValue.MethodByName(name)
		if !method.IsValid() {
			continue
		}
		h := &Handler{}
		h.Func = method
		h.NParams = method.Type().NumIn()
		h.NReturn = method.Type().NumOut()
		if h.NParams == 2 {
			h.ArgsType = method.Type().In(0)
			h.ReplyType = method.Type().In(1)
		}
		handlers[name] = h
	}
	if len(handlers) > 0 {
		svcName := svcType.Elem().Name()
		s.services[svcName] = &Service{name: svcName, handlers: handlers}
		log.Println("register success", "service name", svcName, "services", s.services)
	}
}

func (s *Server) Call(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return s.call(ctx, request)
}

func (s *Server) call(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	curCodec, ok := s.codecMap[request.GetCodecType()]
	if !ok {
		return nil, errors.New("Invalid CodecType " + request.GetCodecType().String())
	}
	handlerInfo := strings.Split(request.GetHandler(), ".")
	if len(handlerInfo) != 2 {
		return nil, errors.New("Invalid Handler " + request.GetHandler())
	}
	srvName := handlerInfo[0]
	srv, ok := s.services[srvName]
	if !ok {
		return nil, errors.New("Not exists service " + srvName)
	}
	handlerName := handlerInfo[1]
	handler, ok := srv.handlers[handlerName]
	if !ok {
		return nil, errors.New("Not exists handler " + handlerName)
	}
	var params []reflect.Value
	data := request.GetData()
	args := reflect.New(handler.ArgsType.Elem())
	if err := curCodec.Decode(data, args.Interface()); err != nil {
		return nil, errors.New("BadRequest " + err.Error())
	}
	params = append(params, args)
	reply := reflect.New(handler.ReplyType.Elem())
	params = append(params, reply)
	handler.Func.Call(params)
	resp := &pb.Response{}
	data, err := curCodec.Encode(reply.Interface())
	if err != nil {
		return nil, errors.New("InternalServerError: invalid return value" + err.Error())
	}
	resp.Data = data
	return resp, nil
}

func (s *Server) Start() {
	if len(s.services) <= 0 {
		log.Fatalf("Please register service!")
		return
	}
	lis, err := net.Listen("tcp", s.cfg.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterRPCServer(srv, s)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewServer(cfg *config.ServerConfig) *Server {
	if cfg == nil {
		cfg = config.GetDefaultServerConfig()
	}
	cm := make(map[pb.CodecType]codec.Codec, len(pb.CodecType_name))
	for ct, _ := range pb.CodecType_name {
		switch pb.CodecType(ct) {
		case pb.CodecType_GOB:
			cm[pb.CodecType_GOB] = codec.NewGobObject()
		case pb.CodecType_JSON:
			cm[pb.CodecType_JSON] = codec.NewJsonCodec()
		}
	}
	s := &Server{
		services: make(map[string]*Service),
		codecMap: cm,
		cfg:      cfg,
	}
	return s
}
