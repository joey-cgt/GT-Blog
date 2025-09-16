package api

import (
	"net/http"
	"strconv"

	"gt-blog/backend/internal/service"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	profile, err := service.GetProfileByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": profile,
	})
}
