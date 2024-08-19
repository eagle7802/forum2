package mocks

import (
	"kkanat/forum/internal/models"
	"time"
)

var mockPost = &models.Post{
	PostID:  1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
}

type PostModel struct{}

func (m *PostModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *PostModel) Get(id int) (*models.Post, error) {
	switch id {
	case 1:
		return mockPost, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *PostModel) Latest() ([]*models.Post, error) {
	return []*models.Post{mockPost}, nil
}
