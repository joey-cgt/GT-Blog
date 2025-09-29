package handler

import (
	"gt-blog/backend_ddd/internal/content_management_context/L0_interface/converter"
	"gt-blog/backend_ddd/internal/content_management_context/L0_interface/dto/request"
	"gt-blog/backend_ddd/internal/content_management_context/L1_application/service"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryAppService *service.CategoryAppService
}

func NewCategoryHandler(categoryAppService *service.CategoryAppService) *CategoryHandler {
	return &CategoryHandler{
		categoryAppService: categoryAppService,
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req request.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertCreateCategoryRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	categoryId, err := h.categoryAppService.CreateCategory(c, cmd)
	if err != nil {
		c.JSON(500, gin.H{"error": "创建分类失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"categoryId": categoryId})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	var req request.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertUpdateCategoryRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	if err := h.categoryAppService.UpdateCategory(c, cmd); err != nil {
		c.JSON(500, gin.H{"error": "更新分类失败: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "更新分类成功"})

}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	var req request.DeleteCategoryRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertDeleteCategoryRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	if err := h.categoryAppService.DeleteCategory(c, cmd); err != nil {
		c.JSON(500, gin.H{"error": "删除分类失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "删除分类成功"})

}

func (h *CategoryHandler) GetCategoryList(c *gin.Context) {
	// var req request.GetCategoryListRequest
	// if err := c.ShouldBindQuery(&req); err != nil {
	// 	c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
	// 	return
	// }

	//qry, err := converter.ConvertGetCategoryListRequestToQuery(req)

	// if err != nil {
	// 	c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
	// 	return
	// }

	result, err := h.categoryAppService.GetCategoryList(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取分类列表失败: " + err.Error()})
		return
	}

	resp, err := converter.ConvertCategoryListResultToResponse(result)
	if err != nil {
		c.JSON(500, gin.H{"error": "转换响应格式错误" + err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": resp})
}

func (h *CategoryHandler) GetCategoryDetail(c *gin.Context) {
	var req request.GetCategoryDetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
	}
	qry, err := converter.ConvertGetCategoryDetailRequestToQuery(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	result, err := h.categoryAppService.GetCategoryDetail(c, qry)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取分类详情失败: " + err.Error()})
		return
	}

	resp, err := converter.ConvertCategoryDetailResultToResponse(result)
	if err != nil {
		c.JSON(500, gin.H{"error": "转换响应格式错误" + err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": resp})
}
