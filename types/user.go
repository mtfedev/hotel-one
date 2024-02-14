package types

type User struct {
	ID        string `bson:"_id" json:"id"`
	FirstName string `bson:"firstName" json:"firstName"`
	lastName  string `bson:"lastName" json:"lastName"`
}
