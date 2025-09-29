package handler

import (
	"gt-blog/backend_ddd/internal/content_management_context/L0_interface/converter"
	"gt-blog/backend_ddd/internal/content_management_context/L0_interface/dto/request"
	"gt-blog/backend_ddd/internal/content_management_context/L1_application/service"

	"github.com/gin-gonic/gin"
)

type ColumnHandler struct {
	columnAppService *service.ColumnAppService
}

func NewColumnHandler(columnAppService *service.ColumnAppService) *ColumnHandler {
	return &ColumnHandler{
		columnAppService: columnAppService,
	}
}

func (h *ColumnHandler) CreateColumn(c *gin.Context) {
	var req request.CreateColumnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertCreateColumnRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	columnId, err := h.columnAppService.CreateColumn(c, cmd)
	if err != nil {
		c.JSON(500, gin.H{"error": "创建专栏失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"columnId": columnId})
}

func (h *ColumnHandler) UpdateColumn(c *gin.Context) {
	var req request.UpdateColumnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "前端参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertUpdateColumnRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "convert参数格式错误" + err.Error()})
		return
	}

	if err := h.columnAppService.UpdateColumn(c, cmd); err != nil {
		c.JSON(500, gin.H{"error": "更新专栏失败: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "更新专栏成功"})
}

func (h *ColumnHandler) DeleteColumn(c *gin.Context) {
	var req request.DeleteColumnRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertDeleteColumnRequestToCommand(req)

	if err != nil {

		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	if err := h.columnAppService.DeleteColumn(c, cmd); err != nil {
		c.JSON(500, gin.H{"error": "删除专栏失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "删除专栏成功"})
}

func (h *ColumnHandler) GetColumnList(c *gin.Context) {
	var req request.GetColumnListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 50
	}

	cmd, err := converter.ConvertGetColumnListRequestToQuery(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	res, err := h.columnAppService.GetColumnList(c, cmd)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取专栏列表失败: " + err.Error()})
		return
	}

	resp, err := converter.ConvertColumnListResultToResponse(res)
	if err != nil {
		return
	}

	c.JSON(200, gin.H{"data": resp})
}

func (h *ColumnHandler) GetColumnDetail(c *gin.Context) {
	var req request.GetColumnDetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertGetColumnDetailRequestToQuery(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	result, err := h.columnAppService.GetColumnDetail(c, cmd)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取专栏详情失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": result})
}

// 删除了置顶相关功能
