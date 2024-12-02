package service

import (
	"Project/pkg/app/post/domain"
)

func (s PostService) PostUpdateService(p domain.PostUpdate, id_Post int) error {

	if p.UserID != id_Post {
		return domain.YouHaveNotAccessError
	}

	err := s.PostStorage.PostUpdate(p)

	if err != nil {
		return err
	}

	return nil
}
