package main

import (
	"grule_study/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	routers.RegisterEbsRoutes(r)
	routers.RegisterGruleRoutes(r)
	routers.RegisterPriceRoutes(r)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.Run()
}
