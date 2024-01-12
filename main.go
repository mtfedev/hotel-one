package main

type Storage interface {
	Get(id int) (any, error)
	Put(id int, val any) error
}

type BobStorage struct{}

func (s *BobStorage) Get(id int) (any, error) {
	return nil, nil
}
func (s *BobStorage) Put(id int) (any, error) {
	return nil, nil
}

type Server struct {
	store Storage
}

func main() {
	s := &Server{}
	s.store.Get(1)
	s.store.Put(1, "bob")
}
