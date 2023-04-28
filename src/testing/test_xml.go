package testing

import (
	"encoding/xml"
	"time"

	"github.com/tagirhamitov/services_practice_1/types"
)

func GetXMLSerializedSize(object types.Struct) (int, error) {
	data, err := xml.Marshal(object)
	if err != nil {
		return 0, err
	}

	return len(data), nil
}

func TestXMLSerialization(object types.Struct) (time.Duration, error) {
	return Measure(func() error {
		_, err := xml.Marshal(object)
		return err
	})
}

func TestXMLDeserialization(object types.Struct) (time.Duration, error) {
	data, err := xml.Marshal(object)
	if err != nil {
		return 0, err
	}
	return Measure(func() error {
		return xml.Unmarshal(data, &object)
	})
}
