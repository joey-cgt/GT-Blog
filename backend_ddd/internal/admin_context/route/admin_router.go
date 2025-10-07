package route

import (
	"gt-blog/backend_ddd/internal/admin_context/controller"
	"gt-blog/backend_ddd/internal/common/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAdminRoutes 注册管理员相关路由
func RegisterAdminRoutes(router *gin.RouterGroup, adminController *controller.AdminController) {
	// 管理员路由组
	adminGroup := router.Group("/admin")
	{
		// 登录路由 - POST /api/v1/admin/login
		// 登录路由不需要认证
		adminGroup.POST("/login", adminController.Login)

		// 需要认证的路由组
		authGroup := adminGroup.Group("")
		// 应用AuthMiddleware中间件
		authGroup.Use(middleware.AuthMiddleware())
		{
			// 获取管理员信息 - GET /api/v1/admin/profile
			authGroup.GET("/profile", adminController.GetAdminInfo)

			// 更新管理员信息 - PUT /api/v1/admin/profile
			authGroup.PUT("/profile", adminController.UpdateAdminInfo)

			// 修改密码 - PUT /api/v1/admin/password
			authGroup.PUT("/password", adminController.ChangePassword)
		}
	}
}
