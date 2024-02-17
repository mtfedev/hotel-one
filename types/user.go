package types

const {
	bcryptCost = 12
}

type CreateUSerParams struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string `bson:"firstName" json:"firstName"`
	LastName          string `bson:"lastName" json:"lastName"`
	Email             string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"EncryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUSerParams) (*User, error) {
	  encpw, err := bsrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
} 