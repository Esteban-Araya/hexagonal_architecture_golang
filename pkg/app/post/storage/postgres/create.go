package postgres

import (
	"Project/pkg/app/post/domain"
	"log"

	"github.com/lib/pq"
)

func (s PostStorage) Create(p domain.CreatePostModel) error {

	_, err := s.DB.Exec("insert into post (user_id, title, content, created_at) values ($1,$2,$3,$4)",
		p.UserID, p.Title, p.Content, p.CreatedAt)

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
