package handlers

import (
	"blog-system/internal/models"
	"blog-system/internal/services"
	"blog-system/internal/utils"
	"blog-system/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler() *PostHandler {
	return &PostHandler{
		postService: services.NewPostService(),
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "用户未认证")
		return
	}

	var req models.PostCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request data:", err)
		utils.BadRequest(c, "请求数据格式错误")
		return
	}

	post, err := h.postService.CreatePost(userID.(uint), &req)
	if err != nil {
		logger.Error("Create post failed:", err)
		utils.BadRequest(c, err.Error())
		return
	}

	utils.SuccessWithMessage(c, "文章创建成功", post)
}

func (h *PostHandler) GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的文章ID")
		return
	}

	post, err := h.postService.GetPostByID(uint(id))
	if err != nil {
		logger.Error("Get post failed:", err)
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, post)
}

func (h *PostHandler) GetPosts(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	posts, total, err := h.postService.GetPosts(page, pageSize)
	if err != nil {
		logger.Error("Get posts failed:", err)
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, gin.H{
		"posts":      posts,
		"total":      total,
		"page":       page,
		"page_size":  pageSize,
		"total_page": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "用户未认证")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的文章ID")
		return
	}

	var req models.PostUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request data:", err)
		utils.BadRequest(c, "请求数据格式错误")
		return
	}

	post, err := h.postService.UpdatePost(uint(id), userID.(uint), &req)
	if err != nil {
		logger.Error("Update post failed:", err)
		utils.BadRequest(c, err.Error())
		return
	}

	utils.SuccessWithMessage(c, "文章更新成功", post)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "用户未认证")
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的文章ID")
		return
	}

	err = h.postService.DeletePost(uint(id), userID.(uint))
	if err != nil {
		logger.Error("Delete post failed:", err)
		utils.BadRequest(c, err.Error())
		return
	}

	utils.SuccessWithMessage(c, "文章删除成功", nil)
}
