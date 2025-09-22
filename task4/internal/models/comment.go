package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Content   string         `json:"content" gorm:"not null;type:text" binding:"required"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	PostID    uint           `json:"post_id" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Post Post `json:"post,omitempty" gorm:"foreignKey:PostID"`
}

type CommentCreateRequest struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
}

type CommentResponse struct {
	ID        uint         `json:"id"`
	Content   string       `json:"content"`
	UserID    uint         `json:"user_id"`
	PostID    uint         `json:"post_id"`
	User      UserResponse `json:"user,omitempty"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func (c *Comment) ToResponse() CommentResponse {
	response := CommentResponse{
		ID:        c.ID,
		Content:   c.Content,
		UserID:    c.UserID,
		PostID:    c.PostID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}

	if c.User.ID != 0 {
		response.User = c.User.ToResponse()
	}

	return response
}
