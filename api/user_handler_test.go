package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/mtfedev/hotel-one/db"
	"github.com/mtfedev/hotel-one/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdburi = "mongodb://localhost:27017"
	dbname    = "hotel-one-test"
)

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client, dbname),
	}
}
func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)
	
	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email: "some@con.com",
		FirstName: "James",
		LastName: "foo",
		Password: "iuedgubvhff",
		
	}
	b, _:+ json.Marshal(params)
	req := httptest.NewRequest("POST", "/", nil, bytes.NewReader(b))
	req.Header.Add("Content-Types", "application/jason")
	resp, err:= app.Test(req)
	if err != nil{
		t.Error(err)
	}
	var user types.User 
	json.NewDecoder(req.Body).Decode(&user)
	if user.FirstName != params.FirstName {
		t.Errorf("expexted firstnem %s but got %s", params.FirstName, user.FirstName)
	}
}
