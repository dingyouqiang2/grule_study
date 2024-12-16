package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterConfigRoutes(r *gin.Engine) {
	configGroup := r.Group("/config")
	{
		configGroup.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "config.html", gin.H{})
		})
	}
}