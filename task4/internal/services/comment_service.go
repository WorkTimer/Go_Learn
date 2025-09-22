package services

import (
	"blog-system/internal/models"
	"blog-system/pkg/database"
	"blog-system/pkg/logger"
	"errors"

	"gorm.io/gorm"
)

type CommentService struct{}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (s *CommentService) CreateComment(userID uint, req *models.CommentCreateRequest) (*models.CommentResponse, error) {
	var post models.Post
	if err := database.DB.First(&post, req.PostID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		logger.Error("Database error:", err)
		return nil, errors.New("获取文章失败")
	}

	comment := &models.Comment{
		Content: req.Content,
		UserID:  userID,
		PostID:  req.PostID,
	}

	if err := database.DB.Create(comment).Error; err != nil {
		logger.Error("Failed to create comment:", err)
		return nil, errors.New("评论创建失败")
	}

	if err := database.DB.Preload("User").First(comment, comment.ID).Error; err != nil {
		logger.Error("Failed to load comment with user:", err)
		return nil, errors.New("获取评论信息失败")
	}

	response := comment.ToResponse()
	return &response, nil
}

func (s *CommentService) GetCommentsByPostID(postID uint) ([]models.CommentResponse, error) {
	var post models.Post
	if err := database.DB.First(&post, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		logger.Error("Database error:", err)
		return nil, errors.New("获取文章失败")
	}

	var comments []models.Comment
	if err := database.DB.Preload("User").
		Where("post_id = ?", postID).
		Order("created_at ASC").
		Find(&comments).Error; err != nil {
		logger.Error("Failed to get comments:", err)
		return nil, errors.New("获取评论列表失败")
	}

	responses := make([]models.CommentResponse, len(comments))
	for i, comment := range comments {
		responses[i] = comment.ToResponse()
	}

	return responses, nil
}

func (s *CommentService) DeleteComment(commentID, userID uint) error {
	var comment models.Comment
	if err := database.DB.First(&comment, commentID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("评论不存在")
		}
		logger.Error("Database error:", err)
		return errors.New("获取评论失败")
	}

	if comment.UserID != userID {
		return errors.New("无权限删除此评论")
	}

	if err := database.DB.Delete(&comment).Error; err != nil {
		logger.Error("Failed to delete comment:", err)
		return errors.New("评论删除失败")
	}

	return nil
}
