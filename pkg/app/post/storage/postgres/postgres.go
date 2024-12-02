package postgres

import (
	"Project/pkg/app/user/storage"
)

type PostStorage struct {
	DB storage.DBinterface
}
