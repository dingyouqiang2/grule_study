package main

import (
	"grule_study/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // 路由
	r.LoadHTMLGlob("templates/*") // 模板
	routers.RegisterEbsRoutes(r) // 云硬盘
	routers.RegisterGruleRoutes(r) // 计费规则
	routers.RegisterPriceRoutes(r) // 价格调整
	r.GET("/", func(c *gin.Context) { // 首页
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.Run() // 运行在8080端口
}
