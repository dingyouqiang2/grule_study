package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterEbsRoutes(r *gin.Engine) {
	ebsGroup := r.Group("/ebs")
	{
		ebsGroup.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "ebs.html", gin.H{})
		})
	}
}