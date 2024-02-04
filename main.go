package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue((context.Background()), "username", "Bob")
	userID, err := fetchUserID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response took %v -> &+v\n:", time.Since(start), userID)
}

func fetchUserID(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	val := ctx.Value("username")
	fmt.Println("the value =", val)

	type result struct {
		userId string
		err    error
	}
	resultch := make(chan result, 1)

	go func() {
		res, err := thirdpartyHTTPCall()
		resultch <- result{
			userId: res,
			err:    err,
		}
	}()
	select {
	//Done()
	//1 the context timeout is exceeded
	//2 the context has been maually canceled -> Cancel()
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultch:
		return res.userId, res.err
	}
}

func thirdpartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 90)
	return "user id 1 ", nil
}
