package routers

import (
	"grule_study/models"
	"grule_study/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterGruleRoutes(r *gin.Engine) {
	grule := r.Group("/grule")
	{
		// 产品列表
		grule.GET("/", func(c *gin.Context) {
			configKeySlice, err := utils.ReadKeys()
			if err != nil {
				log.Panicln(err)
			}
			c.HTML(http.StatusOK, "grule.html", gin.H{"keySlice": configKeySlice})
		})
		// 产品规则
		grule.GET("/:key/", func(c *gin.Context) {
			key := c.Param("key")
			c.HTML(http.StatusOK, "grule_form.html", gin.H{"key": key})
		})
		// 产品新增
		grule.GET("/key/add/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "grule_key_add.html", gin.H{})
		})
		grule.POST("/key/add/", func(c *gin.Context) {
			key := c.PostForm("name")
			utils.AddKey(key)
			c.Redirect(http.StatusMovedPermanently, "/grule/")
		})
		// grule发布
		grule.POST("/:key/post/", func(c *gin.Context) {
			var form models.RuleForm
			c.ShouldBind(&form)
			key := c.Param("key")
			utils.CreateGrule(key, form)
			c.Redirect(http.StatusMovedPermanently, "/grule/")
		})
	}
}