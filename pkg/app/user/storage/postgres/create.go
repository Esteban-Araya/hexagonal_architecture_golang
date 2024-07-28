package storage

import (
	"Project/pkg/app/user/model"
	"log"
	
)
func (s UserStorage) Create(u model.CreateUserModel) (bool, error){

	_, err := s.DB.Exec("insert into users (username, email, created_at) values ($1,$2,$3)",
		u.Username, u.Email, u.CreatedAt)
	log.Printf("%s",u.Username)
	if err != nil {
		log.Printf("cannot save the patient, %s", err.Error())
		return false, err
	} 

	return true, err
}