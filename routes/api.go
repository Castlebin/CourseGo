package routes

import (
	"CourseGo/app/common/request"
	"CourseGo/app/controllers/app"
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
}
