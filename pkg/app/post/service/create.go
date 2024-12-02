package service

import (
	"Project/pkg/app/post/domain"

	"time"
)

func (s PostService) CreatePostService(p domain.CreatePostModel, user_id int) (err error) {

	p.CreatedAt = time.Now().UTC()
	p.UserID = user_id
	err = s.PostStorage.Create(p)

	return err
}
