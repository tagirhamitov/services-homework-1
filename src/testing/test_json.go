package testing

import (
	"encoding/json"
	"time"

	"github.com/tagirhamitov/services_practice_1/types"
)

func GetJSONSerializedSize(object types.Struct) (int, error) {
	data, err := json.Marshal(object)
	if err != nil {
		return 0, err
	}

	return len(data), nil
}

func TestJSONSerialization(object types.Struct) (time.Duration, error) {
	return Measure(func() error {
		_, err := json.Marshal(object)
		return err
	})
}

func TestJSONDeserialization(object types.Struct) (time.Duration, error) {
	data, err := json.Marshal(object)
	if err != nil {
		return 0, err
	}

	return Measure(func() error {
		return json.Unmarshal(data, &object)
	})
}
