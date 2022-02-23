package main

import "fmt"

type sample struct {
	words   string
	list    []int
	dict    map[string]int
	integer int
	float   float32
}

type Serializer interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
}

func main() {
	fmt.Println("Hello, world.")
}
