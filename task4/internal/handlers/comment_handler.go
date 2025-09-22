package handlers

import (
	"blog-system/internal/models"
	"blog-system/internal/services"
	"blog-system/internal/utils"
	"blog-system/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *services.CommentService
}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{
		commentService: services.NewCommentService(),
	}
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "用户未认证")
		return
	}

	var req models.CommentCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Invalid request data:", err)
		utils.BadRequest(c, "请求数据格式错误")
		return
	}

	comment, err := h.commentService.CreateComment(userID.(uint), &req)
	if err != nil {
		logger.Error("Create comment failed:", err)
		utils.BadRequest(c, err.Error())
		return
	}

	utils.SuccessWithMessage(c, "评论创建成功", comment)
}

func (h *CommentHandler) GetCommentsByPostID(c *gin.Context) {
	postIDStr := c.Param("id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的文章ID")
		return
	}

	comments, err := h.commentService.GetCommentsByPostID(uint(postID))
	if err != nil {
		logger.Error("Get comments failed:", err)
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, comments)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "用户未认证")
		return
	}

	commentIDStr := c.Param("id")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的评论ID")
		return
	}

	err = h.commentService.DeleteComment(uint(commentID), userID.(uint))
	if err != nil {
		logger.Error("Delete comment failed:", err)
		utils.BadRequest(c, err.Error())
		return
	}

	utils.SuccessWithMessage(c, "评论删除成功", nil)
}
