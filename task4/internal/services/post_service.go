package services

import (
	"blog-system/internal/models"
	"blog-system/pkg/database"
	"blog-system/pkg/logger"
	"errors"

	"gorm.io/gorm"
)

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}

func (s *PostService) CreatePost(userID uint, req *models.PostCreateRequest) (*models.PostResponse, error) {
	post := &models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := database.DB.Create(post).Error; err != nil {
		logger.Error("Failed to create post:", err)
		return nil, errors.New("文章创建失败")
	}

	if err := database.DB.Preload("User").First(post, post.ID).Error; err != nil {
		logger.Error("Failed to load post with user:", err)
		return nil, errors.New("获取文章信息失败")
	}

	response := post.ToResponse()
	return &response, nil
}

func (s *PostService) GetPostByID(id uint) (*models.PostResponse, error) {
	var post models.Post
	if err := database.DB.Preload("User").Preload("Comments.User").First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		logger.Error("Database error:", err)
		return nil, errors.New("获取文章失败")
	}

	response := post.ToResponse()
	return &response, nil
}

func (s *PostService) GetPosts(page, pageSize int) ([]models.PostListResponse, int64, error) {
	var posts []models.Post
	var total int64

	offset := (page - 1) * pageSize

	if err := database.DB.Model(&models.Post{}).Count(&total).Error; err != nil {
		logger.Error("Failed to count posts:", err)
		return nil, 0, errors.New("获取文章总数失败")
	}

	if err := database.DB.Preload("User").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&posts).Error; err != nil {
		logger.Error("Failed to get posts:", err)
		return nil, 0, errors.New("获取文章列表失败")
	}

	responses := make([]models.PostListResponse, len(posts))
	for i, post := range posts {
		responses[i] = post.ToListResponse()
	}

	return responses, total, nil
}

func (s *PostService) UpdatePost(postID, userID uint, req *models.PostUpdateRequest) (*models.PostResponse, error) {
	var post models.Post
	if err := database.DB.First(&post, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		logger.Error("Database error:", err)
		return nil, errors.New("获取文章失败")
	}

	if post.UserID != userID {
		return nil, errors.New("无权限修改此文章")
	}
	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}

	if err := database.DB.Model(&post).Updates(updates).Error; err != nil {
		logger.Error("Failed to update post:", err)
		return nil, errors.New("文章更新失败")
	}

	if err := database.DB.Preload("User").First(&post, postID).Error; err != nil {
		logger.Error("Failed to reload post:", err)
		return nil, errors.New("获取更新后的文章信息失败")
	}

	response := post.ToResponse()
	return &response, nil
}

func (s *PostService) DeletePost(postID, userID uint) error {
	var post models.Post
	if err := database.DB.First(&post, postID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("文章不存在")
		}
		logger.Error("Database error:", err)
		return errors.New("获取文章失败")
	}

	if post.UserID != userID {
		return errors.New("无权限删除此文章")
	}

	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("post_id = ?", postID).Delete(&models.Comment{}).Error; err != nil {
			return err
		}
		return tx.Delete(&post).Error
	}); err != nil {
		logger.Error("Failed to delete post:", err)
		return errors.New("文章删除失败")
	}

	return nil
}
