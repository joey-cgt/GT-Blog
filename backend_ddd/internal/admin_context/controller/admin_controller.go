package controller

import (
	"gt-blog/backend_ddd/internal/admin_context/dto"
	"gt-blog/backend_ddd/internal/admin_context/service"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService *service.AdminService
}

func NewAdminController(adminService *service.AdminService) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

func (ac *AdminController) GetAdminInfo(c *gin.Context) {
	// 从上下文获取userID，这是由AuthMiddleware设置的
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "未提供有效的认证信息",
			"data":    nil,
		})
		return
	}

	// 将interface{}类型转换为uint
	id, ok := userID.(uint)
	if !ok {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "内部服务器错误",
			"data":    nil,
		})
		return
	}

	// 创建请求对象并设置从JWT中提取的ID
	req := &dto.GetAdminInfoReq{
		ID: id,
	}

	resp, err := ac.adminService.GetAdminInfo(c, req)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    resp,
	})
}

func (ac *AdminController) UpdateAdminInfo(c *gin.Context) {
	req := &dto.UpdateAdminInfoReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 从上下文获取userID，这是由AuthMiddleware设置的
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "未提供有效的认证信息",
			"data":    nil,
		})
		return
	}

	// 将interface{}类型转换为uint
	id, ok := userID.(uint)
	if !ok {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "内部服务器错误",
			"data":    nil,
		})
		return
	}

	// 覆盖请求体中的ID，使用从JWT中提取的ID
	req.ID = id

	if err := ac.adminService.UpdateAdminInfo(c, req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

func (ac *AdminController) ChangePassword(c *gin.Context) {
	req := &dto.ChangePasswordReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := ac.adminService.ChangePassword(c, req); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    nil,
	})
}

func (ac *AdminController) Login(c *gin.Context) {
	req := &dto.LoginReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp, err := ac.adminService.Login(c, req)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    resp,
	})
}
