package main

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Inster(k K, v V) error {
	m.data[k] = v
	return nil
}
func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func cup[T any, B any](val T, x B) {
	fmt.Println(val)
}

func main() {

	//	m1 := NewCustomMap[string, int]()
	//	m1.Inster("bar", 1)
	//	m1.Inster("nor", 2)

	//	m2 := NewCustomMap[int, float64]()
	//	m2.Inster(2, 9.33)
	//m2.Inster(3, 100.2313)
}
