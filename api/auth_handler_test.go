package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/mtfedev/hotel-one/db"
	"github.com/mtfedev/hotel-one/types"
)

func InsertTestUser(t *testing.T, userStore db.UserStore) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     "james@wp.com",
		FirstName: "james",
		LastName:  "foo",
		Password:  "superpassword",
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = userStore.InsertUser(context.TODO(), user)
	if err != nil {
		t.Fatal(err)
	}
	return user
}
func TestAuthenticateWithWrongPassword(t *testing.T) {
	tdb := setup()
	defer tdb.teardown(t)
	insrtedUser := InsertTestUser(t, tdb.UserStore)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.UserStore)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "james@wp.com",
		Password: "superpasswordnotcorrect", //23:00
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/auth ", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expcted http status of 200 but got %d", resp.StatusCode)
	}
	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		t.Error(err)
	}

	if authResp.Token == "" {
		t.Fatalf("expected the JWT token to be present in the auth respose")
	}
	insrtedUser.EncryptedPassword = ""
	if !reflect.DeepEqual(insrtedUser, authResp.User) {
		t.Fatalf("expected the user to be tje inserted user")

	}
}

func TestAuthenticateSuccess(t *testing.T) {
	tdb := setup()
	defer tdb.teardown(t)
	insrtedUser := InsertTestUser(t, tdb.UserStore)

	app := fiber.New()
	authHandler := NewAuthHandler(tdb.UserStore)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "james@wp.com",
		Password: "superpassword",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/auth ", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expcted http status of 200 but got %d", resp.StatusCode)
	}
	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		t.Error(err)
	}

	if authResp.Token == "" {
		t.Fatalf("expected the JWT token to be present in the auth respose")
	}
	insrtedUser.EncryptedPassword = ""
	if !reflect.DeepEqual(insrtedUser, authResp.User) {
		t.Fatalf("expected the user to be tje inserted user")

	}
}
