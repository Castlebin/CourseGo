package routes

import (
	"CourseGo/app/common/request"
	"CourseGo/app/controllers/app"
	"CourseGo/app/middleware"
	"CourseGo/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 用来测试优雅停机
	router.GET("/test", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "success")
	})

	/* 测试自定义验证器
	curl --location 'http://localhost:8888/api/user/register' --header 'Content-Type: application/json' --data '{
		"name": "张三"
	}'

	curl --location 'http://localhost:8888/api/user/register' --header 'Content-Type: application/json' --data '{
		"name": "张三",
		"mobile": "12345678"
	}'
	*/
	router.POST("/user/register", func(c *gin.Context) {
		var form request.Register
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": request.GetErrorMsg(form, err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	// 用户注册
	/*
		curl --location 'http://localhost:8888/api/auth/register' --header 'Content-Type: application/json' --data '{
		    "name": "张三",
		    "mobile": "18912345678",
		    "password": "123456"
		}'
	*/
	router.POST("/auth/register", app.Register)

	// 用户登录
	/*
		curl --location 'http://localhost:8888/api/auth/login' --header 'Content-Type: application/json' --data '{
		    "mobile": "18912345678",
		    "password": "123456"
		}'
	*/
	router.POST("/auth/login", app.Login)

	// 需要登录保护的路由
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		// 获取用户信息
		/* 将 login 接口返回的 token 放到请求头中，请求头的 key 为 Authorization，value 为 Bearer + 空格 + token
		curl --location --request POST 'http://localhost:8888/api/auth/info' --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc5ODYwNTgsImp0aSI6IjEiLCJpc3MiOiJhcHAiLCJuYmYiOjE2OTc5NDE4NTh9.H_HQ8T8b47Rl_3WmmACLCRjlvtMmGzcnxM198AIY16w' --data ''
		*/
		authRouter.POST("/auth/info", app.Info)
	}
}
