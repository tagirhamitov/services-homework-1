package testing

import (
	"time"

	"github.com/tagirhamitov/services_practice_1/types"
	"gopkg.in/yaml.v3"
)

func GetYAMLSerializedSize(object types.Struct) (int, error) {
	data, err := yaml.Marshal(object)
	if err != nil {
		return 0, err
	}

	return len(data), nil
}

func TestYAMLSerialization(object types.Struct) (time.Duration, error) {
	return Measure(func() error {
		_, err := yaml.Marshal(object)
		return err
	})
}

func TestYAMLDeserialization(object types.Struct) (time.Duration, error) {
	data, err := yaml.Marshal(object)
	if err != nil {
		return 0, err
	}
	return Measure(func() error {
		return yaml.Unmarshal(data, &object)
	})
}
