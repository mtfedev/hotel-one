package main

type Database struct {
	user string
}
type Server struct {
	db *Database // unitprt -> 8 bytes long
}

func (s *Server) GetUserFromDB() string {
	//golang is going to "dereference" the db pointer
	// its going to lookup the memory address of the pointer
	if s.db == nil {
		return ""
	}
	return s.db.user
}

func main() {
	s := &Server{}
	s.GetUserFromDB()
}
