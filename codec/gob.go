package codec

import (
	"bytes"
	"encoding/gob"
	pb "github.com/Dganzh/zrpc/core"
)

type Gob struct {
}

func NewGobObject() *Gob {
	return &Gob{}
}

func (g *Gob) Encode(e interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(e)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (g *Gob) Decode(data []byte, e interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(e)
}

func (g *Gob) GetType() pb.CodecType {
	return pb.CodecType_GOB
}

//func (g *Gob) RegisterName(name string, v interface{}) {
//	gob.RegisterName(name, v)
//}
