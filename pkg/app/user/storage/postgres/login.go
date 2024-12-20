package postgres

import (
	"Project/pkg/app/user/domain"
)

func (s UserStorage) GetUserByEmail(email string) (*domain.User, error) {

	rows, err := s.DB.Query("select id, email, password from users where email=$1", email)
	if err != nil {
		return nil, err
	}
	var user domain.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			return nil, err
		}
	}
	if user.ID == 0 {
		return nil, domain.EmailOrPasswordIsWrong
	}
	return &user, nil
}
