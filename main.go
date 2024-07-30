package main

type Database struct {
	user string
}

type Server struct {
	db *Database // uintprt -> 8 bytes long
}

func (s *Server) GetUsersFromDB() string {
	// golang is going to "dereference" the db pointer
	// it's going to look up the memory address of the pointer
	if s.db == nil {
		panic("database == nil hence, is not initialized")
	}
	return s.db.user
}

func main() {
	// since we did not specify a database in out server, there is no valid memory address
	// -> for this reason, you always need to check your pointers for nil
	s := &Server{}
	s.GetUsersFromDB()
}
