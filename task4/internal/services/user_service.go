package services

import (
	"blog-system/internal/models"
	"blog-system/pkg/auth"
	"blog-system/pkg/database"
	"blog-system/pkg/logger"
	"errors"

	"gorm.io/gorm"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register(req *models.UserCreateRequest) (*models.UserResponse, error) {
	var existingUser models.User
	if err := database.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("用户名已存在")
	}

	if err := database.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("邮箱已存在")
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		logger.Error("Failed to hash password:", err)
		return nil, errors.New("密码加密失败")
	}

	user := &models.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	if err := database.DB.Create(user).Error; err != nil {
		logger.Error("Failed to create user:", err)
		return nil, errors.New("用户创建失败")
	}

	response := user.ToResponse()
	return &response, nil
}

func (s *UserService) Login(req *models.UserLoginRequest) (*models.UserResponse, string, error) {
	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("用户名或密码错误")
		}
		logger.Error("Database error during login:", err)
		return nil, "", errors.New("登录失败")
	}

	if err := auth.CheckPassword(user.Password, req.Password); err != nil {
		return nil, "", errors.New("用户名或密码错误")
	}

	token, err := auth.GenerateToken(user.ID, user.Username, "ktjnCkMI6GgMN3w6Nein+BFSl7YThGzlmwuomDSvkzo=")
	if err != nil {
		logger.Error("Failed to generate token:", err)
		return nil, "", errors.New("令牌生成失败")
	}

	response := user.ToResponse()
	return &response, token, nil
}

func (s *UserService) GetUserByID(id uint) (*models.UserResponse, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		logger.Error("Database error:", err)
		return nil, errors.New("获取用户信息失败")
	}

	response := user.ToResponse()
	return &response, nil
}
