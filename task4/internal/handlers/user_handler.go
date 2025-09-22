package handlers

import (
	"blog-system/internal/models"
	"blog-system/internal/services"
	"blog-system/internal/utils"
	"blog-system/pkg/logger"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: services.NewUserService(),
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req models.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request data:", err)
		utils.BadRequest(c, "请求数据格式错误")
		return
	}

	user, err := h.userService.Register(&req)
	if err != nil {
		logger.Error("Registration failed:", err)
		utils.BadRequest(c, err.Error())
		return
	}

	utils.SuccessWithMessage(c, "注册成功", user)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req models.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request data:", err)
		utils.BadRequest(c, "请求数据格式错误")
		return
	}

	user, token, err := h.userService.Login(&req)
	if err != nil {
		logger.Error("Login failed:", err)
		utils.Unauthorized(c, err.Error())
		return
	}

	utils.SuccessWithMessage(c, "登录成功", gin.H{
		"user":  user,
		"token": token,
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "用户未认证")
		return
	}

	user, err := h.userService.GetUserByID(userID.(uint))
	if err != nil {
		logger.Error("Get profile failed:", err)
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, user)
}
