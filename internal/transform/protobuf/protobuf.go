package protobuf

import (
	"github.com/golang/protobuf/proto"
)

type Protobuf struct{}

func (p *Protobuf) Serialize(data interface{}) ([]byte, error) {
	msg := data.(proto.Message)
	return proto.Marshal(msg.(proto.Message))
}

func (p *Protobuf) Deserialize(msg []byte, data interface{}) (interface{}, error) {
	item := data.(proto.Message)

	if err := proto.Unmarshal(msg, item); err != nil {
		return nil, err
	}

	return item, nil
}
