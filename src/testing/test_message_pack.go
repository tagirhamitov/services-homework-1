package testing

import (
	"time"

	"github.com/vmihailenco/msgpack/v5"

	"github.com/tagirhamitov/services_practice_1/types"
)

func GetMessagePackSerializedSize(object types.Struct) (int, error) {
	data, err := msgpack.Marshal(object)
	if err != nil {
		return 0, err
	}

	return len(data), nil
}

func TestMessagePackSerialization(object types.Struct) (time.Duration, error) {
	return Measure(func() error {
		_, err := msgpack.Marshal(object)
		return err
	})
}

func TestMessagePackDeserialization(object types.Struct) (time.Duration, error) {
	data, err := msgpack.Marshal(object)
	if err != nil {
		return 0, err
	}

	return Measure(func() error {
		return msgpack.Unmarshal(data, &object)
	})
}
