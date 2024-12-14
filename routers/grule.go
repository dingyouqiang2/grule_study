package routers

import (
	"grule_study/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterGruleRoutes(r *gin.Engine) {
	grule := r.Group("/grule")
	{
		grule.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "grule.html", gin.H{})
		})
		grule.GET("/ebs/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "grule_ebs.html", gin.H{})
		})
		grule.POST("/post/", controller.WriteGruleForm)
	}
}