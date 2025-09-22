package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null;size:200" binding:"required,max=200"`
	Content   string         `json:"content" gorm:"not null;type:text" binding:"required"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	User     User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:PostID"`
}

type PostCreateRequest struct {
	Title   string `json:"title" binding:"required,max=200"`
	Content string `json:"content" binding:"required"`
}

type PostUpdateRequest struct {
	Title   string `json:"title" binding:"max=200"`
	Content string `json:"content"`
}

type PostResponse struct {
	ID        uint              `json:"id"`
	Title     string            `json:"title"`
	Content   string            `json:"content"`
	UserID    uint              `json:"user_id"`
	User      UserResponse      `json:"user,omitempty"`
	Comments  []CommentResponse `json:"comments,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type PostListResponse struct {
	ID        uint         `json:"id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	UserID    uint         `json:"user_id"`
	User      UserResponse `json:"user,omitempty"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func (p *Post) ToResponse() PostResponse {
	response := PostResponse{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}

	if p.User.ID != 0 {
		response.User = p.User.ToResponse()
	}

	if len(p.Comments) > 0 {
		response.Comments = make([]CommentResponse, len(p.Comments))
		for i, comment := range p.Comments {
			response.Comments[i] = comment.ToResponse()
		}
	}

	return response
}

func (p *Post) ToListResponse() PostListResponse {
	response := PostListResponse{
		ID:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		UserID:    p.UserID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}

	if p.User.ID != 0 {
		response.User = p.User.ToResponse()
	}

	return response
}
