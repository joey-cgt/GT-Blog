package handler

import (
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/converter"
	"gt-blog/backend/internal/ddd/content_management_context/L0_interface/dto/request"
	"gt-blog/backend/internal/ddd/content_management_context/L1_application/service"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	tagAppService *service.TagAppService
}

func NewTagHandler(tagAppService *service.TagAppService) *TagHandler {
	return &TagHandler{
		tagAppService: tagAppService,
	}
}

func (h *TagHandler) CreateTag(c *gin.Context) {
	var req request.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertCreateTagRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	tagId, err := h.tagAppService.CreateTag(c, cmd)
	if err != nil {
		c.JSON(500, gin.H{"error": "创建标签失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"tagId": tagId})
}

func (h *TagHandler) UpdateTag(c *gin.Context) {
	var req request.UpdateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertUpdateTagRequestToCommand(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	if err := h.tagAppService.UpdateTag(c, cmd); err != nil {
		c.JSON(500, gin.H{"error": "更新标签失败: " + err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "更新标签成功"})
}

func (h *TagHandler) DeleteTag(c *gin.Context) {
	var req request.DeleteTagRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertDeleteTagRequestToCommand(req)

	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	if err := h.tagAppService.DeleteTag(c, cmd); err != nil {
		c.JSON(500, gin.H{"error": "删除标签失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "删除标签成功"})
}

func (h *TagHandler) GetTagList(c *gin.Context) {
	var req request.GetTagListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 50
	}

	qry, err := converter.ConvertGetTagListRequestToQuery(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	result, err := h.tagAppService.GetTagList(c, qry)

	if err != nil {
		c.JSON(500, gin.H{"error": "获取标签列表失败: " + err.Error()})
		return
	}

	resp, err := converter.ConvertTagListResultToResponse(result)

	if err != nil {
		c.JSON(500, gin.H{"error": "转换标签列表响应失败: " + err.Error()})
		return

	}

	c.JSON(200, gin.H{"data": resp})
}

func (h *TagHandler) GetTagDetail(c *gin.Context) {
	var req request.GetTagDetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	cmd, err := converter.ConvertGetTagDetailRequestToQuery(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "参数格式错误" + err.Error()})
		return
	}

	result, err := h.tagAppService.GetTagDetail(c, cmd)
	if err != nil {
		c.JSON(500, gin.H{"error": "获取标签详情失败: " + err.Error()})
		return
	}

	resp, err := converter.ConvertTagDetailResultToResponse(result)
	if err != nil {
		c.JSON(500, gin.H{"error": "转换标签详情响应失败: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": resp})
}
