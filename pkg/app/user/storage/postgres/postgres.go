package postgres

import (
	"Project/pkg/app/user/storage"
)

type UserStorage struct {
	DB storage.DBinterface
}
