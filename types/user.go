package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minFirstnameLen = 2
	minLastnameLen  = 2
	minPasswodLen   = 7
)

type CreateUserParams struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() []error {
	errors := []error{}
	if len(params.FirstName) < minFirstnameLen {
		errors = append(errors, fmt.Errorf("firstName length should be at lesast %d characters", minFirstnameLen))
	}
	if len(params.LastName) < minLastnameLen {
		errors = append(errors, fmt.Errorf("LastName length should be at lesast %d characters", minLastnameLen))
	}
	if len(params.Password) < minPasswodLen {
		errors = append(errors, fmt.Errorf("password length should be at lesast %d characters", minPasswodLen))
	}
	if !isEmailValid(params.Email) {
		errors = append(errors, fmt.Errorf("email is invalid"))
	}
	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"EncryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}
