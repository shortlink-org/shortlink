package transform

import (
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/transform/protobuf"
)

type Transform interface {
	Serialize(interface{}) ([]byte, error)
	Deserialize([]byte, interface{}) (interface{}, error)
}

func Serialize(data interface{}) ([]byte, error) {
	return T.Serialize(data)
}

func Deserialize(data []byte, i interface{}) (interface{}, error) {
	return T.Deserialize(data, i)
}

var T Transform

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("TRANSPORT_TYPE", "protobuf")
	transformType := viper.GetString("TRANSPORT_TYPE")

	switch transformType {
	case "protobuf":
		T = &protobuf.Protobuf{}
	// case "avro":
	// 	T = avro.Avro{}
	default:
		T = &protobuf.Protobuf{}
	}
}
