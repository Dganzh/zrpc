package codec

import (
	"encoding/json"
	pb "github.com/Dganzh/zrpc/core"
)

type jsonCodec struct {
}

func NewJsonCodec() *jsonCodec {
	return &jsonCodec{}
}

func (j *jsonCodec) Encode(e interface{}) ([]byte, error) {
	buf, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (j *jsonCodec) Decode(data []byte, e interface{}) error {
	return json.Unmarshal(data, e)
}

func (j *jsonCodec) GetType() pb.CodecType {
	return pb.CodecType_JSON
}
