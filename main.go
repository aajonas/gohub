package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	// new 一个 Gin Engine 实例
	r := gin.New()
	// 注册中间件
	r.Use(gin.Logger(),gin.Recovery())
	// 注册一个路由
	r.GET("/", func(c *gin.Context) {
		//以json格式响应
		c.JSON(http.StatusOK,gin.H{
			"Hello":"World",
		})
	})
	
	//处理404请求
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString,"text/html"){
			// 如果是 HTML 的话
			c.String(http.StatusNotFound,"页面返回404")
		}else{
			//默认返回200
			c.JSON(http.StatusOK,gin.H{
				"error_code":404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

	r.Run(":8000")

}