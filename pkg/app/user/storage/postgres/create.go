package postgres

import (
	"Project/pkg/app/user/domain"
	"log"

	"github.com/lib/pq"
)

func (s UserStorage) Create(u domain.CreateUserModel) error {

	_, err := s.DB.Exec("insert into users (username, email, password, created_at) values ($1,$2,$3,$4)",
		u.Username, u.Email, u.Password, u.CreatedAt)

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
