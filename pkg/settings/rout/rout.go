package rout

import (
	"net/http"
	"Project/pkg/app/user"
	"database/sql"
)

func StartListener(s *http.Server,db *sql.DB) (error){

	user.UserRout(db)	
	
	return s.ListenAndServe()
}