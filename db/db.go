package db

const (
	DBNAME     = "hotel-one"
	TestDBNAME = "hotel-one-test"
	DBURI      = "mongodb://localhost:27017"
)

type Store struct {
	User UserStore
	Hotel HotelStore
	Room RoomStore
}
