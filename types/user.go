package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minFirstnameLen = 2
	minLastnameLen  = 2
	minPasswodLen   = 7
)

type UpateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (p UpateUserParams) ToBSONM() bson.M {
	m := bson.M{}
	if len(p.FirstName) > 0 {
		m["firstName"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		m["lasttName"] = p.LastName
	}
	return m
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstnameLen {
		errors["fisrtName"] = fmt.Sprintf("firstName length should be at lesast %d characters", minFirstnameLen)
	}
	if len(params.LastName) < minLastnameLen {
		errors["lastName"] = fmt.Sprintf("LastName length should be at lesast %d characters", minLastnameLen)
	}
	if len(params.Password) < minPasswodLen {
		errors["password"] = fmt.Sprintf("password length should be at lesast %d characters", minPasswodLen)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = fmt.Sprintf("email is invalid")
	}
	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
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
