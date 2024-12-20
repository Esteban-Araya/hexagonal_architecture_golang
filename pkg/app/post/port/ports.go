package port

import (
	"Project/pkg/app/post/domain"
)

type PostServiceInterface interface {
	CreatePostService(p domain.CreatePostModel, user_id int) (err error)
	GetPostByIdService(id_post int) (*domain.GetPost, error)
	GetAllPostService() (*[]domain.GetPost, error)
	PostUpdateService(p domain.PostUpdate, user_id int) error
}

type PostStorageInterface interface {
	Create(p domain.CreatePostModel) error
	GetPostById(id_post int) (*domain.GetPost, error)
	GetAllPost() (*[]domain.GetPost, error)
	PostUpdate(p domain.PostUpdate) error
}
