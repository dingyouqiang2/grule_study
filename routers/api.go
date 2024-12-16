package routers

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine) {
	configGroup := r.Group("/api")
	{
		configGroup.GET("/config", func(c *gin.Context) {
			data, err := ioutil.ReadFile("config.json")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read config.json"})
				return
			}
       		c.Data(http.StatusOK, "application/json", data)
    	})
		configGroup.POST("/config", func(c *gin.Context) {
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
				return
			}
			err = ioutil.WriteFile("config.json", body, os.ModePerm)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save config.json"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Configuration saved successfully"})
		})
	}
}