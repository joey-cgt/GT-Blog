package controller

import (
	"gt-blog/backend/internal/mvc/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VisitorController struct {
	visitorService *service.VisitorService
}

func NewVisitorController(visitorService *service.VisitorService) *VisitorController {
	return &VisitorController{
		visitorService: visitorService,
	}
}

func (v *VisitorController) GetAuthorInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "无效的作者ID",
			"data":    nil,
		})
		return
	}
	resp, err := v.visitorService.GetAuthorInfo(c, uint(id))
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
