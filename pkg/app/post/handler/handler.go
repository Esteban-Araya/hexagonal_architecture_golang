package handler

import (
	"Project/pkg/app/post/port"
)

type PostHandler struct {
	PostService port.PostServiceInterface
}
