package serializer

import(
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
)


func ProtobufToJSON(message proto.Message) (string, error){
	marshaller := jsonpb.Marshaler{
		OrigName: true,
		EnumsAsInts: false,
		EmitDefaults: true,
		Indent: " ",
	}

	return marshaller.MarshalToString(message)
}
