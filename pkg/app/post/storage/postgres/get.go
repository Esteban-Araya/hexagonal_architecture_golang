package postgres

import "Project/pkg/app/post/domain"

func (s PostStorage) GetPostById(id_post int) (*domain.GetPost, error) {

	rows, err := s.DB.Query("select id, user_id, title, content, created_at from post where id=$1", id_post)
	if err != nil {
		return nil, err
	}
	var post domain.GetPost
	for rows.Next() {
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			return nil, err
		}
	}
	return &post, nil
}
