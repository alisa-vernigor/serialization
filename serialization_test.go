package main

import (
	"testing"

	"github.com/alisa-vernigor/serialization/sample"
	"github.com/alisa-vernigor/serialization/serializers"
	"github.com/alisa-vernigor/serialization/serializers/avroserializer"
	"github.com/alisa-vernigor/serialization/serializers/gobserializer"
	"github.com/alisa-vernigor/serialization/serializers/jsonserializer"
	"github.com/alisa-vernigor/serialization/serializers/msgpackserializer"
	"github.com/alisa-vernigor/serialization/serializers/protoserializer"
	"github.com/alisa-vernigor/serialization/serializers/xmlserializer"
	"github.com/alisa-vernigor/serialization/serializers/yamlserializer"
)

func BenchmarkMarshal(b *testing.B) {
	sample := sample.Sample{
		Words: "Lorem ipsum dolor sit amet, consectetur adipiscing" +
			"elit. Mauris adipiscing adipiscing placerat." +
			"Vestibulum augue augue, " +
			"pellentesque quis sollicitudin id, adipiscing.",
		List: []int32{2, 3, 5, 7, 11, 13},
		Dict: map[string]int32{
			"eggs":    175,
			"bacon":   322,
			"sausage": 189,
		},
		Integer: 22,
		Float:   0.22,
	}

	benchmarks := []struct {
		name       string
		serializer serializers.Serializer
	}{
		{"avro", avroserializer.NewSerializer()},
		{"gob", gobserializer.NewSerializer()},
		{"json", jsonserializer.NewSerializer()},
		{"msgpack", msgpackserializer.NewSerializer()},
		{"proto", protoserializer.NewSerializer()},
		{"xml", xmlserializer.NewSerializer()},
		{"yaml", yamlserializer.NewSerializer()},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			var size uint64

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				data, err := bm.serializer.Marshal(sample)
				if err != nil {
					b.Fatal(err)
				}
				size += uint64(len(data))
			}
			b.ReportMetric(float64(size)/float64(b.N), "size")
		})
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	var sample sample.Sample

	benchmarks := []struct {
		name       string
		serializer serializers.Serializer
	}{
		{"avro", avroserializer.NewSerializer()},
		{"gob", gobserializer.NewSerializer()},
		{"json", jsonserializer.NewSerializer()},
		{"msgpack", msgpackserializer.NewSerializer()},
		{"proto", protoserializer.NewSerializer()},
		{"xml", xmlserializer.NewSerializer()},
		{"yaml", yamlserializer.NewSerializer()},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			data, err := bm.serializer.Marshal(sample)
			if err != nil {
				b.Fatal(err)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				err := bm.serializer.Unmarshal(data, &sample)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
