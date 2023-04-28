package types

import (
	"encoding/xml"
	"io"
)

type StringMap map[string]string

type xmlStringMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (s StringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(s) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range s {
		e.Encode(xmlStringMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (s *StringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*s = make(StringMap)
	for {
		var e xmlStringMapEntry
		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*s)[e.XMLName.Local] = e.Value
	}
	return nil
}
