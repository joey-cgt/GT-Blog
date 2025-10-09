package route

import (
	"gt-blog/backend/internal/mvc/controller"

	"github.com/gin-gonic/gin"
)

func RegisterVisitorRoutes(router *gin.RouterGroup, visitorController *controller.VisitorController) {
	visitorGroup := router.Group("/public")
	{
		visitorGroup.GET("/authorinfo/:id", visitorController.GetAuthorInfo)
	}
}
