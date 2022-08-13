package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	marshalel := jsonpb.Marshaler{
		EnumsAsInts: false,
		EmitDefaults: true,
		Indent: " ",
		OrigName: true,
	}

	return marshalel.MarshalToString(message)
}

// JSONToProtobufMessage converts JSON string to protocol buffer message
func JSONToProtobufMessage(data string, message proto.Message) error {
	return jsonpb.UnmarshalString(data, message)
}