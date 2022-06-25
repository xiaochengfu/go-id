package main

import (
	"fmt"
	"github.com/xiaochengfu/go-id/configs"
	_ "github.com/xiaochengfu/go-id/configs"
	"github.com/xiaochengfu/go-id/internal/api/controller"
	"net/http"
)
import "github.com/gin-gonic/gin"

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	r.GET("/api/segment/:key", controller.Get())
	// 3.监听端口，默认在8080
	port := configs.Get().Service.Port
	fmt.Println(port)
	if port == 0 {
		port = 8000
	}
	_ = r.Run(fmt.Sprintf(":%d", port))
}
