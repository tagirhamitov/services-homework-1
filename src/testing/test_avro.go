package testing

import (
	"time"

	"github.com/hamba/avro"
	"github.com/tagirhamitov/services_practice_1/types"
)

const schemaStr = `
{
	"type": "record",
	"name": "Struct",
	"fields": [
		{"name": "Str", "type": "string"},
        {"name": "Arr", "type": {"type": "array", "items": "int"}},
        {"name": "Dict", "type": {"type": "map", "values": "string"}},
        {"name": "Int", "type": "int"},
        {"name": "Float", "type": "double"}
	]
}`

func GetAvroSerializedSize(object types.Struct) (int, error) {
	schema, err := avro.Parse(schemaStr)
	if err != nil {
		return 0, err
	}

	data, err := avro.Marshal(schema, object)
	if err != nil {
		return 0, err
	}

	return len(data), nil
}

func TestAvroSerialization(object types.Struct) (time.Duration, error) {
	schema, err := avro.Parse(schemaStr)
	if err != nil {
		return 0, err
	}

	return Measure(func() error {
		_, err := avro.Marshal(schema, object)
		return err
	})
}

func TestAvroDeserialization(object types.Struct) (time.Duration, error) {
	schema, err := avro.Parse(schemaStr)
	if err != nil {
		return 0, err
	}

	data, err := avro.Marshal(schema, object)
	if err != nil {
		return 0, err
	}

	return Measure(func() error {
		return avro.Unmarshal(schema, data, &object)
	})
}
