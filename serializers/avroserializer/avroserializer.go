package avroserializer

import (
	"github.com/hamba/avro"
)

const schema = `{
	"name": "MyClass",
	"type": "record",
	"namespace": "com.acme.avro",
	"fields": [
	  {
		"name": "Words",
		"type": "string"
	  },
	  {
		"name": "List",
		"type": {
		  "type": "array",
		  "items": "int"
		}
	  },
	  {
		"name": "Dict",
		"type": {
		  "name": "Dict",
		  "type": "map",
		  "values": "int"
		}
	  },
	  {
		"name": "Integer",
		"type": "int"
	  },
	  {
		"name": "Float",
		"type": "float"
	  }
	]
  }`

func NewSerializer() *Serializer {
	return &Serializer{avro.MustParse(schema)}
}

type Serializer struct {
	schema avro.Schema
}

func (s *Serializer) Marshal(data interface{}) ([]byte, error) {
	res, err := avro.Marshal(s.schema, data)
	if err != nil {
		return nil, nil
	}

	return res, nil
}

func (s *Serializer) Unmarshal(data []byte, format interface{}) error {
	err := avro.Unmarshal(s.schema, data, format)
	if err != nil {
		return err
	}
	return err
}
