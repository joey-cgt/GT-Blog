package middleware

import (
	"gt-blog/backend/pkg/util"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORS 跨域中间件
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许的域名列表(开发环境)
		allowedOrigins := []string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
			// 生产环境添加你的前端域名
		}

		origin := c.Request.Header.Get("Origin")
		for _, allowed := range allowedOrigins {
			if origin == allowed {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		// 允许的HTTP方法
		c.Writer.Header().Set("Access-Control-Allow-Methods",
			"GET, POST, PUT, DELETE, OPTIONS, PATCH")

		// 允许的请求头(添加文件上传相关头)
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Authorization, X-Requested-With, Content-Disposition")

		// 允许暴露的响应头(添加文件下载相关头)
		c.Writer.Header().Set("Access-Control-Expose-Headers",
			"Content-Length, Content-Type, Content-Disposition")

		// 是否允许携带cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 预检请求缓存时间
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// AuthMiddleware 身份验证中间件实现
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取Authorization
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "未提供认证信息",
				"data":    nil,
			})
			c.Abort()
			return
		}

		// 提取token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "无效的认证格式",
				"data":    nil,
			})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		// 解析和验证token
		userID, err := util.ParseJWT(tokenString)
		if err != nil {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "无效的token: " + err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}

		// 将userID存储到context中
		c.Set("userID", userID)

		c.Next()
	}
}

func Logger(c *gin.Context) {
	// 日志中间件实现
	c.Next()
}
