package main

import "fmt"

type Putter interface {
	Put(id int, val any) error
	fmt.Stringer
}
type Storage interface {
	Get(id int) (any, error)
	Putter
}

type simpePutter struct{}

func (*simpePutter) String() string { return "" }
func (s *simpePutter) Put(id int, val any) error {
	return nil
}

type BobStorage struct{}

func (s *BobStorage) Get(id int) (any, error) {
	return nil, nil
}
func (s *BobStorage) Put(id int, val any) error {
	return nil
}

type Server struct {
	store Storage
}

func updateValue(id int, val any, p Putter) error {
	return p.Put(id, val)
}
func main() {
	simpePutter := &simpePutter{}
	updateValue(1, "bob", s.store)
}
