package protoserializer

import (
	"errors"

	"github.com/alisa-vernigor/serialization/sample"
	"google.golang.org/protobuf/proto"
)

func NewSerializer() *Serializer {
	return &Serializer{}
}

type Serializer struct {
}

func (s *Serializer) Marshal(data interface{}) ([]byte, error) {
	godata, ok := data.(sample.Sample)
	if !ok {
		return nil, errors.New("can only convert sample.Sample")
	}

	protodata := &ProtoSample{
		Words:   godata.Words,
		List:    godata.List,
		Dict:    godata.Dict,
		Integer: godata.Integer,
		Float:   godata.Float,
	}

	res, err := proto.Marshal(protodata)
	if err != nil {
		return nil, nil
	}

	return res, nil
}

func (s *Serializer) Unmarshal(data []byte, format interface{}) error {
	godata, ok := format.(*sample.Sample)
	if !ok {
		return errors.New("can only convert *sample.Sample")
	}

	var protodata ProtoSample
	err := proto.Unmarshal(data, &protodata)
	if err != nil {
		return err
	}

	godata.Words = protodata.Words
	godata.List = protodata.List
	godata.Dict = protodata.Dict
	godata.Integer = protodata.Integer
	godata.Float = protodata.Float

	return err
}
