package types

import "fmt"

type Format string

const (
	Native          Format = "native"
	XML             Format = "xml"
	JSON            Format = "json"
	ProtocolBuffers Format = "protobuf"
	Avro            Format = "avro"
	YAML            Format = "yaml"
	MessagePack     Format = "message_pack"
)

func ParseFormat(format string) (Format, error) {
	switch format {
	case "native":
		return Native, nil
	case "xml":
		return XML, nil
	case "json":
		return JSON, nil
	case "protobuf":
		return ProtocolBuffers, nil
	case "avro":
		return Avro, nil
	case "yaml":
		return YAML, nil
	case "message_pack":
		return MessagePack, nil
	default:
		return "", fmt.Errorf("format %v is not supported", format)
	}
}
