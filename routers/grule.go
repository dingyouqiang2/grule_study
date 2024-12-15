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
		grule.GET("/", func(c *gin.Context) {
			configKeySlice, err := utils.ReadKeys()
			if err != nil {
				log.Panicln(err)
			}
			c.HTML(http.StatusOK, "grule.html", gin.H{"keySlice": configKeySlice})
		})
		grule.GET("/:key", func(c *gin.Context) {
			key := c.Param("key")
			c.HTML(http.StatusOK, "grule_form.html", gin.H{"key": key})
		})
		grule.POST("/post/", func(c *gin.Context) {
			var form models.RuleForm
			c.ShouldBind(&form)
			utils.CreateGrule("ebs", form)
		})
	}
}