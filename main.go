package main

import (
	"fmt"
	"github.com/aajonas/gohub/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {

	// new 一个 Gin Engine 实例
	r := gin.New()
	//初始化路由绑定
	bootstrap.SetupRoute(r)
	// 运行服务
	err := r.Run(":3000")
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}

}