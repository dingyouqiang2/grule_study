package main

import (
	"grule_study/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()            // 路由
	r.LoadHTMLGlob("templates/*") // 模板
	r.Static("/static", "./static")
	routers.RegisterApiRoutes(r)      // api路由
	routers.RegisterConfigRoutes(r)   // 配置中心路由
	routers.RegisterEbsRoutes(r)      // 云硬盘路由
	routers.RegisterGruleRoutes(r)    // 计费规则路由
	routers.RegisterPriceRoutes(r)    // 价格调整路由
	r.GET("/", func(c *gin.Context) { // 首页路由
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.Run() // 运行在8080端口
}
