package serializers

type Serializer interface {
	Marshal(data interface{}) ([]byte, error)
	Unmarshal(data []byte, format interface{}) error
}
