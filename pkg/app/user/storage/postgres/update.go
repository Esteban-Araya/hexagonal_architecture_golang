package postgres

import (
	"Project/pkg/app/user/domain"
	"log"

	"github.com/lib/pq"
)

func (s UserStorage) UserUpdateName(u domain.UserUpdateName, id_user int) error {

	_, err := s.DB.Exec("update users set username=$1 where id=$2",
		u.Username, id_user)

	if err != nil {
		// code := err.(*pq.Error).Code
		code := err.(*pq.Error).SQLState()
		if code == "23505" {
			log.Printf("%s", err.Error())
			return domain.EmailAlreadyExistError
		}
		log.Printf("cannot save the patient, %s", err.Error())
		return domain.DatabaseError
	}

	return nil
}
