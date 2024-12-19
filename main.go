package main

import (
	"encoding/json"
	"fmt"
	"grule_study/models"
	"log"
	"reflect"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"

	"github.com/gin-gonic/gin"
)

func main() {
	c := &config.AppConfig{
		AppID:          "TYY",
		Cluster:        "DEV",
		IP:             "http://localhost:8081",
		NamespaceName:  "EBS",
		IsBackupConfig: true,
		Secret:         "7637afdc84f8471daf395c3c7171e5e8",
	}
	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	r := gin.Default()                // 路由
	r.LoadHTMLGlob("templates/*")     // 模板
	r.Static("/static", "./static")   // 静态文件
	r.GET("/test", func(c *gin.Context) {
		ebs := models.EBS{}
		ebsType := reflect.TypeOf(ebs)
		for i := 0; i < ebsType.NumField(); i++ {
			field := ebsType.Field(i)
			log.Printf("字段名: %s, 字段类型: %s\n", field.Name, field.Type)
		}
	})
	r.GET("/agollo", func(ctx *gin.Context) {
		cache := client.GetConfigCache(c.NamespaceName)
		value, _ := cache.Get("rules")
		fmt.Printf("%s \n%T\n", value, value)
	})
	r.GET("/rules/", func(ctx *gin.Context) {
		cache := client.GetConfigCache(c.NamespaceName)
		v, _ := cache.Get("rules")
		var rulesMap map[string]string
		json.Unmarshal([]byte(v.(string)), &rulesMap)
		for k, v := range rulesMap {
			log.Println(k, v)
		}
	})
	r.Run() // 运行在8080端口
}
