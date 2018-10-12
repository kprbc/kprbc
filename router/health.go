package router

import "github.com/gin-gonic/gin"

func healthCheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Status": "OK",
	})
}
