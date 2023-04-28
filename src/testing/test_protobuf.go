package testing

import (
	"time"

	"github.com/tagirhamitov/services_practice_1/pb"
	"github.com/tagirhamitov/services_practice_1/types"
	"google.golang.org/protobuf/proto"
)

func GetProtoBufSerializedSize(object types.Struct) (int, error) {
	pbObject := convert(object)

	data, err := proto.Marshal(&pbObject)
	if err != nil {
		return 0, err
	}

	return len(data), nil
}

func TestProtoBufSerialization(object types.Struct) (time.Duration, error) {
	pbObject := convert(object)
	return Measure(func() error {
		_, err := proto.Marshal(&pbObject)
		return err
	})
}

func TestProtoBufDeserialization(object types.Struct) (time.Duration, error) {
	pbObject := convert(object)
	data, err := proto.Marshal(&pbObject)
	if err != nil {
		return 0, err
	}
	return Measure(func() error {
		return proto.Unmarshal(data, &pbObject)
	})
}

func convert(object types.Struct) pb.Struct {
	return pb.Struct{
		Str:   object.Str,
		Arr:   object.Arr,
		Dict:  object.Dict,
		Int:   object.Int,
		Float: object.Float,
	}
}
