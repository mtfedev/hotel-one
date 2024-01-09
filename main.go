package main

import (
	"fmt"
)

type NumberStoter interface {
	GetAll() ([]int, error)
	Put(int) error
}

type ApiServer struct {
	numberStoter NumberStoter
}

type MongoDbNmumberStrore struct {
	// some values
}

func (m MongoDbNmumberStrore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4}, nil
}

func (m MongoDbNmumberStrore) Put(number int) error {
	fmt.Println("Store the number ito the db storage")
	return nil
}

func main() {
	apiServer := ApiServer{
		numberStoter: MongoDbNmumberStrore{},
	}
	numbers, err := apiServer.numberStoter.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(numbers)
}
