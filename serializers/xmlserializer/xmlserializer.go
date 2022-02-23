package jsonserializer

import (
	"encoding/xml"
)

type Serializer struct {
}

func (s *Serializer) Marshal(data interface{}) ([]byte, error) {
	res, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Serializer) Unmarshal(data []byte, format interface{}) error {
	err := xml.Unmarshal(data, format)

	return err
}
