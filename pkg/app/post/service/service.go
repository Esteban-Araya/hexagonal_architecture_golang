package service

import (
	"Project/pkg/app/post/port"
)

type PostService struct {
	PostStorage port.PostStorageInterface
}
