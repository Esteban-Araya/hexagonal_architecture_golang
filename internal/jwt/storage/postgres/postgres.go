package postgres

import (
	"Project/internal/jwt/storage"
	userDomain "Project/pkg/app/user/domain"
)

type JwtStorage struct {
	DB storage.DBinterface
}

func (s JwtStorage) GetUserById(id float64) (*userDomain.User, error) {

	rows, err := s.DB.Query("select id, email, username from users where id=$1", id)
	if err != nil {
		return nil, err
	}
	var user userDomain.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Email, &user.Username); err != nil {
			return nil, err
		}
	}
	return &user, nil
}
