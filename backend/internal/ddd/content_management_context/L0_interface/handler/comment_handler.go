package handler

import (
	"strconv"

	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/converter"
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/dto/request"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/service"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentAppService *service.CommentAppService
}

func NewCommentHandler(commentAppService *service.CommentAppService) *CommentHandler {
	return &CommentHandler{
		commentAppService: commentAppService,
	}
}

func (h *CommentHandler) SubmitComment(c *gin.Context) {
	var req request.CommentSubmitReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "无效的请求参数: " + err.Error(),
			"data":    nil,
		})
		return
	}

	if err := h.commentAppService.SubmitComment(c, req); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "提交评论失败: " + err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "提交评论成功",
		"data":    nil,
	})
}

func (h *CommentHandler) GetArticleComments(c *gin.Context) {
	// 获取文章ID并转换为uint
	articleIdStr := c.Param("articleId")
	articleID, err := strconv.ParseUint(articleIdStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "无效的文章ID: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 获取评论列表
	result, err := h.commentAppService.GetArticleComments(c, uint(articleID))
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "获取评论列表失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 转换为响应DTO
	resp, err := converter.ConvertGetArticleCommentsToResponse(result)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "转换评论数据失败: " + err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取评论列表成功",
		"data":    resp,
	})
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	// 获取评论ID并转换为uint
	commentIdStr := c.Param("commentId")
	commentID, err := strconv.ParseUint(commentIdStr, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "无效的评论ID: " + err.Error(),
			"data":    nil,
		})
		return
	}

	if err := h.commentAppService.DeleteCommentByID(c, uint(commentID)); err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "删除评论失败: " + err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "删除评论成功",
		"data":    nil,
	})
}
