package sample

import (
	"encoding/xml"
	"io"
)

type Map map[string]int32

type Sample struct {
	Words   string
	List    []int32
	Dict    Map
	Integer int32
	Float   float32
}

type xmlMapEntry struct {
	XMLName xml.Name
	Value   int32 `xml:",chardata"`
}

func (m Map) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (m *Map) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = Map{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}
