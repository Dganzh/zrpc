package codec

import pb "github.com/Dganzh/zrpc/core"

type Codec interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
	GetType() pb.CodecType
}
