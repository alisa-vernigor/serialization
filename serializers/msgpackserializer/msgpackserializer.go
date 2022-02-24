package msgpackserializer

import (
	"bytes"

	"github.com/vmihailenco/msgpack"
)

func NewSerializer() *Serializer {
	return &Serializer{}
}

type Serializer struct {
}

func (s *Serializer) Marshal(data interface{}) ([]byte, error) {
	var res bytes.Buffer
	res.Reset()

	enc := msgpack.NewEncoder(&res)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}

	return res.Bytes(), nil
}

func (s *Serializer) Unmarshal(data []byte, format interface{}) error {
	input := bytes.NewReader(data)

	dec := msgpack.NewDecoder(input)
	err := dec.Decode(format)

	return err
}
