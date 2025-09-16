package middleware

import "github.com/gin-gonic/gin"

// CORS 跨域中间件
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生产环境应替换为具体的前端域名，开发环境可用*
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 允许的HTTP方法
		c.Writer.Header().Set("Access-Control-Allow-Methods",
			"GET, POST, PUT, DELETE, OPTIONS")

		// 允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Authorization, X-Requested-With")

		// 允许暴露的响应头
		c.Writer.Header().Set("Access-Control-Expose-Headers",
			"Content-Length, Content-Type")

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

func AuthMiddleware(c *gin.Context) {
	// 身份验证中间件实现
	c.Next()
}

func Logger(c *gin.Context) {
	// 日志中间件实现
	c.Next()
}
