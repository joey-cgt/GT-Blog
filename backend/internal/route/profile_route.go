package route

import (
	"gt-blog/backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterProfileRoutes(router *gin.Engine) {
	group := router.Group("/api/profiles")
	{
		group.GET("/get/:id", handler.GetProfile)
		group.POST("/update", handler.UpdateProfile)
		group.POST("/avatar/upload", handler.UploadAvatar)
		// 确保没有重复路由
	}
}
