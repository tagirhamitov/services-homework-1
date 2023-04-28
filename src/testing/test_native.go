package testing

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/tagirhamitov/services_practice_1/types"
)

func GetNativeSerializedSize(object types.Struct) (int, error) {
	data, err := marshalNative(object)
	if err != nil {
		return 0, err
	}

	return len(data), nil
}

func TestNativeSerialization(object types.Struct) (time.Duration, error) {
	return Measure(func() error {
		buffer := bytes.NewBuffer(nil)
		encoder := gob.NewEncoder(buffer)
		return encoder.Encode(object)
	})
}

func TestNativeDeserialization(object types.Struct) (time.Duration, error) {
	data, err := marshalNative(object)
	if err != nil {
		return 0, err
	}

	return Measure(func() error {
		buffer := bytes.NewBuffer(data)
		decoder := gob.NewDecoder(buffer)
		return decoder.Decode(&object)
	})
}

func marshalNative(object types.Struct) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(object)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
