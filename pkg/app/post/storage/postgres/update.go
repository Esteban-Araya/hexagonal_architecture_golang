package postgres

import (
	"Project/pkg/app/post/domain"
	"log"

	"github.com/lib/pq"
)

func (s PostStorage) PostUpdate(p domain.PostUpdate) error {

	_, err := s.DB.Exec("update post set title=$1, content=$2 where id=$3",
		p.Title, p.Content, p.ID)

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
