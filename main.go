package main

import (
	"fmt"
	"grule_study/models"
	"log"
	"net/http"
	"reflect"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()                // 路由
	r.LoadHTMLGlob("templates/*")     // 模板
	r.Static("/static", "./static")   // 静态文件
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

	r.GET("/agollo", func(ctx *gin.Context) {
		c := &config.AppConfig{
			AppID:          "SampleApp",
			Cluster:        "DEV",
			IP:             "http://localhost:8081",
			NamespaceName:  "TEST1.EBS",
			IsBackupConfig: true,
			Secret:         "8c140253795e43c6ba85d7baaf2f4705",
		}

		client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
			return c, nil
		})
		fmt.Println("初始化Apollo配置成功")

		//Use your apollo key to test
		cache := client.GetConfigCache(c.NamespaceName)
		value, _ := cache.Get("key")
		fmt.Println(value)
	})

	r.Run() // 运行在8080端口
}
