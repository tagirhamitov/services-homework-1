package testing

import (
	"time"

	"github.com/tagirhamitov/services_practice_1/types"
)

func GetSerializedSizeFunction(
	format types.Format,
) func(types.Struct) (int, error) {
	switch format {
	case types.Native:
		return GetNativeSerializedSize
	case types.XML:
		return GetXMLSerializedSize
	case types.JSON:
		return GetJSONSerializedSize
	case types.ProtocolBuffers:
		return GetProtoBufSerializedSize
	case types.Avro:
		return GetAvroSerializedSize
	case types.YAML:
		return GetYAMLSerializedSize
	case types.MessagePack:
		return GetMessagePackSerializedSize
	default:
		return nil
	}
}

func GetSerializationFunction(
	format types.Format,
) func(types.Struct) (time.Duration, error) {
	switch format {
	case types.Native:
		return TestNativeSerialization
	case types.XML:
		return TestXMLSerialization
	case types.JSON:
		return TestJSONSerialization
	case types.ProtocolBuffers:
		return TestProtoBufSerialization
	case types.Avro:
		return TestAvroSerialization
	case types.YAML:
		return TestYAMLSerialization
	case types.MessagePack:
		return TestMessagePackSerialization
	default:
		return nil
	}
}

func GetDeserializationFunction(
	format types.Format,
) func(types.Struct) (time.Duration, error) {
	switch format {
	case types.Native:
		return TestNativeDeserialization
	case types.XML:
		return TestXMLDeserialization
	case types.JSON:
		return TestJSONDeserialization
	case types.ProtocolBuffers:
		return TestProtoBufDeserialization
	case types.Avro:
		return TestAvroDeserialization
	case types.YAML:
		return TestYAMLDeserialization
	case types.MessagePack:
		return TestMessagePackDeserialization
	default:
		return nil
	}
}
