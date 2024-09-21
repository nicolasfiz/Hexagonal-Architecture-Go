package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitHealthRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Healthy!",
		})
	})

}
