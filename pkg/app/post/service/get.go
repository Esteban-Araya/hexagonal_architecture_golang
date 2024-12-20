package service

import "Project/pkg/app/post/domain"

func (s PostService) GetPostByIdService(id_post int) (*domain.GetPost, error) {
	post, err := s.PostStorage.GetPostById(id_post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s PostService) GetAllPostService() (*[]domain.GetPost, error) {
	post, err := s.PostStorage.GetAllPost()
	if err != nil {
		return nil, err
	}

	return post, nil
}
