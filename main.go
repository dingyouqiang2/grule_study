package main

import (
	"grule_study/models"
	"grule_study/routers"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()            // 路由
	r.LoadHTMLGlob("templates/*") // 模板
	r.Static("/static", "./static") // 静态文件
	routers.RegisterApiRoutes(r)      // api路由
	routers.RegisterConfigRoutes(r)   // 配置中心路由
	r.GET("/", func(c *gin.Context) { // 首页路由
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	
	r.GET("/test", func(c *gin.Context) {
		ebs := models.EBS{}
		ebsType := reflect.TypeOf(ebs)
		for i := 0; i < ebsType.NumField(); i++ {
            field := ebsType.Field(i)
            log.Printf("字段名: %s, 字段类型: %s\n", field.Name, field.Type)
        }
	})
	
	r.Run() // 运行在8080端口
}
