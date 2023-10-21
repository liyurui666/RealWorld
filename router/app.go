package router

import (
	"RealWorld/middleware"
	"RealWorld/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	ginServer := gin.Default()
	//解决跨域
	ginServer.Use(cors())
	//gin框架入门
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello,gin"})
	})
	//用户注册
	ginServer.POST("/api/users", service.Register)
	//用户登录
	ginServer.POST("/api/users/login", service.Login)
	//获取当前用户
	ginServer.GET("/api/user", middleware.AuthMiddleware(), service.CurrentUser)
	//更新用户
	ginServer.PUT("/api/user", middleware.AuthMiddleware(), service.UpdateUserInfo)
	return ginServer
}

// 解决跨域方法
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
